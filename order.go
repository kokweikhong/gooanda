package gooanda

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kokweikhong/gooanda/kw"
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

type order struct {
	connection
	Query  *orderFunc
	Config *configFunc
}

func NewOrderConnection(token string) *order {
	conn := &order{}
	conn.token = token
	return conn
}

func (od *order) connect() ([]byte, error) {
	con := &connection{od.endpoint, od.method, od.token, od.data}
	resp, err := con.connect()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (od *order) GetOrderList(live bool, accountID string, querys ...orderOpts) ([]byte, error) {
	q := newOrderQuery(querys...)
	fmt.Println(q)
	// var url string
	// if live {
	// 	url = fmt.Sprintf(addr.OdOrderList, addr.LiveHost, accountID)
	// } else if !live {
	// 	url = fmt.Sprintf(addr.OdOrderList, addr.PracticeHost, accountID)
	// }
	// u, err := urlAddQuery(url, q)
	// if err != nil {
	// 	return nil, err
	// }
	// od.endpoint = u
	// od.method = http.MethodGet
	// resp, err := od.connect()
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(string(resp))
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

type configOrder struct {
	Type             string               `json:"type"`
	Instrument       string               `json:"instrument"`
	Units            float64              `json:"units,string"`
	TimeInForce      kw.TimeInForce       `json:"timeInForce,string"`
	PriceBound       float64              `json:"priceBound,string"`
	PositionFill     kw.OrderPositionFill `json:"positionFill,string"`
	TakeProfitOnFill sLTpDetails          `json:"takeProfitOnFill,omitempty"`
}

type sLTpDetails struct {
	Price       float64        `json:"price,string"`
	TimeInForce kw.TimeInForce `json:"timeInForce,string"`
}

func (od *order) MarketOrderRequest(instrument string, units float64, opts ...configOpts) {
	conf := newOrderConfig(opts...)
	conf.Instrument = instrument
	conf.Units = units
	data, err := json.Marshal(conf)
	if err != nil {
		log.Fatal(err)
	}
	od.data = data
	od.method = http.MethodPost
	// od.endpoint = fmt.Sprintf(addr.LiveHost + addr.)
}

type configOpts func(*configOrder)

func newOrderConfig(config ...configOpts) *configOrder {
	c := &configOrder{}
	for _, conf := range config {
		conf(c)
	}
	return c
}

type configFunc struct{}

func (*configFunc) WithTimeInForce(timeInForce kw.TimeInForce) configOpts {
	return func(co *configOrder) {
		co.TimeInForce = timeInForce
	}
}

func (*configFunc) WithPriceBound(priceBound float64) configOpts {
	return func(co *configOrder) {
		co.PriceBound = priceBound
	}
}

type orderQuery struct {
	Ids        string `json:"ids,omitempty"`
	State      string `json:"state,omitempty"`
	Instrument string `json:"instrument,omitempty"`
	Count      string `json:"count,omitempty"`
	BeforeID   string `json:"beforeID,omitempty"`
}

type orderOpts func(*orderQuery)

func newOrderQuery(querys ...orderOpts) *orderQuery {
	q := &orderQuery{}
	for _, query := range querys {
		query(q)
	}
	return q
}

type orderFunc struct{}

// integer	The maximum number of Orders to return [default=50, maximum=500]
func (*orderFunc) WithCount(count int) orderOpts {
	return func(oq *orderQuery) {
		if count < 0 || count > 5000 {
			oq.Count = "500"
			return
		} else {
			oq.Count = strconv.Itoa(count)
		}
	}
}

// The state to filter the requested Orders by [default=PENDING]
func (*orderFunc) WithState(state kw.OrderState) orderOpts {
	return func(oq *orderQuery) {
		oq.State = string(state)
	}
}

// The instrument to filter the requested orders by
func (*orderFunc) WithInstrument(instrument string) orderOpts {
	return func(oq *orderQuery) {
		oq.Instrument = instrument
	}
}

// List of OrderID (csv)	List of Order IDs to retrieve
func (*orderFunc) WithListID(ids []string) orderOpts {
	return func(oq *orderQuery) {
		oq.Ids = strings.Join(ids, ",")
	}
}

// The maximum Order ID to return. If not provided the most recent Orders in the Account are returned
func (*orderFunc) WithID(id string) orderOpts {
	return func(oq *orderQuery) {
		oq.BeforeID = id
	}
}
