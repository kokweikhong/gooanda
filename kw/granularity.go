package kw

type granularity struct {
	S5  string // 5 second candlesticks, minute alignment
	S10 string // 10 second candlesticks, minute alignment"
	S15 string // 15 second candlesticks, minute alignment"
	S30 string // 30 second candlesticks, minute alignment"
	M1  string // 1 minute candlesticks, minute alignment
	M2  string // 2 minute candlesticks, hour alignment
	M4  string // 4 minute candlesticks, hour alignment
	M5  string // 5 minute candlesticks, hour alignment
	M10 string // 10 minute candlesticks, hour alignment"
	M15 string // 15 minute candlesticks, hour alignment"
	M30 string // 30 minute candlesticks, hour alignment"
	H1  string // 1 hour candlesticks, hour alignment
	H2  string // 2 hour candlesticks, day alignment
	H3  string // 3 hour candlesticks, day alignment
	H4  string // 4 hour candlesticks, day alignment
	H6  string // 6 hour candlesticks, day alignment
	H8  string // 8 hour candlesticks, day alignment
	H12 string // 12 hour candlesticks, day alignment"
	D   string // 1 day candlesticks, day alignment
	W   string // 1 week candlesticks, aligned to start of week
	M   string // 1 month candlesticks, aligned to first day of the month
}

var GRANULARITY = &granularity{
	S5:  "S5",
	S10: "S10",
	S15: "S15",
	S30: "S30",
	M1:  "M1",
	M2:  "M2",
	M4:  "M4",
	M5:  "M5",
	M10: "M10",
	M15: "M15",
	M30: "M30",
	H1:  "H1",
	H2:  "H2",
	H3:  "H3",
	H4:  "H4",
	H6:  "H6",
	H8:  "H8",
	H12: "H12",
	D:   "D",
	W:   "W",
	M:   "M",
}
