package gooanda

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kokweikhong/gooanda/endpoint"
)

type tradeList struct {
	LastTransactionID string      `json:"lastTransactionID"`
	Trades            []dataTrade `json:"trades"`
}

type specificTrade struct {
	LastTransactionID string    `json:"lastTransactionID"`
	Trade             dataTrade `json:"trade"`
}

type dataTrade struct {
	CurrentUnits float64 `json:"currentUnits,string"`
	Financing    float64 `json:"financing,string"`
	Id           string  `json:"id"`
	InitialUnits float64 `json:"initialUnits,string"`
	Instrument   string  `json:"instrument"`
	OpenTime     string  `json:"openTime"`
	Price        float64 `json:"price,string"`
	RealizePL    float64 `json:"realizedPL,string"`
	State        string  `json:"state"`
	UnrealizePL  float64 `json:"unrealizedPL,string"`
}

type trade struct {
	connection
	Query             *tradeFunc
	TakeProftStopLoss *requestTPSL
}

// NewTradeConnection is to crete connection for TRADE api.
func NewTradeConnection(token string) *trade {
	con := &trade{}
	con.connection.token = token
	con.TakeProftStopLoss = &requestTPSL{}
	return con
}

// GetTradeList is to get a list of Trades for an Account.
func (tr *trade) GetTradeList(live bool, accountID string, opts ...tradeOpts) (*tradeList, error) { // {{{
	var result *tradeList
	query := newTradeQuery(opts...)
	ep := fmt.Sprintf(endpoint.GetEndpoint(live, endpoint.Trade.Trades), accountID)
	url, err := urlAddQuery(ep, query)
	if err != nil {
		return result, err
	}
	tr.method = http.MethodGet
	tr.endpoint = url
	resp, err := tr.connect()
	if err != nil {
		return result, err
	}
	fmt.Println(string(resp))
	if err = json.Unmarshal(resp, &result); err != nil {
		return result, fmt.Errorf("failed to unmarshal data in GetTradeList, %v", err)
	}
	return result, nil
} // }}}

// GetOpenTradeList is to get the list of open Trades for an Account.
func (tr *trade) GetOpenTradeList(live bool, accountID string) (*tradeList, error) { // {{{
	var result *tradeList
	ep := fmt.Sprintf(endpoint.GetEndpoint(live, endpoint.Trade.OpenTrades), accountID)
	tr.method = http.MethodGet
	tr.endpoint = ep
	resp, err := tr.connect()
	if err != nil {
		return result, fmt.Errorf("GetOpenTradeList connect error, %v", err)
	}
	if err = json.Unmarshal(resp, &result); err != nil {
		return result, fmt.Errorf("GetOpenTradeList unmarshal error, %v", err)
	}
	return result, nil
} // }}}

// GetSpecificTradeDetails is to get the details of a specific Trade in an Account.
func (tr *trade) GetSpecificTradeDetails(live bool, accountID, tradeID string) (*specificTrade, error) { // {{{
	var result *specificTrade
	ep := fmt.Sprintf(endpoint.GetEndpoint(live, endpoint.Trade.TradeDetails),
		accountID, tradeID)
	tr.method = http.MethodGet
	tr.endpoint = ep
	resp, err := tr.connect()
	if err != nil {
		return result, fmt.Errorf("GetSpecificTradeDetails connect error, %v", err)
	}
	if err = json.Unmarshal(resp, &result); err != nil {
		return result, fmt.Errorf("GetSpecificTradeDetails unmarshal error, %v", err)
	}
	return result, nil
} // }}}

// CloseTrade is to close (partially or fully) a specific open Trade in an Account.
func (tr *trade) CloseTrade(live bool, accountID, tradeID string, units interface{}) (string, error) { // {{{
	ep := fmt.Sprintf(endpoint.GetEndpoint(live, endpoint.Trade.CloseTrade),
		accountID, tradeID)
	switch t := units.(type) {
	case string:
		if !strings.EqualFold(t, "all") {
			return "", fmt.Errorf("units if string must only be all")
		} else {
			units = strings.ToUpper(t)
		}
	case float64:
		if t <= 0 {
			return "", fmt.Errorf("%v must be greater than 0", units)
		}
	case int:
		if t <= 0 {
			return "", fmt.Errorf("%v must be greater than 0", units)
		}
	}
	tr.method = http.MethodPut
	tr.endpoint = ep
	tr.data = []byte(fmt.Sprintf(`{"units":"%v"}`, units))
	resp, err := tr.connect()
	if err != nil {
		return "", err
	}
	return string(resp), nil
} // }}}

