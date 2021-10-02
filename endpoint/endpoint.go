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

type order struct {
	Orders, PendingOrder, OrderDetails rest
}

type trade struct {
	Trades, OpenTrades, TradeDetails         rest
	CloseTrade, UpdateClientExt, UpdateTrade rest
}

type position struct {
	PositionList, OpenPositionList                       rest
	SingleInstrumentPosition, ClosePositionForInstrument rest
}

type transaction struct {
	Transactions, TransactionById, TransactionIdRange rest
	TransactionSinceId                                rest
	TransactionStream                                 stream
}

var (
	Account     *account
	Instrument  *instrument
	Pricing     *pricing
	Order       *order
	Trade       *trade
	Position    *position
	Transaction *transaction
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
	Order = &order{
		Orders:       "/v3/accounts/%v/orders",
		PendingOrder: "/v3/accounts/%v/pendingOrders",
		OrderDetails: "/v3/accounts/%v/orders/%v",
	}
	Trade = &trade{
		Trades:          "/v3/accounts/%v/trades",
		OpenTrades:      "/v3/accounts/%v/openTrades",
		TradeDetails:    "/v3/accounts/%v/trades/%v",
		CloseTrade:      "/v3/accounts/%v/trades/%v/close",
		UpdateClientExt: "/v3/accounts/%v/trades/%v/clientExtensions",
		UpdateTrade:     "/v3/accounts/%v/trades/%v/orders",
	}
	Position = &position{
		PositionList:               "/v3/accounts/%v/positions",
		OpenPositionList:           "/v3/accounts/%v/openPositions",
		SingleInstrumentPosition:   "/v3/accounts/%v/positions/%v",
		ClosePositionForInstrument: "/v3/accounts/%v/positions/%v/close",
	}
	Transaction = &transaction{
		Transactions:       "/v3/accounts/%v/transactions",
		TransactionById:    "/v3/accounts/%v/transactions/%v",
		TransactionIdRange: "/v3/accounts/%v/transactions/idrange",
		TransactionSinceId: "/v3/accounts/%v/transactions/sinceid",
		TransactionStream:  "/v3/accounts/%v/transactions/stream",
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
