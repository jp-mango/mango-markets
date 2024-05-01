package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TimeSeriesData struct {
	Metadata   Metadata              `json:"Meta Data"`
	TimeSeries map[string]StockPrice `json:"-"`
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
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

// ! Custom unmarshal to handle dynamic json requests
func (t *TimeSeriesData) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	// Unmarshal Metadata
	err = json.Unmarshal(raw["Meta Data"], &t.Metadata)
	if err != nil {
		return err
	}

	// Unmarshal TimeSeries
	for key, value := range raw {
		if strings.HasPrefix(key, "Time Series") {
			err = json.Unmarshal(value, &t.TimeSeries)
			if err != nil {
				return err
			}
			t.Frequency = strings.TrimPrefix(key, "Time Series (")
			t.Frequency = strings.TrimSuffix(t.Frequency, ")")
			break
		} else if strings.HasSuffix(key, "Time Series") {
			err = json.Unmarshal(value, &t.TimeSeries)
			if err != nil {
				return err
			}
			t.Frequency = strings.TrimSuffix(key, " Time Series")
			break
		}
	}

	return nil
}

// ! Pull the data from the api
func dataPull(url string) (*TimeSeriesData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("no info returned from url: %s", err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}

	var data TimeSeriesData

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %s", err)
	}

	return &data, nil
}

func IntradayDataPull(apiKey, ticker, interval string) (*TimeSeriesData, error) {
	/*
	*	- Interval (1m,5m,15m,30m,60m)
	 */
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=%s&entitlement=delayed&apikey=%s", ticker, interval, apiKey)
	return dataPull(url)
}

func DailyDataPull(apiKey, ticker string) (*TimeSeriesData, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", ticker, apiKey)
	return dataPull(url)
}

func WeeklyDataPull(apiKey, ticker string) (*TimeSeriesData, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=%s&apikey=%s", ticker, apiKey)
	return dataPull(url)
}

func MonthlyDataPull(apiKey, ticker string) (*TimeSeriesData, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY&symbol=%s&apikey=%s", ticker, apiKey)
	return dataPull(url)
}
