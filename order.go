package gooanda

type order struct { // {{{
	OrderType string `json:"type"`
	// The Market Order’s Instrument.
	Instrument string `json:"instrument"`
	// The quantity requested to be filled by the Market Order.
	// A positive number of units results in a long Order
	// Negative number of units results in a short Order.
	Units string `json:"units"`
	// The time-in-force requested for the Market Order. Restricted to FOK or IOC for a MarketOrder.
	TimeInForce string `json:"timeInForce"`
	// The worst price that the client is willing to have the Market Order filled at.
	PriceBound string `json:"priceBound,omitempty"`
	// Specification of how Positions in the Account are modified when the Order is filled.
	positionFill string
	// The client extensions to add to the Order. Do not set, modify, or delete clientExtensions if your account is associated with MT4.
	clientExtensions string
	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	takeProfitOnFill string
	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	stopLossOnFill string
	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Guaranteed Stop Loss, or when a
	// Trade’s dependent Guaranteed Stop Loss Order is modified directly through
	// the Trade.
	guaranteedStopLossOnFill string
	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	trailingStopLossOnFill string
	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	tradeClientExtensions string
} // }}}

func PostOrderCreate() *order {
	return nil
}

type order2 connection

func (o *order2) GetOrderList() {

}

func (o *order) MarketOrder() {}

func GetOrderList() {}

func GetOrderPendingList() {}

func GetOrderDetails() {}

func PutOrderReplace() {}

func PutOrderCancel() {}

func PutOrderUpdateClientExt() {}
