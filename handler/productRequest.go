package handler

import (
	"encoding/json"
	"io"
	"lcorequest/data"
	"log"
	"net/http"
)

type pr struct {
}

func Newpr() *pr {
	return &pr{}
}

func (h *pr) ServeHTTP(W http.ResponseWriter, R *http.Request) {

	if R.Method == http.MethodGet {
		h.getProduct(W, R)
		return
	}
	if R.Method == http.MethodPost {
		h.setProduct(W, R)
		return
	}
	log.Printf("pr Not get method")
	return
}

func (h *pr) getProduct(W http.ResponseWriter, R *http.Request) {
	productList := data.NewProduct()
	data, _ := json.Marshal(productList)
	W.Write(data)
}

func (h *pr) setProduct(W http.ResponseWriter, R *http.Request) {
	productList := data.NewProduct()
	data, _ := io.ReadAll(R.Body)
	e := json.Unmarshal(data, productList[0])
	if e != nil {

		panic(e)
	}
	h.getProduct(W, R)
}
