package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Exchange struct {
	ExchangeId string `json:"exchange_id"`
	Name string `json:"name"`
	Website string `json:"website"`
}

type Exchanges []Exchange

func main() {
	exchanges, err := fetchExchanges()
	if err != nil {
		log.Fatalln("Error when fetching exchanges", err)
	}

	report, err := template.New("exchanges").Parse(templ)
	if err != nil {
		log.Fatalln("Error when parsing template", err)
	}

	output, err := os.Create("output.html")
	if err != nil {
		log.Fatalln("Error when opening output file", err)
	}
	
	report.Execute(output, exchanges)
}

func fetchExchanges() (exchanges Exchanges, err error) {
	request, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/exchanges", nil)
	if err != nil {
		return 
	}
	request.Header.Add("X-CoinAPI-Key", "F5143398-66F6-4B6E-A39E-43C09BA01896")

	response, err := new(http.Client).Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return exchanges, errors.New("the API returned an unsuccessful response code")
	}

	json.NewDecoder(response.Body).Decode(&exchanges)
	return
}