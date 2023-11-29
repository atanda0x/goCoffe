package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/atanda0x/goCoffe/data"
	"github.com/gorilla/mux"
)

type Coffee struct {
	l *log.Logger
}

func NewCoffee(l *log.Logger) *Coffee {
	return &Coffee{l}
}

func (c *Coffee) GetCoffees(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle GET Coffee")

	// fetch the coffee from the datastore
	lc := data.GetCoffees()

	// serialized the list to JSON
	err := lc.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (c *Coffee) AddCoffe(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle POST Coffee")

	coff := &data.Coffee{}

	err := coff.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddCoffe(coff)
}

func (c *Coffee) UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "unable to convert id", http.StatusBadRequest)
		return
	}

	c.l.Println("Handle PUT Product", id)

	coff := &data.Coffee{}

	err = coff.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)

	}

	err = data.UpdatedCoffee(id, coff)
	if err == data.ErrorCoffeeNotFound {
		http.Error(w, "Coffe not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Coffee not found", http.StatusInternalServerError)
		return
	}
}
