package gooanda

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kokweikhong/gooanda/endpoint"
	"github.com/kokweikhong/gooanda/kw"
)

type order struct {
	connection
	Config *orderConfigFunc
	Query  *orderQueryFunc
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

// -------------ORDER QUERY SECTION--------------------
// {{{
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

type orderQueryFunc struct{}

// WithCount is the maximum number of Orders to return [default=50, maximum=500]
func (*orderQueryFunc) WithCount(count int) orderOpts {
	return func(oq *orderQuery) {
		if count < 0 || count > 5000 {
			oq.Count = "500"
			return
		} else {
			oq.Count = strconv.Itoa(count)
		}
	}
}

// WithState is the state to filter the requested Orders by.
func (*orderQueryFunc) WithState(orderState string) orderOpts {
	return func(oq *orderQuery) {
		oq.State = orderState
	}
}

// WithIstrument is the instrument to filter the requested orders by.
func (*orderQueryFunc) WithInstrument(instrument string) orderOpts {
	return func(oq *orderQuery) {
		oq.Instrument = instrument
	}
}

// WithListId is List of OrderID (csv), List of Order IDs to retrieve
func (*orderQueryFunc) WithListID(ids []string) orderOpts {
	return func(oq *orderQuery) {
		oq.Ids = strings.Join(ids, ",")
	}
}

// WithID is the maximum Order ID to return. If not provided the most recent Orders in the Account are returned
func (*orderQueryFunc) WithID(id string) orderOpts {
	return func(oq *orderQuery) {
		oq.BeforeID = id
	}
} // }}}

// --------------CREATE ORDER CONFIGURATION SECTION-----------------
// {{{
type configOrder struct {
	Order struct {
		Type                     string      `json:"type"`
		Instrument               string      `json:"instrument,omitempty"`
		Units                    float64     `json:"units,string,omitempty"`
		TimeInForce              string      `json:"timeInForce"`
		Price                    float64     `json:"price,omitempty,string"`
		PriceBound               float64     `json:"priceBound,omitempty,string"`
		PositionFill             string      `json:"positionFill,omitempty"`
		TakeProfitOnFill         *detailSLTP `json:"takeProfitOnFill,omitempty"`
		StopLossOnFill           *detailSLTP `json:"stopLossOnFill,omitempty"`
		TriggerCondition         string      `json:"triggerCondition,omitempty"`
		TradeID                  string      `json:"tradeID,omitempty,string"`
		ClientTradeID            string      `json:"clientTradeID,omitempty,string"`
		Distance                 float64     `json:"distance,omitempty,string"`
		TrailingStopLossOnFill   *tSLOF      `json:"trailingStopLossOnFill,omitempty"`
		GuaranteedStopLossOnFill *gSLOF      `json:"guaranteedStopLossOnFill,omitempty"`
	} `json:"order"`
}

// TrailingStopLossOnFill
type tSLOF struct {
	Distance    float64 `json:"distance,omitempty,string"`
	GtdTime     string  `json:"gtdTime,omitempty"`
	TimeInForce string  `json:"timeInForce"`
}

// GuaranteedStopLossOnFill
type gSLOF struct {
	Price       float64 `json:"price,omitempty,string"`
	Distance    float64 `json:"distance,omitempty,string"`
	GtdTime     string  `json:"gtdTime,omitempty"`
	TimeInForce string  `json:"timeInForce"`
}

type detailSLTP struct {
	Price       float64 `json:"price,string"`
	TimeInForce string  `json:"timeInForce"`
	GtdTime     string  `json:"gtdTime,omitempty"`
}

type configOpts func(*configOrder)

// extendOrderConfig is to add fields to current configuration after default config called.
func (cf *configOrder) extendOrderConfig(config ...configOpts) {
	for _, conf := range config {
		conf(cf)
	}
}

type orderConfigFunc struct{}

// WithGuaranteedStopLossOnFill specifies the details of a Guaranteed Stop Loss
// Order to be created on behalf of a client. This may happen when an Order
// is filled that opens a Trade requiring a Guaranteed Stop Loss, or when a
// Trade’s dependent Guaranteed Stop Loss Order is modified directly through
// the Trade.
func (*orderConfigFunc) WithGuaranteedStopLossOnFill(price, distance float64, timeInForce, gtdTime string) configOpts {
	return func(co *configOrder) {
		co.Order.GuaranteedStopLossOnFill = &gSLOF{}
		co.Order.GuaranteedStopLossOnFill.Price = price
		co.Order.GuaranteedStopLossOnFill.Distance = distance
		co.Order.GuaranteedStopLossOnFill.GtdTime = gtdTime
		switch timeInForce {
		case kw.TIMEINFORCE.GFD, kw.TIMEINFORCE.GTC, kw.TIMEINFORCE.GFD:
			co.Order.GuaranteedStopLossOnFill.TimeInForce = timeInForce
		default:
			co.Order.GuaranteedStopLossOnFill.TimeInForce = kw.TIMEINFORCE.GTC
		}
	}
}

// WithTrailingStopLossOnFill specifies the details of a Trailing Stop Loss
// Order to be created on behalf of a client. This may happen when an Order
// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
// Trade’s dependent Trailing Stop Loss Order is modified directly through
// the Trade.
func (*orderConfigFunc) WithTrailingStopLossOnFill(distance float64, timeInForce, gtdTime string) configOpts {
	return func(co *configOrder) {
		co.Order.TrailingStopLossOnFill = &tSLOF{}
		co.Order.TrailingStopLossOnFill.Distance = distance
		co.Order.TrailingStopLossOnFill.GtdTime = gtdTime
		switch timeInForce {
		case kw.TIMEINFORCE.GFD, kw.TIMEINFORCE.GTC, kw.TIMEINFORCE.GFD:
			co.Order.TrailingStopLossOnFill.TimeInForce = timeInForce
		default:
			co.Order.TrailingStopLossOnFill.TimeInForce = kw.TIMEINFORCE.GTC
		}
	}
}

func (*orderConfigFunc) WithDistance(distance float64) configOpts {
	return func(co *configOrder) { co.Order.Distance = distance }
}

func (*orderConfigFunc) WithClientTradeID(clientTradeID string) configOpts {
	return func(co *configOrder) { co.Order.ClientTradeID = clientTradeID }
}

func (*orderConfigFunc) WithTradeID(tradeID string) configOpts {
	return func(co *configOrder) { co.Order.TradeID = tradeID }
}

func (*orderConfigFunc) WithPrice(price float64) configOpts {
	return func(co *configOrder) { co.Order.Price = price }
}

// WithUnits is the quantity requested to be filled by the Market Order. A positive
// number of units results in a long Order, and a negative number of units
// results in a short Order.
func (*orderConfigFunc) WithUnits(units float64) configOpts {
	return func(co *configOrder) { co.Order.Units = units }
}

func (*orderConfigFunc) WithTriggerCondition(triggerCondition string) configOpts {
	return func(co *configOrder) { co.Order.TriggerCondition = triggerCondition }
}

// WithStopLossOnFill specifies the details of a Stop Loss Order to be created
// on behalf of a client. This may happen when an Order is filled that opens
// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
// Order is modified directly through the Trade.
func (*orderConfigFunc) WithStopLossOnFill(gtdTime, timeInForce string, price float64) configOpts {
	return func(co *configOrder) {
		co.Order.StopLossOnFill = &detailSLTP{}
		co.Order.StopLossOnFill.GtdTime = gtdTime
		co.Order.StopLossOnFill.TimeInForce = timeInForce
		co.Order.StopLossOnFill.Price = price
	}
}

// WithTakeProfitOnFill specifies the details of a Take Profit Order to be
// created on behalf of a client. This may happen when an Order is filled
// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
// Take Profit Order is modified directly through the Trade.
func (*orderConfigFunc) WithTakeProfitOnFill(gtdTime, timeInForce string, price float64) configOpts {
	return func(co *configOrder) {
		co.Order.TakeProfitOnFill = &detailSLTP{}
		co.Order.TakeProfitOnFill.GtdTime = gtdTime
		co.Order.TakeProfitOnFill.TimeInForce = timeInForce
		co.Order.TakeProfitOnFill.Price = price
	}
}

func (*orderConfigFunc) WithPositionFill(positionFill string) configOpts {
	return func(co *configOrder) { co.Order.PositionFill = positionFill }
}

// WithInstrument is the Market Order’s Instrument.
func (*orderConfigFunc) WithInstrument(instrument string) configOpts {
	return func(co *configOrder) { co.Order.Instrument = instrument }
}

// WithTimeInForce is the time-in-force requested for the Market Order. Restricted to FOK or IOC for a MarketOrder.
func (*orderConfigFunc) WithTimeInForce(timeInForce string) configOpts {
	return func(co *configOrder) {
		co.Order.TimeInForce = timeInForce
	}
}

// WithPriceBound is the worst price that the client is willing to have the Market Order filled at.
func (*orderConfigFunc) WithPriceBound(priceBound float64) configOpts {
	return func(co *configOrder) {
		co.Order.PriceBound = priceBound
	}
} // }}}

// --------------UTILS FUNCTION-----------------
// {{{
// defaultConfig is create configuration based on what order request.
func (cf *configOrder) defaultConfig() {
	switch cf.Order.Type {
	case kw.ORDERTYPE.MARKET:
		cf.Order.TimeInForce = kw.TIMEINFORCE.FOK
		cf.Order.PositionFill = kw.POSITIONFILL.DEFAULT
	case kw.ORDERTYPE.LIMIT, kw.ORDERTYPE.STOP,
		kw.ORDERTYPE.MARKET_IF_TOUCHED:
		cf.Order.TimeInForce = kw.TIMEINFORCE.GTC
		cf.Order.PositionFill = kw.POSITIONFILL.DEFAULT
		cf.Order.TriggerCondition = kw.TRIGGERCONDITION.DEFAULT
	case kw.ORDERTYPE.TAKE_PROFIT, kw.ORDERTYPE.STOP_LOSS,
		kw.ORDERTYPE.GUARANTEED_STOP_LOSS,
		kw.ORDERTYPE.TRAILING_STOP_LOSS:
		cf.Order.TriggerCondition = kw.TRIGGERCONDITION.DEFAULT
		cf.Order.TimeInForce = kw.TIMEINFORCE.GTC
	}
}

// convertConfig is to remove fields which not required based on what order request.
func (cf *configOrder) convertConfig() ([]byte, error) {
	switch cf.Order.Type {
	case kw.ORDERTYPE.MARKET:
		cf.Order.Price = 0
		cf.Order.TriggerCondition = ""
		cf.Order.ClientTradeID = ""
		cf.Order.TradeID = ""
		cf.Order.Distance = 0
	case kw.ORDERTYPE.LIMIT:
		cf.Order.PriceBound = 0
		cf.Order.ClientTradeID = ""
		cf.Order.TradeID = ""
		cf.Order.Distance = 0
	case kw.ORDERTYPE.MARKET_IF_TOUCHED, kw.ORDERTYPE.STOP:
		cf.Order.ClientTradeID = ""
		cf.Order.TradeID = ""
		cf.Order.Distance = 0
	case kw.ORDERTYPE.TAKE_PROFIT:
		cf.Order.Instrument = ""
		cf.Order.Units = 0
		cf.Order.PriceBound = 0
		cf.Order.PositionFill = ""
		cf.Order.TakeProfitOnFill = nil
		cf.Order.StopLossOnFill = nil
		cf.Order.Distance = 0
		cf.Order.GuaranteedStopLossOnFill = nil
		cf.Order.TrailingStopLossOnFill = nil
	case kw.ORDERTYPE.STOP_LOSS:
		cf.Order.Instrument = ""
		cf.Order.Units = 0
		cf.Order.PriceBound = 0
		cf.Order.PositionFill = ""
		cf.Order.TakeProfitOnFill = nil
		cf.Order.StopLossOnFill = nil
		cf.Order.GuaranteedStopLossOnFill = nil
		cf.Order.TrailingStopLossOnFill = nil
	case kw.ORDERTYPE.GUARANTEED_STOP_LOSS, kw.ORDERTYPE.TRAILING_STOP_LOSS:
		cf.Order.Instrument = ""
		cf.Order.Units = 0
		cf.Order.PriceBound = 0
		cf.Order.PositionFill = ""
		cf.Order.TakeProfitOnFill = nil
		cf.Order.StopLossOnFill = nil
		cf.Order.GuaranteedStopLossOnFill = nil
		cf.Order.TrailingStopLossOnFill = nil
	}
	data, err := json.Marshal(cf)
	if err != nil {
		return data, fmt.Errorf("failed to marshal config %T to json, %v", cf, err)
	}
	return data, nil
} // }}}

// ----------------  ORDER MAIN FUNCTION-------------------------

func (od *order) GetOrderList(live bool, accountID string, querys ...orderOpts) (string, error) {
	q := newOrderQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Order.Orders)
	ep = fmt.Sprintf(ep, accountID)
	url, err := urlAddQuery(ep, q)
	if err != nil {
		return "", err
	}
	od.endpoint = url
	od.method = http.MethodGet
	resp, err := od.connect()
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (od *order) GetPendingOrders(live bool, accountID string) (string, error) {
	ep := endpoint.GetEndpoint(live, endpoint.Order.PendingOrder)
	od.endpoint = fmt.Sprintf(ep, accountID)
	od.method = http.MethodGet
	resp, err := od.connect()
	if err != nil {
		return string(resp), err
	}
	return string(resp), err
}

func (od *order) GetOrderDetails(live bool, accountID, tradeID string) (string, error) {
	ep := endpoint.GetEndpoint(live, endpoint.Order.OrderDetails)
	od.endpoint = fmt.Sprintf(ep, accountID, tradeID)
	od.method = http.MethodGet
	resp, err := od.connect()
	if err != nil {
		return string(resp), err
	}
	return string(resp), nil

}

func PutOrderReplace() {}

func PutOrderCancel() {}

func PutOrderUpdateClientExt() {}

// MarketOrderRequest specifies the parameters that may be set when creating a Market Order.
func (od *order) MarketOrderRequest(live bool, accountID, instrument string, units float64, opts ...configOpts) (string, error) { // {{{
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.MARKET
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithInstrument(instrument),
		od.Config.WithUnits(units))
	data, err := conf.convertConfig()
	if err != nil {
		return "", err
	}
	od.data = data
	od.method = http.MethodPost
	ep := endpoint.GetEndpoint(live, endpoint.Order.Orders)
	od.endpoint = fmt.Sprintf(ep, accountID)
	resp, err := od.connect()
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// LimitOrderRequest specifies the parameters that may be set when creating a Limit Order.
func (od *order) LimitOrderRequest(live bool, accountID, instrument string, price, units float64, opts ...configOpts) {
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.LIMIT
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithInstrument(instrument),
		od.Config.WithUnits(units),
		od.Config.WithPrice(price))
	data, err := conf.convertConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("order created with below configuration:\n%v\n", string(data))
	od.data = data
	od.method = http.MethodPost
	ep := endpoint.GetEndpoint(live, endpoint.Order.Orders)
	od.endpoint = fmt.Sprintf(ep, accountID)
	resp, err := od.connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp))
} // }}}

