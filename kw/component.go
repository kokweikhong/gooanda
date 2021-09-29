package kw

type priceComponent struct {
	M   string // mid point candles
	B   string // bid candles
	A   string // ask candles
	AB  string // ask and bid candles
	AM  string // ask and mid candles
	BM  string // bid and mid candles
	ALL string // ask, bid and mid candles
}

var PRICECOMPONENT = &priceComponent{
	M:   "M",
	B:   "B",
	A:   "A",
	AB:  "AB",
	AM:  "AM",
	BM:  "BM",
	ALL: "ABM",
}
