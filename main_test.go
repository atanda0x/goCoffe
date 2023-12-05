package main

import (
	"fmt"
	"testing"

	"github.com/go-swagger/go-swagger/examples/cli/client"
	"github.com/go-swagger/go-swagger/examples/file-server/client"
)

func TestClinet(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := coffees.NewListCoffeesParams()
	prod, err := c.Coffee.ListCoffee(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
	t.Fail()
}
