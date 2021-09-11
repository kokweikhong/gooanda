package opts

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kokweikhong/gooanda/kw"
)

type pricingQuery struct {
	// List of CandleSpecification (csv)
	// List of candle specifications to get pricing for. [required]
	CandleSpecifications   string `json:"candleSpecifications,omitempty"`
	Units                  string `json:"units,omitempty"`
	Smooth                 string `json:"smooth,omitempty"`
	DailyAlignment         string `json:"dailyAlignment,omitempty"`
	AlignmentTimezone      string `json:"alignmentTimezone,omitempty"`
	WeeklyAlignment        string `json:"weeklyAlignment,omitempty"`
	Instruments            string `json:"instruments,omitempty"`
	Since                  string `json:"since,omitempty"`
	IncludeHomeConversions string `json:"includeHomeConversions,omitempty"`
	Snapshot               string `json:"snapshot,omitempty"`
	Price                  string `json:"price,omitempty"`
	Granularity            string `json:"granularity,omitempty"`
	Count                  string `json:"count,omitempty"`
	From                   string `json:"from,omitempty"`
	To                     string `json:"to,omitempty"`
	IncludeFirst           string `json:"includeFirst,omitempty"`
}

type PricingOpts func(*pricingQuery)

func NewPricingQuery(querys ...PricingOpts) *pricingQuery {
	q := &pricingQuery{}
	for _, query := range querys {
		query(q)
	}
	return q
}

func QPrCandleSpecifications(instruments []string, granularity kw.Granularity, priceComponent kw.PriceComponent) PricingOpts {
	return func(pq *pricingQuery) {
		for k := range instruments {
			instruments[k] = instruments[k] + ":" + string(granularity) + ":" + string(priceComponent)
		}
		pq.CandleSpecifications = strings.Join(instruments, ",")
	}
}

// QrPrFromTo:
// The start of the time range to fetch candlesticks for.
// The end of the time range to fetch candlesticks for.
func QPrFromTo(from, to time.Time) PricingOpts {
	return func(pq *pricingQuery) {
		if to.Unix() < from.Unix() || to.Unix() > time.Now().Unix() {
			pq.From = ""
			pq.To = ""
			return
		} else {
			pq.From = from.Format(time.RFC3339)
			pq.To = to.Format(time.RFC3339)
		}
	}
}

// QPrSince: Date/Time filter to apply to the response. Only prices and home conversions (if requested) with a time later than this filter (i.e. the price has changed after the since time) will be provided, and are filtered independently.
func QPrSince(since time.Time) PricingOpts {
	return func(pq *pricingQuery) {
		if since.Unix() > time.Now().Unix() {
			pq.Since = time.Now().Format(time.RFC3339)
			return
		} else {
			pq.Since = since.Format(time.RFC3339)
		}
	}
}

// QrPrUnits: The number of units used to calculate the volume-weighted average bid and ask prices in the returned candles. [default=1]
func QPrUnits(units float64) PricingOpts {
	return func(pq *pricingQuery) {
		if units < 1 {
			pq.Units = "1"
			return
		} else {
			pq.Units = fmt.Sprintf("%v", units)
		}
	}
}

// QPrInstruments: List of InstrumentName (csv)	List of Instruments to get pricing for. [required]
func QPrInstruments(instruments []string) PricingOpts {
	return func(pq *pricingQuery) {
		pq.Instruments = strings.Join(instruments, ",")
	}
}

// QPrPrice: The Price component(s) to get candlestick data for. [default=M]
func QPrPrice(component kw.PriceComponent) PricingOpts {
	return func(pq *pricingQuery) {
		pq.Price = string(component)
	}
}

// QPrCount: The number of candlesticks to return in the response. Count should not be specified if both the start and end parameters are provided, as the time range combined with the granularity will determine the number of candlesticks to return. [default=500, maximum=5000]
func QPrCount(count int) PricingOpts {
	return func(pq *pricingQuery) {
		if count < 1 || count > 5000 {
			pq.Count = "500"
			return
		}
		pq.Count = strconv.Itoa(count)
	}
}

// QPrWithoutIncludeFirst:  flag that controls whether the candlestick that is covered by the from time should be included in the results. This flag enables clients to use the timestamp of the last completed candlestick received to poll for future candlesticks but avoid receiving the previous candlestick repeatedly. [default=True]
func QPrWithoutIncludeFirst() PricingOpts {
	return func(pq *pricingQuery) { pq.IncludeFirst = "false" }
}

// QPrGranularity:  granularity of the candlesticks to fetch [default=S5]
func QPrGranularity(granularity kw.Granularity) PricingOpts {
	return func(pq *pricingQuery) { pq.Granularity = string(granularity) }
}

// QPrWithoutSnapshot: that enables/disables the sending of a pricing snapshot when initially connecting to the stream. [default=True]
func QPrWithoutSnapshot() PricingOpts {
	return func(pq *pricingQuery) { pq.Snapshot = "false" }
}

// QPrWithIncludeHomeConversions: that enables the inclusion of the homeConversions field in the returned response. An entry will be returned for each currency in the set of all base and quote currencies present in the requested instruments list. [default=False]
func QPrWithIncludeHomeConversions() PricingOpts {
	return func(pq *pricingQuery) { pq.IncludeHomeConversions = "true" }
}

// QPrWeeklyAlignment:  day of the week used for granularities that have weekly alignment. [default=Friday]
func QPrWeeklyAlignment(day kw.WeeklyAlignment) PricingOpts {
	return func(pq *pricingQuery) { pq.WeeklyAlignment = string(day) }
}

// QPrAlignmentTimezone:  timezone to use for the dailyAlignment parameter. Candlesticks with daily alignment will be aligned to the dailyAlignment hour within the alignmentTimezone. Note that the returned times will still be represented in UTC. [default=America/New_York]
func QPrAlignmentTimezone(timezone string) PricingOpts {
	return func(pq *pricingQuery) { pq.AlignmentTimezone = timezone }
}

// QPrDailyAlignment: hour of the day (in the specified timezone) to use for granularities that have daily alignments. [default=17, minimum=0, maximum=23]
func QPrDailyAlignment(alignment int) PricingOpts {
	return func(pq *pricingQuery) {
		if alignment < 0 || alignment > 23 {
			pq.DailyAlignment = ""
			return
		}
		pq.DailyAlignment = strconv.Itoa(alignment)
	}
}

// QPrWithSmooth: flag that controls whether the candlestick is “smoothed” or not. A smoothed candlestick uses the previous candle’s close price as its open price, while an unsmoothed candlestick uses the first price from its time range as its open price. [default=False]
func QPrWithSmooth() PricingOpts {
	return func(pq *pricingQuery) { pq.Smooth = "true" }
}
