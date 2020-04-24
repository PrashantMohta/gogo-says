package controllers

import (
	"net/http"
)

type QuoteController struct{}

func (qc QuoteController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the Server"))
}

func newQuoteController() *QuoteController {
	return &QuoteController{}
}
