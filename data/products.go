package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the sttructure of the coffee API
type Coffee struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

type Coffees []*Coffee

func (c *Coffees) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func GetCoffees() Coffees {
	return CoffeeList
}

var CoffeeList = []*Coffee{
	&Coffee{
		ID:          1,
		Name:        "Latte",
		Description: "From milky cofFee",
		Price:       3.45,
		SKU:         "abc323",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},

	&Coffee{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.90,
		SKU:         "def34",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
