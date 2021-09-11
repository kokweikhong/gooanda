package gooanda

import (
	"fmt"
	"net/http"

	"github.com/kokweikhong/gooanda/addr"
	"github.com/kokweikhong/gooanda/opts"
)

// type orderC struct { // {{{
// 	OrderType string `json:"type"`
// 	// The Market Order’s Instrument.
// 	Instrument string `json:"instrument"`
// 	// The quantity requested to be filled by the Market Order.
// 	// A positive number of units results in a long Order
// 	// Negative number of units results in a short Order.
// 	Units string `json:"units"`
// 	// The time-in-force requested for the Market Order. Restricted to FOK or IOC for a MarketOrder.
// 	TimeInForce string `json:"timeInForce"`
// 	// The worst price that the client is willing to have the Market Order filled at.
// 	PriceBound string `json:"priceBound,omitempty"`
// 	// Specification of how Positions in the Account are modified when the Order is filled.
// 	positionFill string
// 	// The client extensions to add to the Order. Do not set, modify, or delete clientExtensions if your account is associated with MT4.
// 	clientExtensions string
// 	// TakeProfitDetails specifies the details of a Take Profit Order to be
// 	// created on behalf of a client. This may happen when an Order is filled
// 	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
// 	// Take Profit Order is modified directly through the Trade.
// 	takeProfitOnFill string
// 	// StopLossDetails specifies the details of a Stop Loss Order to be created
// 	// on behalf of a client. This may happen when an Order is filled that opens
// 	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
// 	// Order is modified directly through the Trade.
// 	stopLossOnFill string
// 	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss
// 	// Order to be created on behalf of a client. This may happen when an Order
// 	// is filled that opens a Trade requiring a Guaranteed Stop Loss, or when a
// 	// Trade’s dependent Guaranteed Stop Loss Order is modified directly through
// 	// the Trade.
// 	guaranteedStopLossOnFill string
// 	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
// 	// Order to be created on behalf of a client. This may happen when an Order
// 	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
// 	// Trade’s dependent Trailing Stop Loss Order is modified directly through
// 	// the Trade.
// 	trailingStopLossOnFill string
// 	// Client Extensions to add to the Trade created when the Order is filled
// 	// (if such a Trade is created). Do not set, modify, or delete
// 	// tradeClientExtensions if your account is associated with MT4.
// 	tradeClientExtensions string
// } // }}}

type order connection

func NewOrderConnection(token string) *order {
	return &order{token: token}
}

func (od *order) connect() ([]byte, error) {
	con := &connection{od.endpoint, od.method, od.token, od.data}
	resp, err := con.connect()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (od *order) GetOrderList(live bool, accountID string, querys ...opts.OrderOpts) ([]byte, error) {
	q := opts.NewOrderQuery(querys...)
	var url string
	if live {
		url = fmt.Sprintf(addr.OdOrderList, addr.LiveHost, accountID)
	} else if !live {
		url = fmt.Sprintf(addr.OdOrderList, addr.PracticeHost, accountID)
	}
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	od.endpoint = u
	od.method = http.MethodGet
	resp, err := od.connect()
	if err != nil {
		return nil, err
	}
	fmt.Println(string(resp))
	// var data = &PricingCandleLatest{}
	// if err = json.Unmarshal(resp, &data); err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func GetOrderPendingList() {}

func GetOrderDetails() {}

func PutOrderReplace() {}

func PutOrderCancel() {}

func PutOrderUpdateClientExt() {}

type RequestMarketOrder struct {
	Instrument   string  `json:"instrument"`
	Units        float64 `json:"units"`
	TimeInForce  string  `json:"timeInForce"`
	PositionFill string  `json:"positionFill"`
}

type createOrder struct{}

func (od *order) CreateOrder() *createOrder {
	return &createOrder{}
}

func (no *createOrder) MarketOrder() {

}
