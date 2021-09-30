package gooanda

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kokweikhong/gooanda/endpoint"
)

// PricingCandleLatest data structure
type pricingCandleLatest struct { // {{{
	LatestCandles []struct {
		Instrument  string `json:"instrument"`
		Granularity string `json:"granularity"`
		Candles     []struct {
			Complete bool    `json:"complete"`
			Volume   float64 `json:"volume"`
			Time     string  `json:"time"`
			Bid      ohlc    `json:"bid,omitempty"`
			Ask      ohlc    `json:"ask,omitempty"`
			Mid      ohlc    `json:"mid,omitempty"`
		} `json:"candles"`
	} `json:"latestCandles"`
} // }}}

// PricingInformation data structure
type pricingInformation struct { // {{{
	Time   string `json:"time"`
	Prices []struct {
		Type string `json:"type"`
		Time string `json:"time"`
		Bids []struct {
			Price     string  `json:"price"`
			Liquidity float64 `json:"liquidity"`
		} `json:"bids"`
		Asks []struct {
			Price     string  `json:"price"`
			Liquidity float64 `json:"liquidity"`
		} `json:"asks"`
		CloseoutBid                string `json:"closeoutBid"`
		CloseoutAsk                string `json:"closeoutAsk"`
		Status                     string `json:"status"`
		Tradeable                  bool   `json:"tradeable"`
		QuoteHomeConversionFactors struct {
			PositiveUnits string `json:"positiveUnits"`
			NegativeUnits string `json:"negativeUnits"`
			Instrument    string `json:"instrument"`
		} `json:"quoteHomeConversionFactors"`
	} `json:"prices"`
} // }}}

// PricingCandlestickInstrument data structure
type pricingCandlestickInstrument struct { // {{{
	Instrument  string `json:"instrument"`
	Granularity string `json:"granularity"`
	Candles     []struct {
		Complete bool    `json:"complete"`
		Volume   float64 `json:"volume"`
		Time     string  `json:"time"`
		Mid      ohlc    `json:"mid,omitempty"`
		Bid      ohlc    `json:"bid,omitempty"`
		Ask      ohlc    `json:"ask,omitempty"`
	} `json:"candles"`
} // }}}

type ohlc struct {
	Open  string `json:"o"`
	High  string `json:"h"`
	Low   string `json:"l"`
	Close string `json:"c"`
}

type pricing struct {
	connection
	Query *pricingFunc
}

func NewPricingConnection(token string) *pricing {
	conn := &pricing{}
	conn.token = token
	return conn
}

