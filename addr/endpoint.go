package addr

const (
	livehost = "https://api-fxtrade.oanda.com"

	// Account endpoints
	Accounts          = livehost + "/v3/accounts"
	AccountsById      = livehost + "/v3/accounts/%v"
	AccountSummary    = livehost + "/v3/accounts/%v/summary"
	AccountInstrument = livehost + "/v3/accounts/%v/instruments"

	// Instrument endpoints
	InstrumentCandles      = livehost + "/v3/instruments/%v/candles"
	InstrumentOrderBook    = livehost + "/3/instruments/%v/orderBook"
	InstrumentPositionBook = livehost + "/v3/instruments/%v/positionBook"

	// Order endpoints
	Orders        = livehost + "/v3/accounts/%v/orders"
	PendingOrders = livehost + "/v3/accounts/%v/pendingOrders"
)
