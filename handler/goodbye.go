package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type goodbye struct {
}

func NewBye() *goodbye {
	return &goodbye{}
}

func (h *goodbye) ServeHTTP(W http.ResponseWriter, R *http.Request) {
	log.Printf("Bye Micros")
	dataByte, err := io.ReadAll(R.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Fprintf(W, "Bye Micros\n")
	log.Printf("from Bye with %s", dataByte)
}
