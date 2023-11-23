package handlers

import (
	"log"
	"net/http"

	"github.com/atanda0x/goCoffe/data"
)

type Coffee struct {
	l *log.Logger
}

func NewCoffee(l *log.Logger) *Coffee {
	return &Coffee{l}
}

func (c *Coffee) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.GetCoffees(w, r)
		return
	}

	if r.Method == http.MethodPost {
		c.AddCoffe(w, r)
		return
	}

	// catch all
	// If no method is satisfied return an error
	w.WriteHeader(http.StatusMethodNotAllowed)
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
