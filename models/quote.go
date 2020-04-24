package models

import (
	"errors"
)

//Quote struct
type Quote struct {
	ID    int
	Value string
}

var quoteMap = make(map[int]*Quote)
var MaxIDs = 0

//AddQuote will add a quote to the map
func AddQuote(q Quote) error {
	quoteMap[q.ID] = &q
	MaxIDs++
	return nil
}

//GetQuotes will get the quotes list from the map
func GetQuotes() ([]*Quote, error) {
	var quotes []*Quote
	for _, v := range quoteMap {
		quotes = append(quotes, v)
	}
	if len(quotes) == 0 {
		return quotes, errors.New("Quotes List is Empty")
	}
	return quotes, nil
}

//GetQuoteByID will get a quote from the map based on it's ID
func GetQuoteByID(id int) (*Quote, error) {
	quote := quoteMap[id]
	return quote, nil
}
