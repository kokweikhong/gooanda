package gooanda

import (
	"testing"
)

const (
	token = "f46a82f6b2f1bbf41f3596f05ca2a150-ef7e29de37746ef0cdb155f003840b15"
	id    = "001-011-3698178-001"
)

func TestGetAccounts(t *testing.T) {
	t.Parallel()
	GetInstrumentPositionBook(token, "USD_JPY")
	// GetInstrumentCandles(token, "USD_JPY",
	//	QInFromTo(time.Now().AddDate(0, 0, -2), time.Now()),
	//	QInGranularity(kw.Granularity_D))
	// GetAccountInstruments(token, id, QAccountInstruments([]string{"EUR_USD", "USD_JPY"}))
	// GetAccounts(token)
}
