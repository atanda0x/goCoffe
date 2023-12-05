package handlers

import (
	"net/http"

	"github.com/atanda0x/goCoffe/data"
)

// @swagger:route GET /Coffee/get coffees listCoffees
// @Returns a list of coffees
// @responses:
// 	@200: productsResponse

// GetCoffee returns the coffee from the data db
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
