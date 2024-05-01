package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TimeSeriesData struct {
	Metadata   Metadata              `json:"Meta Data"`
	TimeSeries map[string]StockPrice `json:"Time Series"`
	Frequency  string
}

type Metadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	Size          string `json:"5. Output Size"`
	TZ            string `json:"6. Time Zone"`
}

type StockPrice struct {
	Open           float64 `json:"1. open"`
	High           float64 `json:"2. high"`
	Low            float64 `json:"3. low"`
	Close          float64 `json:"4. close"`
	Volume         int64   `json:"5. volume"`
	AdjustedClose  float64 `json:"5. adjusted close"`
	DividendAmount float64 `json:"7. dividend amount"`
}

func intradayData(apiKey, ticker, interval string) (TimeSeriesData, error) {
	/*
	*	- Interval (1m,5m,15m,30m,60m)
	 */
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=%s&entitlement=delayed&apikey=%s", ticker, interval, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return TimeSeriesData{}, fmt.Errorf("no info returned from url: %s", err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return TimeSeriesData{}, fmt.Errorf("error reading response: %s", err)
	}

	var data TimeSeriesData

	err = json.Unmarshal(content, &data)
	if err != nil {
		return TimeSeriesData{}, fmt.Errorf("error unmarshaling JSON: %s", err)
	}

	return data, nil
}
