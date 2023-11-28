package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPut {
		c.l.Println("PUT", r.URL.Path)

		// expect the id int URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			c.l.Println("Invalid URI more than one id")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			c.l.Println("Invalid URI more than one capture group")
			http.Error(w, "Ivaid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.l.Println("Invalid URI unable to convert to number", idString)
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}

		c.updateCoffee(id, w, r)
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

func (c *Coffee) updateCoffee(id int, w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle PUT Product")

	coff := &data.Coffee{}

	err := coff.FromJSON(r.Body)
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
