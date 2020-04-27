package controllers

import (
	rand "math/rand"
	"net/http"
	"sync"

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

func produceQuoteWords(words chan<- string, await *sync.WaitGroup) {
	// producer
	for _, v := range []string{"1", "2", "3"} {
		words <- "string " + v
	}

	await.Done()
}

//combine Words To Make Gibberish Sentences
func combineWords(sentence *string, mutex *sync.Mutex, words <-chan string, await *sync.WaitGroup) {
	// consumer
	for word := range words {
		mutex.Lock()
		*sentence += word + " "
		mutex.Unlock()
	}

	await.Done()
}

func getSyntheticQuote() string {
	jobs := make(chan string)
	mutex := sync.Mutex{}
	consumerWait := sync.WaitGroup{}
	producerWait := sync.WaitGroup{}

	var sentence string
	for i := 0; i < 3; i++ {
		producerWait.Add(1)
		go produceQuoteWords(jobs, &producerWait)
	}
	for i := 0; i < 3; i++ {
		consumerWait.Add(1)
		go combineWords(&sentence, &mutex, jobs, &consumerWait)
	}

	producerWait.Wait()
	close(jobs)
	consumerWait.Wait()

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
