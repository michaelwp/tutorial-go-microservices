package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello)ServeHttp(w http.ResponseWriter, r http.Request)  {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Ooos", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", d)
}
