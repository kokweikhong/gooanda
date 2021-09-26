package endpoint

import (
	"fmt"
)

type rest string
type stream string

const (
	streamApi    = "https://stream-%v.oanda.com"
	restApi      = "https://api-%v.oanda.com"
	liveHost     = "fxtrade"
	practiceHost = "fxpractice"
)

type instrument struct {
	InstrumentCandles, InstrumentOrderBook, InstrumentPositionBook rest
}

type account struct {
	Accounts, AccountsById, AccountSummary, AccountInstrument rest
}

type pricing struct {
	CandleLatest, PricingInfo, InstrumentCandles rest
	PricingStream                                stream
}

var (
	Account    *account
	Instrument *instrument
	Pricing    *pricing
)

func init() {
	Account = &account{
		Accounts:          "/v3/accounts",
		AccountsById:      "/v3/accounts/%v",
		AccountSummary:    "/v3/accounts/%v/summary",
		AccountInstrument: "/v3/accounts/%v/instruments",
	}
	Instrument = &instrument{
		InstrumentCandles:      "/v3/instruments/%v/candles",
		InstrumentOrderBook:    "/v3/instruments/%v/orderBook",
		InstrumentPositionBook: "/v3/instruments/%v/positionBook",
	}
	Pricing = &pricing{
		CandleLatest:      "/v3/accounts/%v/candles/latest",
		PricingInfo:       "/v3/accounts/%v/pricing",
		PricingStream:     "/v3/accounts/%v/pricing/stream",
		InstrumentCandles: "/v3/accounts/%v/instruments/%v/candles",
	}
}

// GetEndpoint is to get the streaming or rest api endpoint.
// and also for live or practice host.
func GetEndpoint(isLive bool, endpoint interface{}) string {
	var host, url string
	if isLive {
		host = liveHost
	} else if !isLive {
		host = practiceHost
	}
	switch t := endpoint.(type) {
	case rest:
		url = fmt.Sprintf(restApi, host)
		return url + string(t)
	case stream:
		url = fmt.Sprintf(streamApi, host)
		return url + string(t)
	}
	return ""
}
