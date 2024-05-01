package api

import (
	"fmt"
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
