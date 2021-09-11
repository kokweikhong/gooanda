package kw

type OrderState string

const (
	PENDING   OrderState = "PENDING"   // Orders that are currently pending execution
	FILLED    OrderState = "FILLED"    // The Orders that have been filled
	TRIGGERED OrderState = "TRIGGERED" // The Orders that have been triggered
	CANCELLED OrderState = "CANCELLED" // The Orders that have been cancelled
	ALL                                // The Orders that are in any of the possible states listed above
)
