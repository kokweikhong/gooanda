package gooanda

type trade struct {
	connection
}

func NewTradeConnection(token string) *trade {
	con := &trade{}
	con.connection.token = token
	return con
}

// GetTradeList is to get a list of Trades for an Account.
func (tr *trade) GetTradeList() {}

// GetOpenTradeList is to get the list of open Trades for an Account.
func (tr *trade) GetOpenTradeList() {}

// GetSpecificTradeDetails is to get the details of a specific Trade in an Account.
func (tr *trade) GetSpecificTradeDetails() {}