// StopOrderRequest specifies the parameters that may be set when creating a Stop Order.
func (od *order) StopOrderRequest(live bool, accountID, instrument string, price, units float64, opts ...configOpts) (string, error) { // {{{
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.STOP
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithInstrument(instrument),
		od.Config.WithUnits(units),
		od.Config.WithPrice(price))
	data, err := conf.convertConfig()
	if err != nil {
		return "", err
	}
	od.data = data
	od.method = http.MethodPost
	ep := endpoint.GetEndpoint(live, endpoint.Order.Orders)
	od.endpoint = fmt.Sprintf(ep, accountID)
	resp, err := od.connect()
	if err != nil {
		return "", err
	}
	return string(resp), nil
} // }}}

// MarketIfTouchedOrderRequest specifies the parameters that may be set when creating a Market-if-Touched Order.
func (od *order) MarketIfTouchedOrderRequest(live bool, accountID, instrument string, price, units float64, opts ...configOpts) (string, error) { // {{{
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.MARKET_IF_TOUCHED
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithInstrument(instrument),
		od.Config.WithUnits(units),
		od.Config.WithPrice(price))
	data, err := conf.convertConfig()
	if err != nil {
		return "", err
	}
	od.data = data
	od.method = http.MethodPost
	ep := endpoint.GetEndpoint(live, endpoint.Order.Orders)
	od.endpoint = fmt.Sprintf(ep, accountID)
	resp, err := od.connect()
	if err != nil {
		return string(resp), err
	}
	return string(resp), nil
} // }}}

