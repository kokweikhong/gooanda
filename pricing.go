package gooanda

func GetCandlesLatestByAccount() {}

func GetPricingInformationByAccount() {}

func GetStreamingPriceByAccount() {}

func GetCandleStickByAccount() {}

type pricingQuery struct {
	// List of CandleSpecification (csv)
	// List of candle specifications to get pricing for. [required]
	CandleSpecifications string `json:"candleSpecifications"`
	// DecimalNumber	The number of units used to calculate the volume-weighted average bid and ask prices in the returned candles. [default=1]
	units string
	// boolean	A flag that controls whether the candlestick is “smoothed” or not. A smoothed candlestick uses the previous candle’s close price as its open price, while an unsmoothed candlestick uses the first price from its time range as its open price. [default=False]
	smooth string
	// integer	The hour of the day (in the specified timezone) to use for granularities that have daily alignments. [default=17, minimum=0, maximum=23]
	dailyAlignment string
	// The timezone to use for the dailyAlignment parameter. Candlesticks with daily alignment will be aligned to the dailyAlignment hour within the alignmentTimezone. Note that the returned times will still be represented in UTC. [default=America/New_York]
	alignmentTimezone string
	// WeeklyAlignment	The day of the week used for granularities that have weekly alignment. [default=Friday]
	weeklyAlignment string
	// List of InstrumentName (csv)	List of Instruments to get pricing for. [required]
	instruments string
	// DateTime	Date/Time filter to apply to the response. Only prices and home conversions (if requested) with a time later than this filter (i.e. the price has changed after the since time) will be provided, and are filtered independently.
	since string
	// boolean	Flag that enables the inclusion of the homeConversions field in the returned response. An entry will be returned for each currency in the set of all base and quote currencies present in the requested instruments list. [default=False]
	includeHomeConversions string
	// boolean	Flag that enables/disables the sending of a pricing snapshot when initially connecting to the stream. [default=True]
	snapshot string
	// PricingComponent	The Price component(s) to get candlestick data for. [default=M]
	price string
	// CandlestickGranularity	The granularity of the candlesticks to fetch [default=S5]
	granularity string
	// integer	The number of candlesticks to return in the response. Count should not be specified if both the start and end parameters are provided, as the time range combined with the granularity will determine the number of candlesticks to return. [default=500, maximum=5000]
	count string
	// DateTime	The start of the time range to fetch candlesticks for.
	from string
	// DateTime	The end of the time range to fetch candlesticks for.
	to string
	// boolean	A flag that controls whether the candlestick that is covered by the from time should be included in the results. This flag enables clients to use the timestamp of the last completed candlestick received to poll for future candlesticks but avoid receiving the previous candlestick repeatedly. [default=True]
	includeFirst string
}
