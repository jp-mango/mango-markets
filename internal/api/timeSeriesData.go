package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TimeSeriesData interface {
	GetTimeSeriesData() map[string]map[string]string
}

func ParseFloat(value string) float64 {
	price, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("Error parsing float value: %v", err)
	}
	return price
}

func ParseInt(value string) int64 {
	volume, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatalf("Error parsing int value: %v", err)
	}
	return volume
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

type StockPrice struct {
	Symbol    string  `bson:"symbol"`
	Timestamp string  `bson:"timestamp"`
	Open      float64 `bson:"open"`
	High      float64 `bson:"high"`
	Low       float64 `bson:"low"`
	Close     float64 `bson:"close"`
	Volume    int64   `bson:"volume"`
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

func SaveTimeSeriesDaily(apiKey, ticker string, collection *mongo.Collection) error {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to hit endpoint: %v", err)
	}
	defer resp.Body.Close()

	var tsdData TimeSeriesDaily
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &tsdData); err != nil {
		return fmt.Errorf("error parsing API content: %v", err)
	}

	for date, priceData := range tsdData.TimeSeries {
		stockPrice := StockPrice{
			Symbol:    ticker,
			Timestamp: date,
			Open:      ParseFloat(priceData["1. open"]),
			High:      ParseFloat(priceData["2. high"]),
			Low:       ParseFloat(priceData["3. low"]),
			Close:     ParseFloat(priceData["4. close"]),
			Volume:    ParseInt(priceData["5. volume"]),
		}

		_, err := collection.InsertOne(context.Background(), stockPrice)
		if err != nil {
			return fmt.Errorf("error inserting stock price data: %v", err)
		}

		_, err = collection.UpdateOne(
			context.Background(),
			bson.M{"symbol": stockPrice.Symbol, "timestamp": stockPrice.Timestamp},
			bson.M{"$set": stockPrice},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			return fmt.Errorf("error updating MongoDB document: %v", err)
		}
	}

	return nil
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
		return TimeSeriesWeekly{}, fmt.Errorf("error parsing API content: %v", err)
	}
	return tswData, err
}

func SaveTimeSeriesWeekly(apiKey, ticker string, collection *mongo.Collection) error {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var tswData TimeSeriesWeekly
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &tswData); err != nil {
		return fmt.Errorf("error parsing API content: %v", err)
	}

	for date, priceData := range tswData.TimeSeries {
		stockPrice := StockPrice{
			Symbol:    ticker,
			Timestamp: date,
			Open:      ParseFloat(priceData["1. open"]),
			High:      ParseFloat(priceData["2. high"]),
			Low:       ParseFloat(priceData["3. low"]),
			Close:     ParseFloat(priceData["4. close"]),
			Volume:    ParseInt(priceData["5. volume"]),
		}

		_, err := collection.InsertOne(context.Background(), stockPrice)
		if err != nil {
			return fmt.Errorf("error inserting stock price data: %v", err)
		}
		_, err = collection.InsertOne(context.Background(), stockPrice)
		if err != nil {
			return fmt.Errorf("error inserting stock price data: %v", err)
		}
	}

	return nil
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
		return TimeSeriesMonthly{}, fmt.Errorf("error parsing API content: %v", err)
	}
	return tsmData, err
}

func SaveTimeSeriesMonthly(apiKey, ticker string, collection *mongo.Collection) error {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to hit endpoint: %v", err)
	}
	defer resp.Body.Close()

	var tsmData TimeSeriesMonthly
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &tsmData); err != nil {
		return fmt.Errorf("error parsing API content: %v", err)
	}

	for date, priceData := range tsmData.TimeSeries {
		stockPrice := StockPrice{
			Symbol:    ticker,
			Timestamp: date,
			Open:      ParseFloat(priceData["1. open"]),
			High:      ParseFloat(priceData["2. high"]),
			Low:       ParseFloat(priceData["3. low"]),
			Close:     ParseFloat(priceData["4. close"]),
			Volume:    ParseInt(priceData["5. volume"]),
		}

		_, err := collection.UpdateOne(
			context.Background(),
			bson.M{"symbol": stockPrice.Symbol, "timestamp": stockPrice.Timestamp},
			bson.M{"$set": stockPrice},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			return fmt.Errorf("error updating MongoDB document: %v", err)
		}
	}

	return nil
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
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TOP_GAINERS_LOSERS&apikey=%s", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	var marketData TopGainLoss
	err = json.Unmarshal(body, &marketData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %s", err)
	}

	return &marketData, nil
}

// ! Global Market Hours
type MarketHours struct {
	Endpoint string   `json:"endpoint"`
	Markets  []Market `json:"markets"`
}

type Market struct {
	MarketType       string `json:"market_type"`
	Region           string `json:"region"`
	PrimaryExchanges string `json:"primary_exchanges"`
	LocalOpen        string `json:"local_open"`
	LocalClose       string `json:"local_close"`
	CurrentStatus    string `json:"current_status"`
	Notes            string `json:"notes"`
}

func FetchMarketStatus(apiKey string) (*MarketHours, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=MARKET_STATUS&apikey=%s", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	var marketStatus MarketHours
	if err := json.Unmarshal(body, &marketStatus); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}

	return &marketStatus, nil
}
