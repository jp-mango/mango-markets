package api

import (
	"encoding/json"
	"fmt"
)

type TopGainLoss struct {
	Metadata    string    `json:"metadata"`
	LastUpdated string    `json:"last_updated"`
	TopGainers  []Tickers `json:"top_gainers"`
	TopLosers   []Tickers `json:"top_losers"`
	MostActive  []Tickers `json:"most_actively_traded"`
}

type Tickers struct {
	Ticker           string `json:"ticker"`
	Price            string `json:"price"`
	ChangeAmount     string `json:"change_amount"`
	ChangePercentage string `json:"change_percentage"`
	Volume           string `json:"volume"`
}

func FetchTopGainLossData(apiKey string) (*TopGainLoss, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TOP_GAINERS_LOSERS&apikey=%s", apiKey)

	var topGainLoss TopGainLoss

	gainLoss, err := DataPull(url)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	err = json.Unmarshal(gainLoss, &topGainLoss)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON request: %v", err)
	}

	return &topGainLoss, nil
}
