package kw

type priceComponent string

const (
	ComPrice_M priceComponent = "M" // mid point candles
	ComPrice_B priceComponent = "B" // bid candles
	ComPrice_A priceComponent = "A" // ask candles
)
