package kw

type tradeStateFilter struct {
	OPEN, CLOSED, CLOSE_WHEN_TRADEABLE, ALL string
}

var TRADESTATEFILTER = &tradeStateFilter{
	OPEN:                 "OPEN",
	CLOSED:               "CLOSED",
	CLOSE_WHEN_TRADEABLE: "CLOSE_WHEN_TRADEABLE",
	ALL:                  "ALL",
}
