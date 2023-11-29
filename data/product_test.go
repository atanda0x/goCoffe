package data

import (
	"log"
	"testing"
)

func TestValidate(t *testing.T) {
	c := &Coffee{
		Name:  "litt",
		Price: 1.00,
		SKU:   "abs-all-gds",
		Ratio: "1 shot of Latte",
		Cup:   "2-4 oz",
	}

	err := c.Validate()
	if err != nil {
		log.Fatal(err)
	}
}
