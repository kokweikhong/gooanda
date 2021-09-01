package kw

type Granularity string

const (
	Granularity_S5  Granularity = "S5"  // 5 second candlesticks, minute alignment
	Granularity_S10 Granularity = "S10" // 10 second candlesticks, minute alignment
	Granularity_S15 Granularity = "S15" // 15 second candlesticks, minute alignment
	Granularity_S30 Granularity = "S30" // 30 second candlesticks, minute alignment
	Granularity_M1  Granularity = "M1"  // 1 minute candlesticks, minute alignment
	Granularity_M2  Granularity = "M2"  // 2 minute candlesticks, hour alignment
	Granularity_M4  Granularity = "M4"  // 4 minute candlesticks, hour alignment
	Granularity_M5  Granularity = "M5"  // 5 minute candlesticks, hour alignment
	Granularity_M10 Granularity = "M10" // 10 minute candlesticks, hour alignment
	Granularity_M15 Granularity = "M15" // 15 minute candlesticks, hour alignment
	Granularity_M30 Granularity = "M30" // 30 minute candlesticks, hour alignment
	Granularity_H1  Granularity = "H1"  // 1 hour candlesticks, hour alignment
	Granularity_H2  Granularity = "H2"  // 2 hour candlesticks, day alignment
	Granularity_H3  Granularity = "H3"  // 3 hour candlesticks, day alignment
	Granularity_H4  Granularity = "H4"  // 4 hour candlesticks, day alignment
	Granularity_H6  Granularity = "H6"  // 6 hour candlesticks, day alignment
	Granularity_H8  Granularity = "H8"  // 8 hour candlesticks, day alignment
	Granularity_H12 Granularity = "H12" // 12 hour candlesticks, day alignment
	Granularity_D   Granularity = "D"   // 1 day candlesticks, day alignment
	Granularity_W   Granularity = "W"   // 1 week candlesticks, aligned to start of week
	Granularity_M   Granularity = "M"   // 1 month candlesticks, aligned to first day of the month
)
