package kw

type TimeInForce string

const (
	TimeIF_GTC TimeInForce = "GTC" // The Order is “Good unTil Cancelled”
	TimeIF_GTD TimeInForce = "GTD" // The Order is “Good unTil Date” and will be cancelled at the provided time
	TimeIF_GFD TimeInForce = "GFD" // The Order is “Good For Day” and will be cancelled at 5pm New York time
	TimeIF_FOK TimeInForce = "FOK" // The Order must be immediately “Filled Or Killed”
	TimeIF_IOC TimeInForce = "IOC" // The Order must be “Immediately partially filled Or Cancelled”
)
