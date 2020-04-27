package controllers

import (
	"net/http"
)

func RegisterControllers() {
	qc := newQuoteController()
	http.Handle("/", *qc)
}
