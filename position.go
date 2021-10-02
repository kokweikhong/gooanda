package gooanda

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kokweikhong/gooanda/endpoint"
)

type position struct {
	connection
}

func NewPositionConnection(token string) *position {
	conn := &position{}
	conn.token = token
	return conn
}

func (ps *position) GetPositionList(live bool, accountID string) {
	ep := endpoint.GetEndpoint(live, endpoint.Position.PositionList)
	ps.endpoint = fmt.Sprintf(ep, accountID)
	fmt.Println(ps.endpoint)
	ps.method = http.MethodGet
	data, err := ps.connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func (ps *position) GetOpenPositionList(live bool, accountID string) {
	ep := endpoint.GetEndpoint(live, endpoint.Position.OpenPositionList)
	ps.endpoint = fmt.Sprintf(ep, accountID)
	fmt.Println(ps.endpoint)
	ps.method = http.MethodGet
	data, err := ps.connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func (ps *position) GetOpenPositionForInstrument(live bool, accountID, instrument string) {
	ep := endpoint.GetEndpoint(live, endpoint.Position.SingleInstrumentPosition)
	ps.endpoint = fmt.Sprintf(ep, accountID, instrument)
	fmt.Println(ps.endpoint)
	ps.method = http.MethodGet
	data, err := ps.connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func (ps *position) CloseOpenPositionForInstrument(live bool, accountID, instrument string, isLongPosition bool, units interface{}) {
	var pos string
	if isLongPosition {
		pos = "longUnits"
	} else if !isLongPosition {
		pos = "shortUnits"
	}
	switch t := units.(type) {
	case int, float64, nil:
		_, okInt := t.(int)
		_, okFloat := t.(float64)
		if okInt && t.(int) < 1 {
			log.Fatalf("Type INT %v must be greater than 0", t)
		} else if okFloat && t.(float64) < 1 {
			log.Fatalf("Type FLOAT64 %v must be greater than 0", t)
		} else if units == nil {
			units = "ALL"
		}
	default:
		log.Fatal("only accepted types are INT, FLOAT64, NIL")
	}
	body := fmt.Sprintf(`{"%v":"%v"}`, pos, units)
	ep := endpoint.GetEndpoint(live, endpoint.Position.ClosePositionForInstrument)
	ps.endpoint = fmt.Sprintf(ep, accountID, instrument)
	ps.method = http.MethodPut
	ps.data = []byte(body)
	data, err := ps.connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
