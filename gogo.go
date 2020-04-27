package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	port := 3000

	fmt.Println("Getting Json from Repo...")
	initialiseMap()

	fmt.Println("Registering Controller...")
	controllers.RegisterControllers()

	fmt.Println("Starting Server at http://localhost:" + strconv.Itoa(port) + "/")
	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
