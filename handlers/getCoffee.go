package handlers

import (
	"net/http"

	"github.com/atanda0x/goCoffe/data"
)

// @sroute GET /Coffee/get coffees listCoffees
// @Returns a list of coffees
// @responses:
// 	@200: coffeesResponse

// GetCoffee returns the coffee from the data db
func (c *Coffee) GetCoffees(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle GET Coffee")

	w.Header().Add("Content-Type", "application/json")

	// fetch the coffee from the datastore
	lc := data.GetCoffees()

	// serialized the list to JSON
	err := lc.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}
