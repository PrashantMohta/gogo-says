package main

import (
	"net/http"

	"github.com/PrashantMohta/gogo-says/controllers"

	"github.com/PrashantMohta/gogo-says/models"
)

func main() {

	models.AddQuote(models.Quote{ID: 0, Value: "Teja! yeh kaise aadmi tune paal rakhe hain, saale suit to 10–10 hazaar ka pehente hain, lekin akal 10 paise ki bhi nahi hai"})
	models.AddQuote(models.Quote{ID: 1, Value: "Crime Master Gogo naam hai mera! Ankhen nikal ke gotiyaan kheltaan hoon."})
	models.AddQuote(models.Quote{ID: 2, Value: "Aayan hoon, kuch to le ke jaoonga! Khandani chor hoon, Mogambo ka bhateeja"})
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