// TakeProfitOrderRequest specifies the parameters that may be
// set when creating a Take Profit Order.
func (od *order) TakeProfitOrderRequest(live bool, accountID, tradeID string, price float64, opts ...configOpts) (string, error) { // {{{
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.TAKE_PROFIT
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithTradeID(tradeID),
		od.Config.WithPrice(price))
	data, err := conf.convertConfig()
	if err != nil {
		return "", err
	}
	fmt.Printf("order created with below configuration:\n%v\n", string(data))
	od.data = data
	od.method = http.MethodPost
	ep := endpoint.GetEndpoint(live, endpoint.Order.Orders)
	od.endpoint = fmt.Sprintf(ep, accountID)
	resp, err := od.connect()
	if err != nil {
		return string(resp), err
	}
	return string(resp), nil
} // }}}

// StopLossOrderRequest specifies the parameters that may be set
// when creating a Stop Loss Order. Only one of the price and
// distance fields may be specified.
func (od *order) StopLossOrderRequest(live bool, accountID, tradeID string, price float64, opts ...configOpts) (string, error) { // {{{
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.STOP_LOSS
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithTradeID(tradeID),
		od.Config.WithPrice(price))
	data, err := conf.convertConfig()
	if err != nil {
		return "", err
	}
	od.data = data
	od.method = http.MethodPost
	ep := endpoint.GetEndpoint(live, endpoint.Order.Orders)
	od.endpoint = fmt.Sprintf(ep, accountID)
	resp, err := od.connect()
	if err != nil {
		return string(resp), err
	}
	return string(resp), nil
} // }}}

