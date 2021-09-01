package gooanda

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kokweikhong/gooanda/addr"
)

type Accounts struct {
	Accounts []struct {
		ID           string   `json:"id"`
		Mt4AccountID int      `json:"mt4AccountID,omitempty"`
		Tags         []string `json:"tags"`
	}
}

type AccountById struct {
	Account struct {
		ID              string      `json:"id"`
		Alias           string      `json:"alias"`
		Currency        string      `json:"currency"`
		CreatedByUserID int         `json:"createdByUserID"`
		CreatedTime     interface{} `json:"createdTime"`
	} `json:"account"`
}

func GetAccountInstruments(token, accountID string, querys ...accountOpts) {
	q := newAccountQuery(querys...)
	u, err := urlAddQuery(fmt.Sprintf(addr.AccountInstrument, accountID), q)
	if err != nil {
		log.Fatal(err)
	}
	conn := &connection{u, http.MethodGet, token}
	resp := conn.connect()
	fmt.Println(string(resp))
}

func GetAccountSummary(token, accountID string, querys ...accountOpts) {
	q := newAccountQuery(querys...)
	fmt.Println(q.Instruments, q.SinceTransactionID)
	u, err := urlAddQuery(fmt.Sprintf(addr.AccountSummary, accountID), q)
	if err != nil {
		log.Fatal(err)
	}
	conn := &connection{u, http.MethodGet, token}
	resp := conn.connect()
	fmt.Println(string(resp))
}

// Get the full details for a single Account that a client has access to. Full pending Order, open Trade and open Position representations are provided.
func GetAccountById(token, accountID string, querys ...accountOpts) {
	q := newAccountQuery(querys...)
	u, err := urlAddQuery(fmt.Sprintf(addr.AccountsById, accountID), q)
	if err != nil {
		log.Fatal(err)
	}
	conn := &connection{u, http.MethodGet, token}
	resp := conn.connect()
	fmt.Println(string(resp))
}

// Get the full details for a single Account that a client has access to. Full pending Order, open Trade and open Position representations are provided.
func GetAccounts(token string, querys ...accountOpts) *Accounts {
	q := newAccountQuery(querys...)
	u, err := urlAddQuery(addr.Accounts, q)
	if err != nil {
		log.Fatal(err)
	}
	conn := &connection{
		endpoint: u,
		method:   http.MethodGet,
		token:    token,
	}

	resp := conn.connect()
	var data = &Accounts{}
	if err = json.Unmarshal(resp, &data); err != nil {
		log.Fatal(err)
	}
	return data
}

type accountQuery struct {
	SinceTransactionID string `json:"sinceTransactionID"`
	Instruments        string `json:"instruments"`
}

type accountOpts func(*accountQuery)

func newAccountQuery(querys ...accountOpts) *accountQuery {
	q := &accountQuery{}
	for _, query := range querys {
		query(q)
	}
	return q
}

// List of instruments to query specifically.
func QAccountInstruments(instruments []string) accountOpts {
	return func(aq *accountQuery) {
		aq.Instruments = strings.Join(instruments, ",")
	}
}

// ID of the Transaction to get Account changes since.
func QAccountSinceTransactionID(transactionID string) accountOpts {
	return func(aq *accountQuery) {
		aq.SinceTransactionID = transactionID
	}
}
