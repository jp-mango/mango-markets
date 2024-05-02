package api

import (
	"encoding/json"
	"fmt"
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

/*
! Market News
This API returns live and historical market news & sentiment data from a large & growing selection of premier news outlets around the world, covering stocks, cryptocurrencies, forex, and a wide range of topics such as fiscal policy, mergers & acquisitions, IPOs, etc.
*/
func FetchNewsSentimentData(apiKey string, ticker string) (*NewsSentimentData, error) {
	// Construct the API URL
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers=%s&apikey=%s", ticker, apiKey)

	var news NewsSentimentData

	newsData, err := DataPull(url)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	err = json.Unmarshal(newsData, &news)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %v", err)
	}

	return &news, nil
}
