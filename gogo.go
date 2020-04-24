package main

import (
	"fmt"
	"net/http"

	"github.com/PrashantMohta/gogo-says/controllers"

	"github.com/PrashantMohta/gogo-says/models"
)

func main() {
	foo := models.Quote{ID: 1, Value: "crime master gogo"}
	models.AddQuote(foo)
	models.AddQuote(models.Quote{ID: 2, Value: "crime master gogo"})
	models.AddQuote(models.Quote{ID: 3, Value: "crime master gogo"})

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

	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
