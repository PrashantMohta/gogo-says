package controllers

import (
	rand "math/rand"
	"net/http"
	"strconv"

	"github.com/PrashantMohta/gogo-says/models"
)

type QuoteController struct{}

func (qc QuoteController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	//println((*request.URL).Path)
	if models.MaxIDs > 0 {
		randomID := rand.Intn(models.MaxIDs)
		response.Write([]byte(strconv.Itoa(randomID)))
		quote, err := models.GetQuoteByID(randomID)
		if err == nil && quote != nil {
			quoteAsString := quote.Value
			if len(quoteAsString) > 0 {
				response.Write([]byte(quoteAsString))
			} else {
				response.Write([]byte("Empty Quote"))
			}
		} else {
			response.Write([]byte("Some Error Occurred"))
		}
	} else {
		response.Write([]byte("No Quotes Present"))
	}
}

func newQuoteController() *QuoteController {
	return &QuoteController{}
}
