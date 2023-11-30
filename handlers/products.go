// Package classification of Coffee Product API
//
// Documentation for Coffee Product API
//
// SCHEMA: http
// BasePath: /Coffee
// Version: 1.0.0
//
//Consumes:
// - application/json
//
// Produces:
// - application/json
//swagger:meta

package handlers

import (
	"context"
	"fmt"
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

	coff := r.Context().Value(KeyCoffee{}).(data.Coffee)
	data.AddCoffe(&coff)
}

func (c *Coffee) UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "unable to convert id", http.StatusBadRequest)
		return
	}

	c.l.Println("Handle PUT Product", id)
	coff := r.Context().Value(KeyCoffee{}).(data.Coffee)

	err = data.UpdatedCoffee(id, &coff)
	if err == data.ErrorCoffeeNotFound {
		http.Error(w, "Coffe not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Coffee not found", http.StatusInternalServerError)
		return
	}
}

type KeyCoffee struct{}

func (c *Coffee) MiddlewareCoffeeValid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		coff := data.Coffee{}

		err := coff.FromJSON(r.Body)
		if err != nil {
			c.l.Println("[ERROR] deserialing coffee", err)
			http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		err = coff.Validate()
		if err != nil {
			c.l.Println("[ERROR] validating coffee", err)
			http.Error(w, fmt.Sprintf("Error reading coffee: %s", err), http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), KeyCoffee{}, coff)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})

}
