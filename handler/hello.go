package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type hello struct {
}

func NewHello() *hello {
	return &hello{}
}

func (h *hello) ServeHTTP(W http.ResponseWriter, R *http.Request) {
	log.Printf("Hello Micros")
	dataByte, err := io.ReadAll(R.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Fprintf(W, "hii Micros\n")
	log.Printf("from Hi with %s", dataByte)
}
