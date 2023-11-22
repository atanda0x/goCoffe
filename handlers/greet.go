package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Greet struct {
	l *log.Logger
}

func NewGeet(l *log.Logger) *Greet {
	return &Greet{l}
}

func (h *Greet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Hi")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oopp", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hi %s \n", d)
}
