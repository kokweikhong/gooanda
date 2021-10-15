package gooanda

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kokweikhong/gooanda/endpoint"
)

// GetAccountList data structure
type accountList struct { // {{{
	Accounts []struct {
		ID           string   `json:"id"`
		Mt4AccountID int      `json:"mt4AccountID,omitempty"`
		Tags         []string `json:"tags"`
	}
} // }}}

// GetAccountById data structure
type accountById struct { // {{{
	Account struct {
		accountGlobal
		Positions []struct {
			Instrument   string                    `json:"instrument"`
			Long         struct{ accountPosition } `json:"long"`
			Short        struct{ accountPosition } `json:"short"`
			UnrealizedPL float64                   `json:"unrealizedPL,string"`
		} `json:"positions"`
	} `json:"account"`
} // }}}

// GetAccountSummary data structure.
type accountSummary struct { // {{{
	Account struct {
		accountGlobal
		ResettablePL    float64 `json:"resettablePL,string"`
		UnrealizedPL    float64 `json:"unrealizedPL,string"`
		WithdrawalLimit float64 `json:"withdrawalLimit,string"`
	} `json:"account"`
} // }}}

// GetAccountInstruments data structure.
type accountInstruments struct { // {{{
	Intruments []struct {
		DisplayName                 string  `json:"displayName"`
		DisplayPrecision            int     `json:"displayPrecision"`
		MarginRate                  float64 `json:"marginRate,string"`
		MaximumOrderUnits           float64 `json:"maximumOrderUnits,string"`
		MaximumPositionSize         float64 `json:"maximumPositionSize,string"`
		MaximumTrailingStopDistance float64 `json:"maximumTrailingStopDistance,string"`
		MinimumTradeSize            float64 `json:"minimumTradeSize,string"`
		MinimumTrailingStopDistance float64 `json:"minimumTrailingStopDistance,string"`
		Name                        string  `json:"name"`
		PipLocation                 int     `json:"pipLocation"`
		TradeUnitsPrecision         int     `json:"tradeUnitsPrecision"`
		Type                        string  `json:"type"`
	} `json:"instruments"`
	LastTransactionID string `json:"lastTransactionID"`
} // }}}

type accountGlobal struct { // {{{
	NAV                         string      `json:"NAV"`
	Alias                       string      `json:"alias"`
	Balance                     float64     `json:"balance,string"`
	CreatedByUserID             int         `json:"createdByUserID"`
	CreatedTime                 time.Time   `json:"createdTime"`
	Currency                    string      `json:"currency"`
	HedgingEnabled              bool        `json:"hedgingEnabled"`
	Id                          string      `json:"id"`
	LastTransactionID           string      `json:"lastTransactionID"`
	MarginAvailable             float64     `json:"marginAvailable,string"`
	MarginCloseoutMarginUsed    float64     `json:"marginCloseoutMarginUsed,string"`
	MarginCloseoutNAV           float64     `json:"marginCloseoutNAV,string"`
	MarginCloseoutPercent       float64     `json:"marginCloseoutPercent,string"`
	MarginCloseoutPositionValue float64     `json:"marginCloseoutPositionValue,string"`
	MarginCloseoutUnrealizedPL  float64     `json:"marginCloseoutUnrealizedPL,string"`
	MarginRate                  float64     `json:"marginRate,string"`
	MarginUsed                  float64     `json:"marginUsed,string"`
	OpenPositionCount           int         `json:"openPositionCount"`
	OpenTradeCount              int         `json:"openTradeCount"`
	Orders                      interface{} `json:"orders"`
	PendingOrderCount           int         `json:"pendingOrderCount"`
	PL                          string      `json:"pl"`
	PositionValue               string      `json:"positionValue"`
} // }}}

type accountPosition struct {
	PL           float64 `json:"pl,string"`
	ResettablePL float64 `json:"resettablePL,string"`
	Units        float64 `json:"units,string"`
	UnrealizedPL float64 `json:"unrealizedPL,string"`
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

// GetAccountInstruments is to get the list of tradeable instruments for the given Account.
// The list of tradeable instruments is dependent on the regulatory division
// that the Account is located in, thus should be the same for all Accounts owned by a single user.
func (ac *account) GetAccountInstruments(live bool, accountID string, querys ...accountOpts) (*accountInstruments, error) { // {{{
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
	data := &accountInstruments{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// GetAccountSummary is to get a summary for a single Account that a client has access to.
func (ac *account) GetAccountSummary(live bool, accountID string, querys ...accountOpts) (*accountSummary, error) { // {{{
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
	data := &accountSummary{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// GetAccountById is to get the full details for a single Account that a client has access to.
// Full pending Order, open Trade and open Position representations are provided.
func (ac *account) GetAccountById(live bool, accountID string, querys ...accountOpts) (*accountById, error) { // {{{
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
	data := &accountById{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}
	return data, nil
} // }}}

// GetAccountList is to get the full details for a single Account that a client has access to.
// Full pending Order, open Trade and open Position representations are provided.
func (ac *account) GetAccountList(live bool, querys ...accountOpts) (*accountList, error) { // {{{
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
	var data = &accountList{}
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

// WithInstruments is list of instruments to query specifically.
func (*accountFunc) WithInstruments(instruments []string) accountOpts { // {{{
	return func(aq *accountQuery) {
		aq.Instruments = strings.Join(instruments, ",")
	}
} // }}}

// WithSinceTransactionID is ID of the Transaction to get Account changes since.
func (*accountFunc) WithSinceTransactionID(transactionID string) accountOpts { // {{{
	return func(aq *accountQuery) {
		aq.SinceTransactionID = transactionID
	}
} // }}}