// GuaranteedStopLossOrderRequest specifies the parameters that
// may be set when creating a Guaranteed Stop Loss Order.
// Only one of the price and distance fields may be specified.
func (od *order) GuaranteedStopLossOrderRequest(live bool, accountID, tradeID string, price float64, opts ...configOpts) (string, error) { // {{{
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.GUARANTEED_STOP_LOSS
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithTradeID(tradeID),
		od.Config.WithPrice(price))
	data, err := conf.convertConfig()
	if err != nil {
		return "", err
	}
	od.data = data
	od.method = http.MethodPost
	ep := endpoint.GetEndpoint(live, endpoint.Order.Orders)
	od.endpoint = fmt.Sprintf(ep, accountID)
	resp, err := od.connect()
	if err != nil {
		return string(resp), err
	}
	return string(resp), nil
} // }}}

// TrailingStopLossOrderRequest specifies the parameters that
// may be set when creating a Trailing Stop Loss Order.
func (od *order) TrailingStopLossOrderRequest(live bool, accountID, tradeID string, distance float64, opts ...configOpts) (string, error) { // {{{
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.TRAILING_STOP_LOSS
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithTradeID(tradeID),
		od.Config.WithPrice(distance))
	data, err := conf.convertConfig()
	if err != nil {
		return "", err
	}
	od.data = data
	od.method = http.MethodPost
	ep := endpoint.GetEndpoint(live, endpoint.Order.Orders)
	od.endpoint = fmt.Sprintf(ep, accountID)
	resp, err := od.connect()
	if err != nil {
		return string(resp), err
	}
	return string(resp), nil
} // }}}
