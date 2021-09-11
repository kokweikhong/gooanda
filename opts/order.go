package opts

import (
	"strconv"
	"strings"

	"github.com/kokweikhong/gooanda/kw"
)

type orderQuery struct {
	Ids        string `json:"ids,omitempty"`
	State      string `json:"state,omitempty"`
	Instrument string `json:"instrument,omitempty"`
	Count      string `json:"count,omitempty"`
	BeforeID   string `json:"beforeID,omitempty"`
}

type OrderOpts func(*orderQuery)

func NewOrderQuery(querys ...OrderOpts) *orderQuery {
	q := &orderQuery{}
	for _, query := range querys {
		query(q)
	}
	return q
}

// integer	The maximum number of Orders to return [default=50, maximum=500]
func QOrderCount(count int) OrderOpts {
	return func(oq *orderQuery) {
		if count < 0 || count > 5000 {
			oq.Count = "500"
			return
		} else {
			oq.Count = strconv.Itoa(count)
		}
	}
}

// The state to filter the requested Orders by [default=PENDING]
func QOrderState(state kw.OrderState) OrderOpts {
	return func(oq *orderQuery) {
		oq.State = string(state)
	}
}

// The instrument to filter the requested orders by
func QOrderInstrument(instrument string) OrderOpts {
	return func(oq *orderQuery) {
		oq.Instrument = instrument
	}
}

// List of OrderID (csv)	List of Order IDs to retrieve
func QOrderListID(ids []string) OrderOpts {
	return func(oq *orderQuery) {
		oq.Ids = strings.Join(ids, ",")
	}
}

// The maximum Order ID to return. If not provided the most recent Orders in the Account are returned
func QOrderID(id string) OrderOpts {
	return func(oq *orderQuery) {
		oq.BeforeID = id
	}
}
