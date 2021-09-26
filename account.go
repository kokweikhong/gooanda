package gooanda

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kokweikhong/gooanda/endpoint"
)

// Accounts data structure
type AccountList struct { // {{{
	Accounts []struct {
		ID           string   `json:"id"`
		Mt4AccountID int      `json:"mt4AccountID,omitempty"`
		Tags         []string `json:"tags"`
	}
} // }}}

// AccountById data structure
type AccountById struct { // {{{
	Account struct {
		accountGlobal
		Postions []struct {
			Instrument   string             `json:"instrument"`
			Long         struct{ position } `json:"long"`
			Short        struct{ position } `json:"short"`
			UnrealizedPL string             `json:"unrealizedPL"`
		} `json:"positions"`
	} `json:"account"`
} // }}}

// AccountSummary data structure
type AccountSummary struct { // {{{
	Account struct {
		accountGlobal
		ResettablePL    string `json:"resettablePL"`
		UnrealizedPL    string `json:"unrealizedPL"`
		WithdrawalLimit string `json:"withdrawalLimit"`
	} `json:"account"`
} // }}}

// AccountInstruments data structure
type AccountInstruments struct { // {{{
	Intruments []struct {
		DisplayName                 string `json:"displayName"`
		DisplayPrecision            int    `json:"displayPrecision"`
		MarginRate                  string `json:"marginRate"`
		MaximumOrderUnits           string `json:"maximumOrderUnits"`
		MaximumPositionSize         string `json:"maximumPositionSize"`
		MaximumTrailingStopDistance string `json:"maximumTrailingStopDistance"`
		MinimumTradeSize            string `json:"minimumTradeSize"`
		MinimumTrailingStopDistance string `json:"minimumTrailingStopDistance"`
		Name                        string `json:"name"`
		PipLocation                 int    `json:"pipLocation"`
		TradeUnitsPrecision         int    `json:"tradeUnitsPrecision"`
		Type                        string `json:"type"`
	} `json:"instruments"`
	LastTransactionID string `json:"lastTransactionID"`
} // }}}

type accountGlobal struct { // {{{
	NAV                         string `json:"NAV"`
	Alias                       string `json:"alias"`
	Balance                     string `json:"balance"`
	CreatedByUserID             int    `json:"createdByUserID"`
	CreatedTime                 string `json:"createdTime"`
	Currency                    string `json:"currency"`
	HedgingEnabled              bool   `json:"hedgingEnabled"`
	Id                          string `json:"id"`
	LastTransactionID           string `json:"lastTransactionID"`
	MarginAvailable             string `json:"marginAvailable"`
	MarginCloseoutMarginUsed    string `json:"marginCloseoutMarginUsed"`
	MarginCloseoutNAV           string `json:"marginCloseoutNAV"`
	MarginCloseoutPercent       string `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue string `json:"marginCloseoutPositionValue"`
	MarginCloseoutUnrealizedPL  string `json:"marginCloseoutUnrealizedPL"`
	MarginRate                  string `json:"marginRate"`
	MarginUsed                  string `json:"marginUsed"`
	OpenPositionCount           int    `json:"openPositionCount"`
	OpenTradeCount              int    `json:"openTradeCount"`
	Orders                      []int  `json:"orders"`
	PendingOrderCount           int    `json:"pendingOrderCount"`
	PL                          string `json:"pl"`
	PositionValue               string `json:"positionValue"`
} // }}}

type position struct {
	PL           string `json:"pl"`
	ResettablePL string `json:"resettablePL"`
	Units        string `json:"units"`
	UnrealizedPL string `json:"unrealizedPL"`
}

type account struct {
	connection
	Query *accountFunc
}

// NewAccountConnection create a new connection for account endpoint
func NewAccountConnection(token string) *account {
	con := &account{}
	con.connection.token = token
	return con
}

func (ac *account) connect() ([]byte, error) {
	con := &connection{ac.endpoint, ac.method, ac.token, ac.data}
	resp, err := con.connect()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Get the list of tradeable instruments for the given Account. The list of tradeable instruments is dependent on the regulatory division that the Account is located in, thus should be the same for all Accounts owned by a single user.
func (ac *account) GetAccountInstruments(live bool, accountID string, querys ...accountOpts) (*AccountInstruments, error) { // {{{
	q := newAccountQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Account.AccountInstrument)
	url := fmt.Sprintf(ep, accountID)
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	ac.endpoint = u
	ac.method = http.MethodGet
	resp, err := ac.connect()
	if err != nil {
		return nil, err
	}
	data := &AccountInstruments{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// Get a summary for a single Account that a client has access to.
func (ac *account) GetAccountSummary(live bool, accountID string, querys ...accountOpts) (*AccountSummary, error) { // {{{
	q := newAccountQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Account.AccountSummary)
	url := fmt.Sprintf(ep, accountID)
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	ac.endpoint = u
	ac.method = http.MethodGet
	resp, err := ac.connect()
	if err != nil {
		return nil, err
	}
	data := &AccountSummary{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// Get the full details for a single Account that a client has access to. Full pending Order, open Trade and open Position representations are provided.
func (ac *account) GetAccountById(live bool, accountID string, querys ...accountOpts) (*AccountById, error) { // {{{
	q := newAccountQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Account.AccountsById)
	url := fmt.Sprintf(ep, accountID)
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	ac.endpoint = u
	ac.method = http.MethodGet
	resp, err := ac.connect()
	if err != nil {
		return nil, err
	}
	data := &AccountById{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// Get the full details for a single Account that a client has access to. Full pending Order, open Trade and open Position representations are provided.
func (ac *account) GetAccountList(live bool, querys ...accountOpts) (*AccountList, error) { // {{{
	q := newAccountQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Account.Accounts)
	url := ep
	u, err := urlAddQuery(url, q)
	if err != nil {
		return nil, err
	}
	ac.endpoint = u
	ac.method = http.MethodGet
	resp, err := ac.connect()
	if err != nil {
		return nil, err
	}
	var data = &AccountList{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

type accountQuery struct { // {{{
	SinceTransactionID string `json:"sinceTransactionID,omitempty"`
	Instruments        string `json:"instruments,omitempty"`
} // }}}

type accountOpts func(*accountQuery)

func newAccountQuery(querys ...accountOpts) *accountQuery {
	q := &accountQuery{}
	for _, query := range querys {
		query(q)
	}
	return q
}

type accountFunc struct{}

// List of instruments to query specifically.
func (*accountFunc) WithInstruments(instruments []string) accountOpts { // {{{
	return func(aq *accountQuery) {
		aq.Instruments = strings.Join(instruments, ",")
	}
} // }}}

// ID of the Transaction to get Account changes since.
func (*accountFunc) WithSinceTransactionID(transactionID string) accountOpts { // {{{
	return func(aq *accountQuery) {
		aq.SinceTransactionID = transactionID
	}
} // }}}
