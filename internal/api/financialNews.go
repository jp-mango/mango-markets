package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type News struct {
	Items             string        `json:"items"`
	SentimentScoreDef string        `json:"sentiment_score_definition"`
	RelevanceScoreDef string        `json:"relevance_score_definition"`
	Feed              []FeedContent `json:"feed"`
}

type FeedContent struct {
	Title           string     `json:"title"`
	Url             string     `json:"url"`
	PublishDate     string     `json:"time_published"`
	Authors         []string   `json:"authors"`
	Summary         string     `json:"summary"`
	Image           string     `json:"banner_image"`
	Source          string     `json:"source"`
	FeedSource      string     `json:"category_within_source"`
	SourceDomain    string     `json:"source_domain"`
	Topics          []Topic    `json:"topics"`
	SentimentScore  float64    `json:"overall_sentiment_score"`
	SentimentLabel  string     `json:"overall_sentiment_label"`
	TickerSentiment []TickSent `json:"ticker_sentiment"`
}

type Topic struct {
	Topic     string `json:"topic"`
	Relevance string `json:"relevance_score"`
}

type TickSent struct {
	Ticker               string `json:"ticker"`
	Relevance            string `json:"relevance_score"`
	TickerSentimentScore string `json:"ticker_sentiment_score"`
	TickerSentimentLabel string `json:"ticker_sentiment_label"`
}

func FetchNewsByTicker(ticker, apiKey string) (*News, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	var news News
	if err := json.Unmarshal(body, &news); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}
	return &news, nil
}

func FetchNewsByTopic(topicsUrl, apiKey string) (*News, error) {
	resp, err := http.Get(topicsUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	var news News
	if err := json.Unmarshal(body, &news); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}
	return &news, nil
}
