package addr

const (
	livehost          = "https://api-fxtrade.oanda.com"
	Accounts          = livehost + "/v3/accounts"
	AccountsById      = livehost + "/v3/accounts/%v"
	AccountSummary    = livehost + "/v3/accounts/%v/summary"
	AccountInstrument = livehost + "/v3/accounts/%v/instruments"

	InstrumentCandles      = livehost + "/v3/instruments/%v/candles"
	InstrumentOrderBook    = livehost + "/3/instruments/%v/orderBook"
	InstrumentPositionBook = livehost + "/v3/instruments/%v/positionBook"
)
