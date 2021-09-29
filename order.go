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
	Order struct {
		Type             string      `json:"type"`
		Instrument       string      `json:"instrument,omitempty"`
		Units            float64     `json:"units,string"`
		TimeInForce      string      `json:"timeInForce"`
		Price            float64     `json:"price,omitempty,string"`
		PriceBound       float64     `json:"priceBound,omitempty,string"`
		PositionFill     string      `json:"positionFill,omitempty"`
		TakeProfitOnFill *detailSLTP `json:"takeProfitOnFill,omitempty"`
		StopLossOnFill   *detailSLTP `json:"stopLossOnFill,omitempty"`
		TriggerCondition string      `json:"triggerCondition,omitempty"`
		TradeID          string      `json:"tradeID,omitempty,string"`
		ClientTradeID    string      `json:"clientTradeID,omitempty,string"`
	} `json:"order"`
}

type detailSLTP struct {
	Price       float64 `json:"price,string"`
	TimeInForce string  `json:"timeInForce"`
	GtdTime     string  `json:"gtdTime"`
}

// type orderResponse struct {
//     OrderCreateTransaction *transaction `json:"orderCreateTransaction,omitempty"`
//     orderFillTransaction : (OrderFillTransaction),
//     orderCancelTransaction : (OrderCancelTransaction),
//     orderReissueTransaction : (Transaction),
//     orderReissueRejectTransaction : (Transaction),
//     relatedTransactionIDs : (Array[TransactionID]),
//     lastTransactionID : (TransactionID)
// }

// type transaction struct {

// }
// MarketOrderRequest specifies the parameters that may be set when creating a Market Order.
func (od *order) MarketOrderRequest(live bool, accountID, instrument string, units float64, opts ...configOpts) {
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.MARKET
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithInstrument(instrument),
		od.Config.WithUnits(units))
	fmt.Println(conf.Order.Instrument)
	data, err := convertConfig(conf)
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
	data, err := convertConfig(conf)
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
}

// StopOrderRequest specifies the parameters that may be set when creating a Stop Order.
func (od *order) StopOrderRequest(live bool, accountID, instrument string, price, units float64, opts ...configOpts) {
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.STOP
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithInstrument(instrument),
		od.Config.WithUnits(units),
		od.Config.WithPrice(price))
	data, err := convertConfig(conf)
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
}

// MarketIfTouchedOrderRequest specifies the parameters that may be set when creating a Market-if-Touched Order.
func (od *order) MarketIfTouchedOrderRequest(live bool, accountID, instrument string, price, units float64, opts ...configOpts) {
	conf := &configOrder{}
	conf.Order.Type = kw.ORDERTYPE.MARKET_IF_TOUCHED
	conf.defaultConfig()
	conf.extendOrderConfig(opts...)
	conf.extendOrderConfig(
		od.Config.WithInstrument(instrument),
		od.Config.WithUnits(units),
		od.Config.WithPrice(price))
	data, err := convertConfig(conf)
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
}

// TakeProfitOrderRequest specifies the parameters that may be set when creating a Take Profit Order.
func (od *order) TakeProfitOrderRequest() {}

// StopLossOrderRequest specifies the parameters that may be set when creating a Stop Loss Order. Only one of the price and distance fields may be specified.
func (od *order) StopLossOrderRequest() {}

// GuaranteedStopLossOrderRequest specifies the parameters that may be set when creating a Guaranteed Stop Loss Order. Only one of the price and distance fields may be specified.
func (od *order) GuaranteedStopLossOrderRequest() {}

// TrailingStopLossOrderRequest specifies the parameters that may be set when creating a Trailing Stop Loss Order.
func (od *order) TrailingStopLossOrderRequest() {}

// convertConfig is to remove fields which not required based on what order request.
func convertConfig(config *configOrder) ([]byte, error) {
	switch strings.ToLower(config.Order.Type) {
	case kw.ORDERTYPE.MARKET:
		config.Order.TriggerCondition = ""
	case kw.ORDERTYPE.LIMIT:
		config.Order.PriceBound = 0
	}
	fmt.Println(config.Order.Type)
	data, err := json.Marshal(config)
	if err != nil {
		return data, fmt.Errorf("failed to marshal config to json, %v", err)
	}
	return data, nil
}

type configOpts func(*configOrder)

// extendOrderConfig is to add fields to current configuration after default config called.
func (cf *configOrder) extendOrderConfig(config ...configOpts) {
	for _, conf := range config {
		conf(cf)
	}
}

// defaultConfig is create configuration based on what order request.
func (cf *configOrder) defaultConfig() {
	switch cf.Order.Type {
	case kw.ORDERTYPE.MARKET: // Market order request
		cf.Order.TimeInForce = kw.TIMEINFORCE.FOK
		cf.Order.PositionFill = kw.POSITIONFILL.DEFAULT
	case kw.ORDERTYPE.LIMIT, kw.ORDERTYPE.STOP,
		kw.ORDERTYPE.MARKET_IF_TOUCHED:
		cf.Order.TimeInForce = kw.TIMEINFORCE.GTC
		cf.Order.PositionFill = kw.POSITIONFILL.DEFAULT
		cf.Order.TriggerCondition = kw.TRIGGERCONDITION.DEFAULT
	}
}

type orderConfigFunc struct{}

func (*orderConfigFunc) WithPrice(price float64) configOpts {
	return func(co *configOrder) { co.Order.Price = price }
}

func (*orderConfigFunc) WithUnits(units float64) configOpts {
	return func(co *configOrder) { co.Order.Units = units }
}

func (*orderConfigFunc) WithTriggerCondition(triggerCondition string) configOpts {
	return func(co *configOrder) { co.Order.TriggerCondition = triggerCondition }
}

func (*orderConfigFunc) WithStopLossOnFill(gtdTime, timeInForce string, price float64) configOpts {
	return func(co *configOrder) {
		co.Order.StopLossOnFill.GtdTime = gtdTime
		co.Order.StopLossOnFill.TimeInForce = timeInForce
		co.Order.StopLossOnFill.Price = price
	}
}

func (*orderConfigFunc) WithTakeProfitOnFill(gtdTime, timeInForce string, price float64) configOpts {
	return func(co *configOrder) {
		co.Order.TakeProfitOnFill.GtdTime = gtdTime
		co.Order.TakeProfitOnFill.TimeInForce = timeInForce
		co.Order.TakeProfitOnFill.Price = price
	}
}

func (*orderConfigFunc) WithPositionFill(positionFill string) configOpts {
	return func(co *configOrder) { co.Order.PositionFill = positionFill }
}

func (*orderConfigFunc) WithInstrument(instrument string) configOpts {
	return func(co *configOrder) { co.Order.Instrument = instrument }
}

func (*orderConfigFunc) WithTimeInForce(timeInForce string) configOpts {
	return func(co *configOrder) {
		co.Order.TimeInForce = timeInForce
	}
}

func (*orderConfigFunc) WithPriceBound(priceBound float64) configOpts {
	return func(co *configOrder) {
		co.Order.PriceBound = priceBound
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
}
