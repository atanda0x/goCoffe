package handlers

import (
	"net/http"
	"strconv"

	"github.com/atanda0x/goCoffe/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (c *Coffee) DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	// This will always convert because of the router
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	c.l.Println("Handle Delete Coffee handler", id)

	err := data.DeleteCoffee(id)

	if err == data.ErrorCoffeeNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
