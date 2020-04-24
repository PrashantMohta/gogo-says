package controllers

import (
	"net/http"
)

func RegisterControllers() {
	qc := newQuoteController()
	http.Handle("/", *qc)
	http.Handle("/gogo-says", *qc)
	http.Handle("/gogo-says/", *qc)
}
