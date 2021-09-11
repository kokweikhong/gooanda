package kw

type PriceComponent string

const (
	ComPrice_M  PriceComponent = "M"  // mid point candles
	ComPrice_B  PriceComponent = "B"  // bid candles
	ComPrice_A  PriceComponent = "A"  // ask candles
	ComPrice_AM PriceComponent = "AM" // ask/mid point candles
	ComPrice_BM PriceComponent = "BM" // bid/mid point candles
	ComPrice_AB PriceComponent = "AB" // ask/bid candles

	ComPrice_BAM PriceComponent = "BAM" // ask/mid/bid candles
)
