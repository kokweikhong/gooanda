package gooanda

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kokweikhong/gooanda/addr"
	"github.com/kokweikhong/gooanda/kw"
	"github.com/kokweikhong/gooanda/opts"
)

// PricingCandleLatest data structure
type PricingCandleLatest struct { // {{{
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
type PricingInformation struct { // {{{
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
type PricingCandlestickInstrument struct { // {{{
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

type pricing connection

func NewPricingConnection(token string) *pricing {
	return &pricing{token: token}
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
func (pr *pricing) GetCandlesLatest(live bool, accountID string, instruments []string, granularity kw.Granularity, priceComponent kw.PriceComponent, querys ...opts.PricingOpts) (*PricingCandleLatest, error) { // {{{
	querys = append(querys, opts.QPrCandleSpecifications(instruments, granularity, priceComponent))
	q := opts.NewPricingQuery(querys...)
	var url string
	if live {
		url = fmt.Sprintf(addr.PrCandleLatest, addr.LiveHost, accountID)
	} else if !live {
		url = fmt.Sprintf(addr.PrCandleLatest, addr.PracticeHost, accountID)
	}
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
	var data = &PricingCandleLatest{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// GetPricingInformation
func (pr *pricing) GetPricingInformation(live bool, accountID string, instruments []string, querys ...opts.PricingOpts) (*PricingInformation, error) { // {{{
	querys = append(querys, opts.QPrInstruments(instruments))
	q := opts.NewPricingQuery(querys...)
	var url string
	if live {
		url = fmt.Sprintf(addr.PrPricingInfo, addr.LiveHost, accountID)
	} else if !live {
		url = fmt.Sprintf(addr.PrPricingInfo, addr.PracticeHost, accountID)
	}
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
	var data = &PricingInformation{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// GetStreamingPrice
func (pr *pricing) GetStreamingPrice(live bool, accountID string, instruments []string, querys ...opts.PricingOpts) ([]byte, error) { // {{{
	querys = append(querys, opts.QPrInstruments(instruments))
	q := opts.NewPricingQuery(querys...)
	var url string
	if live {
		url = fmt.Sprintf(addr.PrPricingStream, addr.LiveHost, accountID)
	} else if !live {
		url = fmt.Sprintf(addr.PrPricingStream, addr.PracticeHost, accountID)
	}
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
func (pr *pricing) GetCandlestickInstrument(live bool, accountID string, instrument string, querys ...opts.PricingOpts) (*PricingCandlestickInstrument, error) { // {{{
	q := opts.NewPricingQuery(querys...)
	var url string
	if live {
		url = fmt.Sprintf(addr.PrInstrumentCandles,
			addr.LiveHost, accountID, instrument)
	} else if !live {
		url = fmt.Sprintf(addr.PrInstrumentCandles,
			addr.PracticeHost, accountID, instrument)
	}
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
	var data = &PricingCandlestickInstrument{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}
