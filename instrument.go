package gooanda

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/kokweikhong/gooanda/endpoint"
)

// InstrumentCandles data structure
type InstrumentCandles struct {
	Candles []struct {
		Ask      struct{ instrumentOHLC } `json:"ask"`
		Bid      struct{ instrumentOHLC } `json:"bid"`
		Mid      struct{ instrumentOHLC } `json:"mid"`
		Complete bool                     `json:"complete"`
		Time     string                   `json:"time"`
		Volume   float64                  `json:"volume"`
	} `json:"candles"`
	Granularity string `json:"granularity"`
	Instrument  string `json:"instrument"`
}

// InstrumentOrderBook data structure
type InstrumentOrderBook struct {
	OrderBook instrumentBook `json:"orderBook"`
}

// InstrumentPositionBook data structure
type InstrumentPositionBook struct {
	PositionBook instrumentBook `json:"positionBook"`
}

type instrumentBook struct {
	Instrument  string `json:"instrument"`
	Time        string `json:"time"`
	Price       string `json:"price"`
	BucketWidth string `json:"bucketWidth"`
	Buckets     []struct {
		Price             string `json:"price"`
		LongCountPercent  string `json:"longCountPercent"`
		ShortCountPercent string `json:"shortCountPercent"`
	} `json:"buckets"`
}

type instrumentOHLC struct {
	Close string `json:"c"`
	High  string `json:"h"`
	Low   string `json:"l"`
	Open  string `json:"o"`
}

type instrument struct {
	connection
	Query *instrumentFunc
}

// NewInstrumentConnection create new connection for instrument endpoint
func NewInstrumentConnection(token string) *instrument {
	conn := &instrument{}
	conn.token = token
	return conn
}