// UpdateTPSLForTrade is to create, replace and cancel a Trade’s dependent
// Orders (Take Profit, Stop Loss and Trailing Stop Loss) through the Trade itself
func (tr *trade) UpdateTPSLForTrade(live bool, accountID, tradeID string) (string, error) { // {{{
	body, err := json.Marshal(tr.TakeProftStopLoss)
	if err != nil {
		log.Fatalf("failed to unmarshal takeprofit and stop loss, %v", err)
	}
	tr.endpoint = fmt.Sprintf(endpoint.GetEndpoint(live,
		endpoint.Trade.UpdateTrade), accountID, tradeID)
	tr.data = body
	tr.method = http.MethodPut
	resp, err := tr.connect()
	if err != nil {
		return "", err
	}
	return string(resp), nil
} // }}}

type requestTPSL struct {
	TakeProfit         *tpsl `json:"takeProfit,omitempty"`
	StopLoss           *tpsl `json:"stopLoss,omitempty"`
	TrailingStopLoss   *tpsl `json:"trailingStopLoss,omitempty"`
	GuaranteedStopLoss *tpsl `json:"guaranteedStopLoss,omitempty"`
}

type tpsl struct {
	Price       float64 `json:"price,omitempty,string"`
	TimeInForce string  `json:"timeInForce,omitempty"`
	GtdTime     string  `json:"gtdTime,omitempty"`
	Distance    float64 `json:"distance,omitempty,string"`
}

// WithTakeProfit specifies the details of a Take Profit Order to be
// created on behalf of a client. This may happen when an Order is filled
// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
// Take Profit Order is modified directly through the Trade.
func (rtpsl *requestTPSL) WithTakeProfit(price float64, timeInForce, gtdTime string) {
	rtpsl.TakeProfit = &tpsl{
		Price:       price,
		TimeInForce: timeInForce,
		GtdTime:     gtdTime,
	}
}

// WithStopLoss specifies the details of a Stop Loss Order
// to be created on behalf of a client. This may happen when an Order
// is filled that opens a Trade requiring a Stop Loss, or when a Trade’s
// dependent Stop Loss Order is modified directly through the Trade.
func (rtpsl *requestTPSL) WithStopLoss(price float64, timeInForce, gtdTime string) {
	rtpsl.StopLoss = &tpsl{
		Price:       price,
		TimeInForce: timeInForce,
		GtdTime:     gtdTime,
	}
}

// WithTrailingStopLoss specifies the details of a Trailing Stop Loss Order
// to be created on behalf of a client. This may happen when an Order is
// filled that opens a Trade requiring a Trailing Stop Loss, or when a Trade’s
// dependent Trailing Stop Loss Order is modified directly through the Trade.
func (rtpsl *requestTPSL) WithTrailingStopLoss(price float64, timeInForce, gtdTime string) {
	rtpsl.TrailingStopLoss = &tpsl{
		Price:       price,
		TimeInForce: timeInForce,
		GtdTime:     gtdTime,
	}
}

// WithGuaranteedStopLoss specifies the details of a Guaranteed Stop Loss Order
// to be created on behalf of a client. This may happen when an Order
// is filled that opens a Trade requiring a Guaranteed Stop Loss,
// or when a Trade’s dependent Guaranteed Stop Loss Order is
// modified directly through the Trade.
func (rtpsl *requestTPSL) WithGuaranteedStopLoss(price, distance float64, timeInForce, gtdTime string) {
	rtpsl.GuaranteedStopLoss = &tpsl{
		Price:       price,
		TimeInForce: timeInForce,
		GtdTime:     gtdTime,
		Distance:    distance,
	}
}

type tradeQuery struct {
	Ids        string `json:"ids,omitempty"`
	State      string `json:"state,omitempty"`
	Instrument string `json:"instrument,omitempty"`
	Count      int    `json:"count,omitempty,string"`
	BeforeID   string `json:"beforeID,omitempty"`
}

type tradeOpts func(*tradeQuery)

type tradeFunc struct{}

func newTradeQuery(opts ...tradeOpts) *tradeQuery {
	q := &tradeQuery{}
	for _, query := range opts {
		query(q)
	}
	return q
}

func (tf *tradeFunc) WithInstrument(instrument string) tradeOpts {
	return func(tq *tradeQuery) { tq.Instrument = instrument }
}

func (tf *tradeFunc) WithBeforeID(beforeID string) tradeOpts {
	return func(tq *tradeQuery) { tq.BeforeID = beforeID }
}

func (tf *tradeFunc) WithCount(count int) tradeOpts {
	return func(tq *tradeQuery) { tq.Count = count }
}

func (tf *tradeFunc) WithIds(ids []string) tradeOpts {
	return func(tq *tradeQuery) {
		tq.Ids = strings.Join(ids, ",")
	}
}

func (tf *tradeFunc) WithState(state string) tradeOpts {
	return func(tq *tradeQuery) { tq.State = state }
}
