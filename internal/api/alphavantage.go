package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ! Time series data
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

func FetchTimeSeriesDaily(apiKey, ticker string) TimeSeriesResponse {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var tsData TimeSeriesResponse
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &tsData); err != nil {
		log.Fatal("Error parsing API content:", err)
	}

	return tsData
}

//
