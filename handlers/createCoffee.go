package handlers

import (
	"net/http"

	"github.com/atanda0x/goCoffe/data"
)

func (c *Coffee) AddCoffe(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle POST Coffee")

	coff := r.Context().Value(KeyCoffee{}).(data.Coffee)
	data.AddCoffe(&coff)
}
