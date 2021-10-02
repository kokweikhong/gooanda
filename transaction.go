package gooanda

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kokweikhong/gooanda/endpoint"
)

type transaction struct {
	connection
	Query *transactionFunc
}

func NewTransactionConnection(token string) *transaction {
	conn := &transaction{}
	conn.token = token
	return conn
}

// GetTransactions data structure
type transactions struct {
	From              string   `json:"from"`
	To                string   `json:"to"`
	PageSize          int      `json:"pageSize"`
	Type              []string `json:"type,omitempty"`
	Count             int      `json:"count"`
	Pages             []string `json:"pages"`
	LastTransactionID string   `json:"lastTransactionID"`
}

// GetTransactions is to get a list of Transactions pages
// that satisfy a time-based Transaction query.
func (tc *transaction) GetTransactions(live bool, accountID string, querys ...transactionOpts) (*transactions, error) {
	var result *transactions
	query := newTransactionQuery(querys...)
	ep := endpoint.GetEndpoint(live, endpoint.Transaction.Transactions)
	url, _ := urlAddQuery(fmt.Sprintf(ep, accountID), query)
	tc.endpoint = url
	tc.method = http.MethodGet
	data, err := tc.connect()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s to %T, %v", string(data), result, err)
	}
	return result, nil
}

// GetTransactionById is to get the details of a single Account Transaction.
func (tc *transaction) GetTransactionById(live bool, accountID, transactionID string) (string, error) {
	ep := endpoint.GetEndpoint(live, endpoint.Transaction.TransactionById)
	tc.endpoint = fmt.Sprintf(ep, accountID, transactionID)
	tc.method = http.MethodGet
	data, err := tc.connect()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// GetTransactionRangeById is to get a range of Transactions
// for an Account based on the Transaction IDs.
func (tc *transaction) GetTransactionRangeById(live bool, accountID, fromID, toID string, opts ...transactionOpts) (string, error) {
	query := newTransactionQuery(opts...)
	queryMap := struct {
		From string `json:"from"`
		To   string `json:"to"`
		Type string `json:"type,omitempty"`
	}{fromID, toID, query.Type}
	ep := endpoint.GetEndpoint(live, endpoint.Transaction.TransactionIdRange)
	url, err := urlAddQuery(fmt.Sprintf(ep, accountID), queryMap)
	if err != nil {
		return "", err
	}
	tc.endpoint = url
	tc.method = http.MethodGet
	data, err := tc.connect()
	if err != nil {
		return "", err
	}
	return string(data), err
}

// GetTransactionRange is to get a range of Transactions for
// an Account starting at (but not including) a provided Transaction ID.
func (tc *transaction) GetTransactionRange(live bool, accountID, transactionID string, opts ...transactionOpts) (string, error) {
	query := newTransactionQuery(opts...)
	ep := endpoint.GetEndpoint(live, endpoint.Transaction.TransactionIdRange)
	url := fmt.Sprintf(ep, accountID) + "?id=" + transactionID
	url, err := urlAddQuery(url, &transactionQuery{Type: query.Type})
	if err != nil {
		return "", err
	}
	tc.endpoint = url
	tc.method = http.MethodGet
	data, err := tc.connect()
	if err != nil {
		return "", nil
	}
	return string(data), nil
}

//
// GET	/v3/accounts/{accountID}/transactions/stream
// Get a stream of Transactions for an Account starting from when the request is made.

type transactionQuery struct {
	FromDate string `json:"from,omitempty"`
	ToDate   string `json:"to,omitempty"`
	PageSize int    `json:"pageSize,omitempty,string"`
	Type     string `json:"type,omitempty"`
}

type transactionFunc struct{}

type transactionOpts func(*transactionQuery)

func newTransactionQuery(opts ...transactionOpts) *transactionQuery {
	t := &transactionQuery{}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

// WithType is a filter for restricting the types of Transactions to retrieve.
func (tf *transactionFunc) WithType(transactionType ...string) transactionOpts {
	return func(tq *transactionQuery) {
		tq.Type = strings.Join(transactionType, ",")
	}
}

// WithFromDate is the starting time (inclusive) of the time range for the
// Transactions being queried. [default=Account Creation Time]
func (tf *transactionFunc) WithFromDate(from time.Time) transactionOpts {
	return func(tq *transactionQuery) { tq.FromDate = from.Format(time.RFC3339) }
}

// WithToDate is the ending time (inclusive) of the time range for the
// Transactions being queried. [default=Request Time]
func (tf *transactionFunc) WithToDate(to time.Time) transactionOpts {
	return func(tq *transactionQuery) { tq.ToDate = to.Format(time.RFC3339) }
}

// WithPageSize is the number of Transactions to include in each page
// of the results. [default=100, maximum=1000]
func (tf *transactionFunc) WithPageSize(size int) transactionOpts {
	return func(tq *transactionQuery) {
		if size < 1 || size > 1000 {
			tq.PageSize = 100
			return
		}
		tq.PageSize = size
	}
}
