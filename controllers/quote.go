package controllers

import (
	rand "math/rand"
	"net/http"

	"github.com/PrashantMohta/gogo-says/models"
)

type QuoteController struct{}

func (qc QuoteController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	println((*request.URL).Path)
	path := request.URL.Path
	if path == "/gogo-says" {
		response.Write([]byte(getRandomQuote()))
	} else if path == "/gogo-pc" {
		response.Write([]byte(getSyntheticQuote()))
	}
}

func getSyntheticQuote() string {
	jobs := make(chan string)
	done := make(chan bool)

	var sentence string
	go (func(jobs chan<- string) {
		// producer
		for _, v := range []string{"1", "2", "3"} {
			jobs <- "string " + v
		}

		close(jobs)

	})(jobs)

	go (func(jobs <-chan string, completion chan<- bool) {
		// consumer
		for word := range jobs {
			sentence += word + " "
		}

		completion <- true
	})(jobs, done)

	<-done
	return sentence
}

func getRandomQuote() string {
	if models.MaxIDs > 0 {
		randomID := rand.Intn(models.MaxIDs)
		quote, err := models.GetQuoteByID(randomID)
		if err == nil && quote != nil {
			quoteAsString := quote.Value
			if len(quoteAsString) > 0 {
				return quoteAsString
			} else {
				return "Empty Quote"
			}
		} else {
			return "Some Error Occurred"
		}
	} else {
		return "No Quotes Present"
	}
}

func newQuoteController() *QuoteController {
	return &QuoteController{}
}
