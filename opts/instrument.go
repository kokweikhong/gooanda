package opts

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kokweikhong/gooanda/kw"
)

type instrumentQuery struct {
	Price             string `json:"price,omitempty"`
	Granularity       string `json:"granularity,omitempty"`
	Count             string `json:"count,omitempty"`
	From              string `json:"from,omitempty"`
	To                string `json:"to,omitempty"`
	Smooth            string `json:"smooth,omitempty"`
	IncludeFirst      string `json:"includeFirst,omitempty"`
	DailyAlignment    string `json:"dailyAlignment,omitempty"`
	AlignmentTimezone string `json:"alignmentTimezone,omitempty"`
	WeeklyAlignment   string `json:"weeklyAlignment,omitempty"`
	Time              string `json:"time,omitempty"`
}

type InstrumentOpts func(*instrumentQuery)

func NewInstrumentQuery(querys ...InstrumentOpts) *instrumentQuery {
	q := &instrumentQuery{}
	for _, query := range querys {
		query(q)
	}
	return q
}

// The time of the snapshot to fetch. If not specified, then the most recent snapshot is fetched.
func QIntTime(setTime time.Time) InstrumentOpts {
	return func(iq *instrumentQuery) {
		if setTime.Unix() > time.Now().Unix() {
			iq.Time = ""
		}
		iq.Time = setTime.Format(time.RFC3339)
	}
}

// The day of the week used for granularities that have weekly alignment. [default=Friday]
func QIntWeeklyAlignment(day kw.WeeklyAlignment) InstrumentOpts {
	return func(iq *instrumentQuery) {
		iq.WeeklyAlignment = string(day)
	}
}

// The timezone to use for the dailyAlignment parameter. Candlesticks with daily alignment will be aligned to the dailyAlignment hour within the alignmentTimezone. Note that the returned times will still be represented in UTC. [default=America/New_York]
func QIntAlignmentTimezone(timezone string) InstrumentOpts {
	return func(iq *instrumentQuery) {
		iq.AlignmentTimezone = timezone
	}
}

// The hour of the day (in the specified timezone) to use for granularities that have daily alignments. [default=17, minimum=0, maximum=23]
func QIntDailyAlignment(alignment int) InstrumentOpts {
	return func(iq *instrumentQuery) {
		if alignment > 23 || alignment < 0 {
			iq.DailyAlignment = "17"
			return
		}
		iq.DailyAlignment = strconv.Itoa(alignment)
	}
}

// A flag that controls whether the candlestick that is covered by the from time should be included in the results. This flag enables clients to use the timestamp of the last completed candlestick received to poll for future candlesticks but avoid receiving the previous candlestick repeatedly. [default=True]
func QIntWithoutIncludeFirst() InstrumentOpts {
	return func(iq *instrumentQuery) {
		iq.IncludeFirst = "false"
	}
}

// A flag that controls whether the candlestick is “smoothed” or not. A smoothed candlestick uses the previous candle’s close price as its open price, while an un-smoothed candlestick uses the first price from its time range as its open price. [default=False]
func QIntWithSmooth() InstrumentOpts {
	return func(iq *instrumentQuery) {
		iq.Smooth = "true"
	}
}

// from: The start of the time range to fetch candlesticks for.
func QInFrom(from time.Time) InstrumentOpts {
	return func(iq *instrumentQuery) {
		if from.Unix() < time.Now().Unix() {
			iq.From = ""
			return
		}
		iq.From = from.Format(time.RFC3339)
	}
}

// from: The start of the time range to fetch candlesticks for.
// to: The end of the time range to fetch candlesticks for.
func QInFromTo(from, to time.Time) InstrumentOpts {
	return func(iq *instrumentQuery) {
		if to.Unix() > from.Unix() || to.Unix() > time.Now().Unix() {
			iq.From = ""
			iq.To = ""
			return
		}
		iq.From = from.Format(time.RFC3339)
		fmt.Println(iq.From)
		iq.To = to.Format(time.RFC3339)
	}
}

// The number of candlesticks to return in the response. Count should not be specified if both the start and end parameters are provided, as the time range combined with the granularity will determine the number of candlesticks to return. [default=500, maximum=5000]
func QInCount(count int) InstrumentOpts {
	return func(iq *instrumentQuery) {
		c := strconv.Itoa(count)
		if count > 5000 {
			iq.Count = "500"
		}
		iq.Count = c
	}
}

// The granularity of the candlesticks to fetch [default=S5]
func QInGranularity(granularity kw.Granularity) InstrumentOpts {
	return func(iq *instrumentQuery) {
		iq.Granularity = string(granularity)
	}
}

// The Price component(s) to get candlestick data for. [default=M]
func QInPrice(price string) InstrumentOpts {
	return func(iq *instrumentQuery) {
		iq.Price = price
	}
}
