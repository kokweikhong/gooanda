package opts

import (
	"strconv"

	"github.com/kokweikhong/gooanda/kw"
)

type pricingQuery struct {
	// List of CandleSpecification (csv)
	// List of candle specifications to get pricing for. [required]
	CandleSpecifications string `json:"candleSpecifications,omitempty"`
	// DecimalNumber	The number of units used to calculate the volume-weighted average bid and ask prices in the returned candles. [default=1]
	Units             string `json:"units,omitempty"`
	Smooth            string `json:"smooth,omitempty"`
	DailyAlignment    string `json:"dailyAlignment,omitempty"`
	AlignmentTimezone string `json:"alignmentTimezone,omitempty"`
	WeeklyAlignment   string `json:"weeklyAlignment,omitempty"`
	// List of InstrumentName (csv)	List of Instruments to get pricing for. [required]
	instruments string
	// DateTime	Date/Time filter to apply to the response. Only prices and home conversions (if requested) with a time later than this filter (i.e. the price has changed after the since time) will be provided, and are filtered independently.
	since                  string
	IncludeHomeConversions string `json:"includeHomeConversions,omitempty"`
	Snapshot               string `json:"snapshot,omitempty"`
	// PricingComponent	The Price component(s) to get candlestick data for. [default=M]
	price       string
	Granularity string `json:"granularity,omitempty"`
	Count       string `json:"count,omitempty"`
	// DateTime	The start of the time range to fetch candlesticks for.
	from string
	// DateTime	The end of the time range to fetch candlesticks for.
	to           string
	IncludeFirst string `json:"includeFirst,omitempty"`
}

type PricingOpts func(*pricingQuery)

func NewPricingQuery(querys ...PricingOpts) *pricingQuery {
	q := &pricingQuery{}
	for _, query := range querys {
		query(q)
	}
	return q
}

// The number of candlesticks to return in the response. Count should not be specified if both the start and end parameters are provided, as the time range combined with the granularity will determine the number of candlesticks to return. [default=500, maximum=5000]
func QPrCount(count int) PricingOpts {
	return func(pq *pricingQuery) {
		if count < 1 || count > 5000 {
			pq.Count = "500"
			return
		}
		pq.Count = strconv.Itoa(count)
	}
}

// A flag that controls whether the candlestick that is covered by the from time should be included in the results. This flag enables clients to use the timestamp of the last completed candlestick received to poll for future candlesticks but avoid receiving the previous candlestick repeatedly. [default=True]
func QPrWithoutIncludeFirst() PricingOpts {
	return func(pq *pricingQuery) { pq.IncludeFirst = "false" }
}

// CandlestickGranularity	The granularity of the candlesticks to fetch [default=S5]
func QPrGranularity(granularity kw.Granularity) PricingOpts {
	return func(pq *pricingQuery) { pq.Granularity = string(granularity) }
}

// Flag that enables/disables the sending of a pricing snapshot when initially connecting to the stream. [default=True]
func QPrWithoutSnapshot() PricingOpts {
	return func(pq *pricingQuery) { pq.Snapshot = "false" }
}

// Flag that enables the inclusion of the homeConversions field in the returned response. An entry will be returned for each currency in the set of all base and quote currencies present in the requested instruments list. [default=False]
func QPrWithIncludeHomeConversions() PricingOpts {
	return func(pq *pricingQuery) { pq.IncludeHomeConversions = "true" }
}

// WeeklyAlignment	The day of the week used for granularities that have weekly alignment. [default=Friday]
func QPrWeeklyAlignment(day kw.WeeklyAlignment) PricingOpts {
	return func(pq *pricingQuery) { pq.WeeklyAlignment = string(day) }
}

// The timezone to use for the dailyAlignment parameter. Candlesticks with daily alignment will be aligned to the dailyAlignment hour within the alignmentTimezone. Note that the returned times will still be represented in UTC. [default=America/New_York]
func QPrAlignmentTimezone(timezone string) PricingOpts {
	return func(pq *pricingQuery) { pq.AlignmentTimezone = timezone }
}

// The hour of the day (in the specified timezone) to use for granularities that have daily alignments. [default=17, minimum=0, maximum=23]
func QPrDailyAlignment(alignment int) PricingOpts {
	return func(pq *pricingQuery) {
		if alignment < 0 || alignment > 23 {
			pq.DailyAlignment = ""
			return
		}
		pq.DailyAlignment = strconv.Itoa(alignment)
	}
}

// A flag that controls whether the candlestick is “smoothed” or not. A smoothed candlestick uses the previous candle’s close price as its open price, while an unsmoothed candlestick uses the first price from its time range as its open price. [default=False]
func QPrWithSmooth() PricingOpts {
	return func(pq *pricingQuery) { pq.Smooth = "true" }
}
