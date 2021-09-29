package kw

type orderState struct {
	PENDING   string // Orders that are currently pending execution
	FILLED    string // The Orders that have been filled
	TRIGGERED string // The Orders that have been triggered
	CANCELLED string // The Orders that have been cancelled
	ALL       string // The Orders that are in any of the possible states listed above
}

var OrderState = &orderState{
	PENDING:   "PENDING",
	FILLED:    "FILLED",
	TRIGGERED: "TRIGGERED",
	CANCELLED: "CANCELLED",
}
