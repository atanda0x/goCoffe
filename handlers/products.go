package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/atanda0x/goCoffe/data"
)

// @A list of Coffee returns in the response
// @swagger:response coffeesResponse
type coffeeResponse struct {
	// @All coffees in the system
	// @in: body
	Body []data.Coffee
}

// @swagger:parameters deleteCoffee
type coffeeIDParameterWrapper struct {
	// @The id of the product to delete from the db
	// @in: path
	// @required: true
	ID int `json:"id"`
}

// Coffee is a http.Handler
type Coffee struct {
	l *log.Logger
}

// NewCoffee create a products handler with the given logger
func NewCoffee(l *log.Logger) *Coffee {
	return &Coffee{l}
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
