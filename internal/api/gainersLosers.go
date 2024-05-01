package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top gainers and losers: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var gainLoss TopGainLoss
	err = json.Unmarshal(body, &gainLoss)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON request: %v", err)
	}

	return &gainLoss, nil
}
