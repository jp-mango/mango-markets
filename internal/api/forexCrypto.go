package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ForexTimeSeriesData interface {
	GetTimeSeriesData() map[string]map[string]string
}

type ForexMetaData struct {
	Information   string `json:"1. Information"`
	FromSymbol    string `json:"2. From Symbol"`
	ToSymbol      string `json:"3. To Symbol"`
	OutputSize    string `json:"4. Output Size"`
	LastRefreshed string `json:"5. Last Refreshed"`
	TimeZone      string `json:"6. Time Zone"`
}

// ! Daily Forex Time Series
type ForexTimeSeriesDaily struct {
	MetaData   ForexMetaData                `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Time Series FX (Daily)"`
}

func (ftsd ForexTimeSeriesDaily) GetTimeSeriesData() map[string]map[string]string {
	return ftsd.TimeSeries
}

func FetchForexTimeSeriesDaily(apiKey, fromSymbol, toSymbol string) (ForexTimeSeriesDaily, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=FX_DAILY&from_symbol=%s&to_symbol=%s&apikey=%s", fromSymbol, toSymbol, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var ftsdData ForexTimeSeriesDaily
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &ftsdData); err != nil {
		return ForexTimeSeriesDaily{}, fmt.Errorf("error parsing API content: %v", err)
	}
	return ftsdData, nil
}

// !  Weekly Forex Time Series
type ForexTimeSeriesWeekly struct {
	MetaData   ForexMetaData                `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Time Series FX (Weekly)"`
}

func (ftsw ForexTimeSeriesWeekly) GetTimeSeriesData() map[string]map[string]string {
	return ftsw.TimeSeries
}

func FetchForexTimeSeriesWeekly(apiKey, fromSymbol, toSymbol string) (ForexTimeSeriesWeekly, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=FX_WEEKLY&from_symbol=%s&to_symbol=%s&apikey=%s", fromSymbol, toSymbol, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var ftswData ForexTimeSeriesWeekly
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &ftswData); err != nil {
		return ForexTimeSeriesWeekly{}, fmt.Errorf("error parsing API content: %v", err)
	}
	return ftswData, nil
}

// Monthly Forex Time Series
type ForexTimeSeriesMonthly struct {
	MetaData   ForexMetaData                `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Time Series FX (Monthly)"`
}

func (ftsm ForexTimeSeriesMonthly) GetTimeSeriesData() map[string]map[string]string {
	return ftsm.TimeSeries
}

func FetchForexTimeSeriesMonthly(apiKey, fromSymbol, toSymbol string) (ForexTimeSeriesMonthly, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=FX_MONTHLY&from_symbol=%s&to_symbol=%s&apikey=%s", fromSymbol, toSymbol, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var ftsmData ForexTimeSeriesMonthly
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &ftsmData); err != nil {
		return ForexTimeSeriesMonthly{}, fmt.Errorf("error parsing API content: %v", err)
	}
	return ftsmData, nil
}
