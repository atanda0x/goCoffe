package handlers

import (
	"net/http"
	"strconv"

	"github.com/atanda0x/goCoffe/data"
	"github.com/gorilla/mux"
)

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
