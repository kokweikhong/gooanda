package addr

const (
	streamhost = "https://stream-fxtrade.oanda.com"
	livehost   = "https://api-fxtrade.oanda.com"

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

	// Pricing endpoints
	PrCandleLatest       = livehost + "/v3/accounts/%v/candles/latest"
	PrPricingInfo        = livehost + "/v3/accounts/%v/pricing"
	PrPricingStream      = streamhost + "/v3/accounts/%v/pricing/stream"
	PrInstrumentCandlees = livehost + "/v3/accounts/%v/instruments/instrument/candles"
)
