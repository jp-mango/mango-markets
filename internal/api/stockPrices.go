package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TimeSeriesProvider is an interface for fetching time series data.
type TimeSeriesProvider interface {
	GetTimeSeriesData() map[string]map[string]string
}

// ! Time series daily
func (tsd TimeSeriesDaily) GetTimeSeriesData() map[string]map[string]string {
	return tsd.TimeSeries
}

type TimeSeriesDaily struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		OutputSize    string `json:"4. Output Size"`
		TimeZone      string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Time Series (Daily)"`
}

func FetchTimeSeriesDaily(apiKey, ticker string) (TimeSeriesDaily, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var tsdData TimeSeriesDaily
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &tsdData); err != nil {
		return TimeSeriesDaily{}, fmt.Errorf("error parsing API content:%e", err)
	}
	return tsdData, err
}

// ! Time series weekly
func (tsw TimeSeriesWeekly) GetTimeSeriesData() map[string]map[string]string {
	return tsw.TimeSeries
}

type TimeSeriesWeekly struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		TimeZone      string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Weekly Adjusted Time Series"`
}

func FetchTimeSeriesWeekly(apiKey, ticker string) (TimeSeriesWeekly, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY_ADJUSTED&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var tswData TimeSeriesWeekly
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &tswData); err != nil {
		return TimeSeriesWeekly{}, fmt.Errorf("error parsing API content:%e", err)
	}
	return tswData, err
}

// ! Time series monthly
func (tsm TimeSeriesMonthly) GetTimeSeriesData() map[string]map[string]string {
	return tsm.TimeSeries
}

type TimeSeriesMonthly struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		TimeZone      string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Monthly Adjusted Time Series"`
}

func FetchTimeSeriesMonthly(apiKey, ticker string) (TimeSeriesMonthly, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY_ADJUSTED&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var tsmData TimeSeriesMonthly
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &tsmData); err != nil {
		return TimeSeriesMonthly{}, fmt.Errorf("error parsing API content:%e", err)
	}
	return tsmData, err
}

// ! Top Gainers and Losers
type TopGainLoss struct {
	Metadata           string   `json:"metadata"`
	LastUpdated        string   `json:"last_updated"`
	TopGainers         []Ticker `json:"top_gainers"`
	TopLosers          []Ticker `json:"top_losers"`
	MostActivelyTraded []Ticker `json:"most_actively_traded"`
}

type Ticker struct {
	Ticker           string `json:"ticker"`
	Price            string `json:"price"`
	ChangeAmount     string `json:"change_amount"`
	ChangePercentage string `json:"change_percentage"`
	Volume           string `json:"volume"`
}

func FetchGainLoss(apiKey string) (*TopGainLoss, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TOP_GAINERS_LOSERS&apikey=%v", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var marketData TopGainLoss
	err = json.Unmarshal(body, &marketData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &marketData, nil
}