func (pr *pricing) connect() ([]byte, error) {
	con := &connection{pr.endpoint, pr.method, pr.token, pr.data}
	resp, err := con.connect()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetCandlesLatest get dancing bears and most recently completed candles within an Account for specified combinations of instrument, granularity, and price component.
func (pr *pricing) GetCandlesLatest(live bool, accountID string, instruments []string, granularity string, priceComponent string, querys ...pricingOpts) (*pricingCandleLatest, error) { // {{{
	querys = append(querys, pr.Query.WithCandleSpecifications(instruments, granularity, priceComponent))
	q := newPricingQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Pricing.CandleLatest)
	url := fmt.Sprintf(ep, accountID)
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	pr.endpoint = u
	pr.method = http.MethodGet
	resp, err := pr.connect()
	if err != nil {
		return nil, err
	}
	var data = &pricingCandleLatest{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// GetPricingInformation
func (pr *pricing) GetPricingInformation(live bool, accountID string, instruments []string, querys ...pricingOpts) (*pricingInformation, error) { // {{{
	querys = append(querys, pr.Query.WithInstruments(instruments))
	q := newPricingQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Pricing.PricingInfo)
	url := fmt.Sprintf(ep, accountID)
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	pr.endpoint = u
	pr.method = http.MethodGet
	resp, err := pr.connect()
	if err != nil {
		return nil, err
	}
	var data = &pricingInformation{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// GetStreamingPrice
func (pr *pricing) GetStreamingPrice(live bool, accountID string, instruments []string, querys ...pricingOpts) ([]byte, error) { // {{{
	querys = append(querys, pr.Query.WithInstruments(instruments))
	q := newPricingQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Pricing.PricingStream)
	url := fmt.Sprintf(ep, accountID)
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	fmt.Println(u)
	pr.endpoint = u
	pr.method = http.MethodGet
	resp, err := pr.connect()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(resp))
	return nil, nil
} // }}}

// GetCandlestickInstrument fetch candlestick data for an instrument.
func (pr *pricing) GetCandlestickInstrument(live bool, accountID string, instrument string, querys ...pricingOpts) (*pricingCandlestickInstrument, error) { // {{{
	q := newPricingQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Pricing.InstrumentCandles)
	url := fmt.Sprintf(ep, accountID, instrument)
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	pr.endpoint = u
	pr.method = http.MethodGet
	resp, err := pr.connect()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var data = &pricingCandlestickInstrument{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

type pricingQuery struct {
	// List of CandleSpecification (csv)
	// List of candle specifications to get pricing for. [required]
	CandleSpecifications   string `json:"candleSpecifications,omitempty"`
	Units                  string `json:"units,omitempty"`
	Smooth                 string `json:"smooth,omitempty"`
	DailyAlignment         string `json:"dailyAlignment,omitempty"`
	AlignmentTimezone      string `json:"alignmentTimezone,omitempty"`
	WeeklyAlignment        string `json:"weeklyAlignment,omitempty"`
	Instruments            string `json:"instruments,omitempty"`
	Since                  string `json:"since,omitempty"`
	IncludeHomeConversions string `json:"includeHomeConversions,omitempty"`
	Snapshot               string `json:"snapshot,omitempty"`
	Price                  string `json:"price,omitempty"`
	Granularity            string `json:"granularity,omitempty"`
	Count                  string `json:"count,omitempty"`
	From                   string `json:"from,omitempty"`
	To                     string `json:"to,omitempty"`
	IncludeFirst           string `json:"includeFirst,omitempty"`
}

type pricingOpts func(*pricingQuery)

func newPricingQuery(querys ...pricingOpts) *pricingQuery {
	q := &pricingQuery{}
	for _, query := range querys {
		query(q)
	}
	return q
}

type pricingFunc struct{}

func (*pricingFunc) WithCandleSpecifications(instruments []string, granularity, priceComponent string) pricingOpts {
	return func(pq *pricingQuery) {
		for k := range instruments {
			instruments[k] = instruments[k] + ":" + string(granularity) + ":" + string(priceComponent)
		}
		pq.CandleSpecifications = strings.Join(instruments, ",")
	}
}

// QrPrFromTo:
// The start of the time range to fetch candlesticks for.
// The end of the time range to fetch candlesticks for.
func (*pricingFunc) WithFromTo(from, to time.Time) pricingOpts {
	return func(pq *pricingQuery) {
		if to.Unix() < from.Unix() || to.Unix() > time.Now().Unix() {
			pq.From = ""
			pq.To = ""
			return
		} else {
			pq.From = from.Format(time.RFC3339)
			pq.To = to.Format(time.RFC3339)
		}
	}
}

// WithSince: Date/Time filter to apply to the response. Only prices and home conversions (if requested) with a time later than this filter (i.e. the price has changed after the since time) will be provided, and are filtered independently.
func (*pricingFunc) WithSince(since time.Time) pricingOpts {
	return func(pq *pricingQuery) {
		if since.Unix() > time.Now().Unix() {
			pq.Since = time.Now().Format(time.RFC3339)
			return
		} else {
			pq.Since = since.Format(time.RFC3339)
		}
	}
}

// WithUnits: The number of units used to calculate the volume-weighted average bid and ask prices in the returned candles. [default=1]
func (*pricingFunc) WithUnits(units float64) pricingOpts {
	return func(pq *pricingQuery) {
		if units < 1 {
			pq.Units = "1"
			return
		} else {
			pq.Units = fmt.Sprintf("%v", units)
		}
	}
}

// WithInstruments: List of InstrumentName (csv)	List of Instruments to get pricing for. [required]
func (*pricingFunc) WithInstruments(instruments []string) pricingOpts {
	return func(pq *pricingQuery) {
		pq.Instruments = strings.Join(instruments, ",")
	}
}

// WithPrice: The Price component(s) to get candlestick data for. [default=M]
func (*pricingFunc) WithPrice(priceComponent string) pricingOpts {
	return func(pq *pricingQuery) {
		pq.Price = priceComponent
	}
}

// WithCount: The number of candlesticks to return in the response. Count should not be specified if both the start and end parameters are provided, as the time range combined with the granularity will determine the number of candlesticks to return. [default=500, maximum=5000]
func (*pricingFunc) WithCount(count int) pricingOpts {
	return func(pq *pricingQuery) {
		if count < 1 || count > 5000 {
			pq.Count = "500"
			return
		}
		pq.Count = strconv.Itoa(count)
	}
}

// WithoutIncludeFirst:  flag that controls whether the candlestick that is covered by the from time should be included in the results. This flag enables clients to use the timestamp of the last completed candlestick received to poll for future candlesticks but avoid receiving the previous candlestick repeatedly. [default=True]
func (*pricingFunc) WithoutIncludeFirst() pricingOpts {
	return func(pq *pricingQuery) { pq.IncludeFirst = "false" }
}

// WithGranularity:  granularity of the candlesticks to fetch [default=S5]
func (*pricingFunc) WithGranularity(granularity string) pricingOpts {
	return func(pq *pricingQuery) { pq.Granularity = granularity }
}

// WithoutSnapshot: that enables/disables the sending of a pricing snapshot when initially connecting to the stream. [default=True]
func (*pricingFunc) WithoutSnapshot() pricingOpts {
	return func(pq *pricingQuery) { pq.Snapshot = "false" }
}

// WithIncludeHomeConversions: that enables the inclusion of the homeConversions field in the returned response. An entry will be returned for each currency in the set of all base and quote currencies present in the requested instruments list. [default=False]
func (*pricingFunc) WithIncludeHomeConversions() pricingOpts {
	return func(pq *pricingQuery) { pq.IncludeHomeConversions = "true" }
}

// WithWeeklyAlignment:  day of the week used for granularities that have weekly alignment. [default=Friday]
func (*pricingFunc) WithWeeklyAlignment(weeklyAlignment string) pricingOpts {
	return func(pq *pricingQuery) { pq.WeeklyAlignment = weeklyAlignment }
}

// WithAlignmentTimezone:  timezone to use for the dailyAlignment parameter. Candlesticks with daily alignment will be aligned to the dailyAlignment hour within the alignmentTimezone. Note that the returned times will still be represented in UTC. [default=America/New_York]
func (*pricingFunc) WithAlignmentTimezone(timezone string) pricingOpts {
	return func(pq *pricingQuery) { pq.AlignmentTimezone = timezone }
}

// WithDailyAlignment: hour of the day (in the specified timezone) to use for granularities that have daily alignments. [default=17, minimum=0, maximum=23]
func (*pricingFunc) WithDailyAlignment(alignment int) pricingOpts {
	return func(pq *pricingQuery) {
		if alignment < 0 || alignment > 23 {
			pq.DailyAlignment = ""
			return
		}
		pq.DailyAlignment = strconv.Itoa(alignment)
	}
}

// WithSmooth: flag that controls whether the candlestick is “smoothed” or not. A smoothed candlestick uses the previous candle’s close price as its open price, while an unsmoothed candlestick uses the first price from its time range as its open price. [default=False]
func (*pricingFunc) WithSmooth() pricingOpts {
	return func(pq *pricingQuery) { pq.Smooth = "true" }
}
