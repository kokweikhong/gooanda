package kw

type timeInForce struct {
	GTC string // The Order is “Good unTil Cancelled”
	GTD string // The Order is “Good unTil Date” and will be cancelled at the provided time
	GFD string // The Order is “Good For Day” and will be cancelled at 5pm New York time
	FOK string // The Order must be immediately “Filled Or Killed”
	IOC string // The Order must be “Immediately partially filled Or Cancelled”
}

var TIMEINFORCE = &timeInForce{
	GTC: "GTC",
	GTD: "GTD",
	GFD: "GFD",
	FOK: "FOK",
	IOC: "IOC",
}
