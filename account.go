package gooanda

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kokweikhong/gooanda/addr"
	"github.com/kokweikhong/gooanda/opts"
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

type accountGlobal struct {
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
}

type position struct {
	PL           string `json:"pl"`
	ResettablePL string `json:"resettablePL"`
	Units        string `json:"units"`
	UnrealizedPL string `json:"unrealizedPL"`
}

type account connection

// NewAccountConnection create a new connection for account endpoint
func NewAccountConnection(token string) *account {
	return &account{token: token}
}

func (ac *account) connect() []byte {
	con := &connection{ac.endpoint, ac.method, ac.token, ac.data}
	return con.connect()
}

// Get the list of tradeable instruments for the given Account. The list of tradeable instruments is dependent on the regulatory division that the Account is located in, thus should be the same for all Accounts owned by a single user.
func (ac *account) GetAccountInstruments(accountID string, querys ...opts.AccountOpts) *AccountInstruments { // {{{
	q := opts.NewAccountQuery(querys...)
	u, err := urlAddQuery(fmt.Sprintf(addr.AccountInstrument, accountID), q)
	if err != nil {
		log.Fatal(err)
	}
	ac.endpoint = u
	ac.method = http.MethodGet
	resp := ac.connect()
	data := &AccountInstruments{}
	if err = json.Unmarshal(resp, &data); err != nil {
		log.Fatal(err)
	}
	return data
} // }}}

// Get a summary for a single Account that a client has access to.
func (ac *account) GetAccountSummary(accountID string, querys ...opts.AccountOpts) *AccountSummary { // {{{
	q := opts.NewAccountQuery(querys...)
	fmt.Println(q.Instruments, q.SinceTransactionID)
	u, err := urlAddQuery(fmt.Sprintf(addr.AccountSummary, accountID), q)
	if err != nil {
		log.Fatal(err)
	}
	ac.endpoint = u
	ac.method = http.MethodGet
	resp := ac.connect()
	data := &AccountSummary{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
} // }}}

// Get the full details for a single Account that a client has access to. Full pending Order, open Trade and open Position representations are provided.
func (ac *account) GetAccountById(accountID string, querys ...opts.AccountOpts) *AccountById { // {{{
	q := opts.NewAccountQuery(querys...)
	u, err := urlAddQuery(fmt.Sprintf(addr.AccountsById, accountID), q)
	if err != nil {
		log.Fatal(err)
	}
	ac.endpoint = u
	ac.method = http.MethodGet
	resp := ac.connect()
	data := &AccountById{}
	if err = json.Unmarshal(resp, &data); err != nil {
		log.Fatal(err)
	}
	return data
} // }}}

// Get the full details for a single Account that a client has access to. Full pending Order, open Trade and open Position representations are provided.
func (ac *account) GetAccountList(querys ...opts.AccountOpts) *AccountList { // {{{
	q := opts.NewAccountQuery(querys...)
	u, err := urlAddQuery(addr.Accounts, q)
	if err != nil {
		log.Fatal(err)
	}
	ac.endpoint = u
	ac.method = http.MethodGet
	resp := ac.connect()
	var data = &AccountList{}
	if err = json.Unmarshal(resp, &data); err != nil {
		log.Fatal(err)
	}
	return data
} // }}}
