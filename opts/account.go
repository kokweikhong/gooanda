package opts

import "strings"

type accountQuery struct { // {{{
	SinceTransactionID string `json:"sinceTransactionID,omitempty"`
	Instruments        string `json:"instruments,omitempty"`
} // }}}

type AccountOpts func(*accountQuery)

func NewAccountQuery(querys ...AccountOpts) *accountQuery {
	q := &accountQuery{}
	for _, query := range querys {
		query(q)
	}
	return q
}

// List of instruments to query specifically.
func QAccountInstruments(instruments []string) AccountOpts { // {{{
	return func(aq *accountQuery) {
		aq.Instruments = strings.Join(instruments, ",")
	}
} // }}}

// ID of the Transaction to get Account changes since.
func QAccountSinceTransactionID(transactionID string) AccountOpts { // {{{
	return func(aq *accountQuery) {
		aq.SinceTransactionID = transactionID
	}
} // }}}
