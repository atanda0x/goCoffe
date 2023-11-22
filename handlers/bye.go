package handlers

import (
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (b *Bye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye\n"))
}
