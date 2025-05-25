package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BitcoinData struct {
	Time struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	Bpi        struct {
		USD struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"USD"`
		GBP struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"GBP"`
		EUR struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func main() {

	// := ist eine kürzere versione Variablen zu deklarieren Muss kein var schreiben

	url := "https://api.coindesk.com/v1/bpi/currentprice.json"

	client := &http.Client{}

	//Bei go wird oft err als zweite Variable gesetzt für Errors
	res, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
		return // Bricht hier den code ab
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fehler beim Lesen der Antwort:", err)
		return
	}
	fmt.Println(string(body)) // Gibt den Body als String aus

	var bitcoinData BitcoinData
	json.Unmarshal([]byte(body), &bitcoinData)

	fmt.Println("Bitcoin Daten vom: ", bitcoinData.Time.Updated)
	fmt.Println("BitcoinData Kurs in EUR", bitcoinData.Bpi.EUR.Rate)

}
