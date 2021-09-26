package kw

type OrderPositionFill string

const (
	OPF_OPEN_ONLY    OrderPositionFill = "" // When the Order is filled, only allow Positions to be opened or extended.
	OPF_REDUCE_FIRST OrderPositionFill = "" // When the Order is filled, always fully reduce an existing Position before opening a new Position.
	OPF_REDUCE_ONLY  OrderPositionFill = "" // When the Order is filled, only reduce an existing Position.
	OPF_DEFAULT      OrderPositionFill = "" // When the Order is filled, use REDUCE_FIRST behaviour for non-client hedging Accounts, and OPEN_ONLY behaviour for client hedging Accounts.
)
