package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/kokweikhong/gooanda/endpoint"
)

func TestEndpoint(t *testing.T) {
	fmt.Println(endpoint.Account.Accounts)
	fmt.Println(endpoint.GetEndpoint(true, endpoint.Account.Accounts))
}
