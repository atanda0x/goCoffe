package main

// import (
// 	"fmt"
// 	"testing"

// )

// func TestClinet(t *testing.T) {
// 	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
// 	c := client.NewHTTPClientWithConfig(nil, cfg)

// 	params := coffees.NewListCoffeesParams()
// 	prod, err := c.Coffee.ListCoffee(params)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Printf("%#v", prod.GetPayload()[0])
// 	t.Fail()
// }
