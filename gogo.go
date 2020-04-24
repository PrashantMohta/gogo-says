package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PrashantMohta/gogo-says/controllers"

	"github.com/PrashantMohta/gogo-says/models"
)

func initialiseMap() {
	resp, err := http.Get("https://raw.githubusercontent.com/PrashantMohta/gogo-says/master/data/gogo.json")
	var quotes = []models.Quote{}
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			json.Unmarshal(body, &quotes)
		}
	}
	for _, v := range quotes {
		models.AddQuote(v)
	}
}

func main() {
	initialiseMap()
	/*
		quotes, err := models.GetQuotes()
		if err == nil {
			for _, v := range quotes {
				fmt.Println(*v)
			}
			fmt.Println(quotes)
		}

		quote1, err := models.GetQuoteByID(1)
		if err == nil {
			fmt.Println(*quote1)
		}
	*/

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
