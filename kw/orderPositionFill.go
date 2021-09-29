package kw

type positionFill struct {
	OPEN_ONLY    string // When the Order is filled, only allow Positions to be opened or extended.
	REDUCE_FIRST string // When the Order is filled, always fully reduce an existing Position before opening a new Position.
	REDUCE_ONLY  string // When the Order is filled, only reduce an existing Position.
	DEFAULT      string // When the Order is filled, use REDUCE_FIRST behaviour for non-client hedging Accounts, and OPEN_ONLY behaviour for client hedging Accounts.
}

var POSITIONFILL = &positionFill{
	OPEN_ONLY:    "OPEN_ONLY",
	REDUCE_FIRST: "REDUCE_FIRST",
	REDUCE_ONLY:  "REDUCE_ONLY",
	DEFAULT:      "DEFAULT",
}
