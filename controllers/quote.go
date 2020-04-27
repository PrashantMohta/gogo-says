package controllers

import (
	"fmt"
	rand "math/rand"
	"net/http"
	"strings"
	"sync"

	"github.com/PrashantMohta/gogo-says/models"
)

type QuoteController struct{}

func (qc QuoteController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	println((*request.URL).Path)
	path := request.URL.Path
	var out string
	var ok bool
	response.Header().Set("Content-Type", "text/html")
	response.Write([]byte(`<h2>Welcome to Go Introduction project on Crime Master Gogo</h2> <hr/>`))
	if path == "/" {
		out = `<a href='/gogo-says'>Get a Random GoGo quote</a> <br/> <a href='/gogo-pc'>Get a Synthetic GoGo quote</a> `
		ok = true
	} else if path == "/gogo-says" {
		out, ok = getRandomQuote()
	} else if path == "/gogo-pc" {
		out, ok = getSyntheticQuote()
	}

	if ok {
		response.Write([]byte(out))
	} else {
		fmt.Println(out)
		response.Write([]byte("Something went wrong, try again later"))
	}
	response.Write([]byte(`<hr/> <a href="` + path + `">Reload</a>`))
}

func produceQuoteWords(words chan<- string, await *sync.WaitGroup) {
	// producer
	if quote, ok := getRandomQuote(); ok {
		quoteSlice := strings.Split(quote, " ")
		for i := 0; i < 3; i++ {
			word := rand.Intn(len(quoteSlice))
			words <- quoteSlice[word]
		}
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

func getSyntheticQuote() (string, bool) {
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

	if len(sentence) > 0 {
		return sentence, true
	}
	return "Something bad happened", false
}

func getRandomQuote() (string, bool) {
	if models.MaxIDs > 0 {
		randomID := rand.Intn(models.MaxIDs)
		quote, err := models.GetQuoteByID(randomID)
		if err == nil && quote != nil {
			quoteAsString := quote.Value
			if len(quoteAsString) > 0 {
				return quoteAsString, true
			}
			return "Empty Quote", false
		}
		return "Some Error Occurred", false
	}
	return "No Quotes Present", false
}

func newQuoteController() *QuoteController {
	return &QuoteController{}
}