func (in *instrument) connect() ([]byte, error) {
	con := &connection{in.endpoint, in.method, in.token, in.data}
	resp, err := con.connect()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetInstrumentCandles: Fetch candlestick data for an instrument.
func (in *instrument) GetCandles(live bool, instrument string, querys ...instrumentOpts) (*InstrumentCandles, error) { // {{{
	q := newInstrumentQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Instrument.InstrumentCandles)
	url := fmt.Sprintf(ep, instrument)
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	in.endpoint = u
	in.method = http.MethodGet
	resp, err := in.connect()
	if err != nil {
		return nil, err
	}
	var data = &InstrumentCandles{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// GetInstrumentOrderBook: Fetch an order book for an instrument.
func (in *instrument) GetOrderBook(live bool, instrument string, querys ...instrumentOpts) (*InstrumentOrderBook, error) { // {{{
	q := newInstrumentQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Instrument.InstrumentOrderBook)
	url := fmt.Sprintf(ep, instrument)
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	in.endpoint = u
	in.method = http.MethodGet
	resp, err := in.connect()
	if err != nil {
		return nil, err
	}
	var data = &InstrumentOrderBook{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// GetInstrumentPositionBook: Fetch a position book for an instrument.
func (in *instrument) GetPositionBook(live bool, instrument string, querys ...instrumentOpts) (*InstrumentPositionBook, error) { // {{{
	q := newInstrumentQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Instrument.InstrumentPositionBook)
	url := fmt.Sprintf(ep, instrument)
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	in.endpoint = u
	in.method = http.MethodGet
	resp, err := in.connect()
	if err != nil {
		return nil, err
	}
	var data = &InstrumentPositionBook{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

type instrumentQuery struct {
	Price             string `json:"price,omitempty"`
	Granularity       string `json:"granularity,omitempty"`
	Count             string `json:"count,omitempty"`
	From              string `json:"from,omitempty"`
	To                string `json:"to,omitempty"`
	Smooth            string `json:"smooth,omitempty"`
	IncludeFirst      string `json:"includeFirst,omitempty"`
	DailyAlignment    string `json:"dailyAlignment,omitempty"`
	AlignmentTimezone string `json:"alignmentTimezone,omitempty"`
	WeeklyAlignment   string `json:"weeklyAlignment,omitempty"`
	Time              string `json:"time,omitempty"`
}

type instrumentOpts func(*instrumentQuery)

func newInstrumentQuery(querys ...instrumentOpts) *instrumentQuery {
	q := &instrumentQuery{}
	for _, query := range querys {
		query(q)
	}
	return q
}

type instrumentFunc struct{}

// The time of the snapshot to fetch. If not specified, then the most recent snapshot is fetched.
func (*instrumentFunc) WithTime(setTime time.Time) instrumentOpts {
	return func(iq *instrumentQuery) {
		if setTime.Unix() > time.Now().Unix() {
			iq.Time = ""
		}
		iq.Time = setTime.Format(time.RFC3339)
	}
}

// The day of the week used for granularities that have weekly alignment. [default=Friday]
func (*instrumentFunc) WithWeeklyAlignment(weeklyAlignment string) instrumentOpts {
	return func(iq *instrumentQuery) {
		iq.WeeklyAlignment = weeklyAlignment
	}
}

// The timezone to use for the dailyAlignment parameter. Candlesticks with daily alignment will be aligned to the dailyAlignment hour within the alignmentTimezone. Note that the returned times will still be represented in UTC. [default=America/New_York]
func (*instrumentFunc) WithAlignmentTimezone(timezone string) instrumentOpts {
	return func(iq *instrumentQuery) {
		iq.AlignmentTimezone = timezone
	}
}

// The hour of the day (in the specified timezone) to use for granularities that have daily alignments. [default=17, minimum=0, maximum=23]
func (*instrumentFunc) WithDailyAlignment(alignment int) instrumentOpts {
	return func(iq *instrumentQuery) {
		if alignment > 23 || alignment < 0 {
			iq.DailyAlignment = "17"
			return
		}
		iq.DailyAlignment = strconv.Itoa(alignment)
	}
}

// A flag that controls whether the candlestick that is covered by the from time should be included in the results. This flag enables clients to use the timestamp of the last completed candlestick received to poll for future candlesticks but avoid receiving the previous candlestick repeatedly. [default=True]
func (*instrumentFunc) WithWithoutIncludeFirst() instrumentOpts {
	return func(iq *instrumentQuery) {
		iq.IncludeFirst = "false"
	}
}

// A flag that controls whether the candlestick is “smoothed” or not. A smoothed candlestick uses the previous candle’s close price as its open price, while an un-smoothed candlestick uses the first price from its time range as its open price. [default=False]
func (*instrumentFunc) WithWithSmooth() instrumentOpts {
	return func(iq *instrumentQuery) {
		iq.Smooth = "true"
	}
}

// from: The start of the time range to fetch candlesticks for.
func (*instrumentFunc) WithFrom(from time.Time) instrumentOpts {
	return func(iq *instrumentQuery) {
		if from.Unix() < time.Now().Unix() {
			iq.From = ""
			return
		}
		iq.From = from.Format(time.RFC3339)
	}
}

// from: The start of the time range to fetch candlesticks for.
// to: The end of the time range to fetch candlesticks for.
func (*instrumentFunc) WithFromTo(from, to time.Time) instrumentOpts {
	return func(iq *instrumentQuery) {
		if to.Unix() > from.Unix() || to.Unix() > time.Now().Unix() {
			iq.From = ""
			iq.To = ""
			return
		}
		iq.From = from.Format(time.RFC3339)
		fmt.Println(iq.From)
		iq.To = to.Format(time.RFC3339)
	}
}

// The number of candlesticks to return in the response. Count should not be specified if both the start and end parameters are provided, as the time range combined with the granularity will determine the number of candlesticks to return. [default=500, maximum=5000]
func (*instrumentFunc) WithCount(count int) instrumentOpts {
	return func(iq *instrumentQuery) {
		c := strconv.Itoa(count)
		if count > 5000 {
			iq.Count = "500"
		}
		iq.Count = c
	}
}

// The granularity of the candlesticks to fetch [default=S5]
func (iq *instrumentFunc) WithGranularity(granularity string) instrumentOpts {
	return func(iq *instrumentQuery) {
		iq.Granularity = string(granularity)
	}
}

// The Price component(s) to get candlestick data for. [default=M]
func (*instrumentFunc) WithPrice(price string) instrumentOpts {
	return func(iq *instrumentQuery) {
		iq.Price = price
	}
}
