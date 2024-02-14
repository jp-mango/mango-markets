package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//? loads env variables from file in root
	godotenv.Load("../../.env")

	//? check for API key
	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		log.Fatal("API_KEY is not found in the env")
	}
	fmt.Println("API key:", API_KEY)

	//? time series daily
	// retrieve ticker from user
	var ticker string
	fmt.Print("Enter ticker: ")
	fmt.Scanln(&ticker)
	// API call w/ ticker and API info
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", ticker, API_KEY)
	// GET request to API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint")
	}
	defer resp.Body.Close()
	// read contents & convert to string
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to load content")
	}
	// place content into struct
	type TimeSeriesResponse struct {
		MetaData struct {
			Information   string `json:"1. Information"`
			Symbol        string `json:"2. Symbol"`
			LastRefreshed string `json:"3. Last Refreshed"`
			OutputSize    string `json:"4. Output Size"`
			TimeZone      string `json:"5. Time Zone"`
		} `json:"Meta Data"`
		TimeSeries map[string]map[string]string `json:"Time Series (Daily)"`
	}
	// unpack response
	var tsData TimeSeriesResponse
	unmarsh_err := json.Unmarshal(content, &tsData)
	if unmarsh_err != nil {
		log.Fatal("Error pulling API content")
	}
}
