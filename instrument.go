package gooanda

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kokweikhong/gooanda/addr"
	"github.com/kokweikhong/gooanda/opts"
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

type instrument connection

// NewInstrumentConnection create new connection for instrument endpoint
func NewInstrumentConnection(token string) *instrument {
	return &instrument{token: token}
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
func (in *instrument) GetCandles(live bool, instrument string, querys ...opts.InstrumentOpts) (*InstrumentCandles, error) {
	q := opts.NewInstrumentQuery(querys...)
	var url string
	if live {
		url = fmt.Sprintf(addr.InstrumentCandles, addr.LiveHost, instrument)
	} else if !live {
		url = fmt.Sprintf(addr.InstrumentCandles, addr.PracticeHost, instrument)
	}
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
}

// GetInstrumentOrderBook: Fetch an order book for an instrument.
func (in *instrument) GetOrderBook(live bool, instrument string, querys ...opts.InstrumentOpts) (*InstrumentOrderBook, error) {
	q := opts.NewInstrumentQuery(querys...)
	var url string
	if live {
		url = fmt.Sprintf(addr.InstrumentOrderBook, addr.LiveHost, instrument)
	} else if !live {
		url = fmt.Sprintf(addr.InstrumentOrderBook, addr.PracticeHost, instrument)
	}
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
}

// GetInstrumentPositionBook: Fetch a position book for an instrument.
func (in *instrument) GetPositionBook(live bool, instrument string, querys ...opts.InstrumentOpts) (*InstrumentPositionBook, error) {
	q := opts.NewInstrumentQuery(querys...)
	var url string
	if live {
		url = fmt.Sprintf(addr.InstrumentPositionBook, addr.LiveHost, instrument)
	} else if !live {
		url = fmt.Sprintf(addr.InstrumentPositionBook, addr.PracticeHost, instrument)
	}
	fmt.Println(url)
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
}
