package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NewsSentimentData struct {
	Items                    string `json:"items"`
	SentimentScoreDefinition string `json:"sentiment_score_definition"`
	RelevanceScoreDefinition string `json:"relevance_score_definition"`
	Feed                     []struct {
		Title                 string   `json:"title"`
		URL                   string   `json:"url"`
		TimePublished         string   `json:"time_published"`
		Authors               []string `json:"authors"`
		Summary               string   `json:"summary"`
		BannerImage           string   `json:"banner_image"`
		Source                string   `json:"source"`
		CategoryWithinSource  string   `json:"category_within_source"`
		SourceDomain          string   `json:"source_domain"`
		Topics                []Topic  `json:"topics"`
		OverallSentimentScore float64  `json:"overall_sentiment_score"`
		OverallSentimentLabel string   `json:"overall_sentiment_label"`
		TickerSentiment       []Ticker `json:"ticker_sentiment"`
	} `json:"feed"`
}

type Topic struct {
	Topic          string  `json:"topic"`
	RelevanceScore float64 `json:"relevance_score,string"`
}

type Ticker struct {
	Ticker               string  `json:"ticker"`
	RelevanceScore       float64 `json:"relevance_score,string"`
	TickerSentimentScore float64 `json:"ticker_sentiment_score,string"`
	TickerSentimentLabel string  `json:"ticker_sentiment_label"`
}

func FetchNewsSentimentData(apiKey string, ticker string) (*NewsSentimentData, error) {
	// Construct the API URL
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers=%s&apikey=%s", ticker, apiKey)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch news sentiment data: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON response
	var data NewsSentimentData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %v", err)
	}

	return &data, nil
}
