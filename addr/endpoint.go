package addr

const (
	streamApi    = "https://stream-%v.oanda.com"
	restApi      = "https://api-%v.oanda.com"
	LiveHost     = "fxtrade"
	PracticeHost = "fxpractice"

	// https://api-fxpractice.oanda.com
	// Account endpoints
	Accounts          = restApi + "/v3/accounts"
	AccountsById      = restApi + "/v3/accounts/%v"
	AccountSummary    = restApi + "/v3/accounts/%v/summary"
	AccountInstrument = restApi + "/v3/accounts/%v/instruments"

	// Instrument endpoints
	InstrumentCandles      = restApi + "/v3/instruments/%v/candles"
	InstrumentOrderBook    = restApi + "/v3/instruments/%v/orderBook"
	InstrumentPositionBook = restApi + "/v3/instruments/%v/positionBook"

	// Order endpoints
	OdOrderList        = restApi + "/v3/accounts/%v/orders"
	OdPendingOrderList = restApi + "/v3/accounts/%v/pendingOrders"

	// Pricing endpoints
	PrCandleLatest      = restApi + "/v3/accounts/%v/candles/latest"
	PrPricingInfo       = restApi + "/v3/accounts/%v/pricing"
	PrPricingStream     = streamApi + "/v3/accounts/%v/pricing/stream"
	PrInstrumentCandles = restApi + "/v3/accounts/%v/instruments/%v/candles"
)
