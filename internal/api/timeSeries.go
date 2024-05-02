package api

import (
	"encoding/json"
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

/*
! Intraday Time Series Data
This API returns current and 20+ years of historical intraday OHLCV time series of the equity specified, covering extended trading hours where applicable (e.g., 4:00am to 8:00pm Eastern Time for the US market). You can query both raw (as-traded) and split/dividend-adjusted intraday data from this endpoint. The OHLCV data is sometimes called "candles" in finance literature.
*/
func FetchIntradayTSData(apiKey, ticker, interval string) (*TimeSeriesData, error) {
	/*
	*	- Interval (1m,5m,15m,30m,60m)
	 */
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&outputsize=full&interval=%s&entitlement=delayed&apikey=%s", ticker, interval, apiKey)

	var data TimeSeriesData

	content, err := DataPull(url)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)

	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %s", err)
	}

	return &data, nil
}

/*
! Daily Time Series Data
This API returns raw (as-traded) daily time series (date, daily open, daily high, daily low, daily close, daily volume) of the global equity specified, covering 20+ years of historical data. The OHLCV data is sometimes called "candles" in finance literature.
*/
func FetchDailyTSData(apiKey, ticker string) (*TimeSeriesData, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&outputsize=full&apikey=%s", ticker, apiKey)

	var data TimeSeriesData

	content, err := DataPull(url)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)

	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %s", err)
	}

	return &data, nil
}

/*
! Weekly Time Series Data
This API returns weekly time series (last trading day of each week, weekly open, weekly high, weekly low, weekly close, weekly volume) of the global equity specified, covering 20+ years of historical data.
*/
func FetchWeeklyTSData(apiKey, ticker string) (*TimeSeriesData, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=%s&outputsize=full&apikey=%s", ticker, apiKey)

	var data TimeSeriesData

	content, err := DataPull(url)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)

	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %s", err)
	}

	return &data, nil
}

/*
! Monthly Time Series Data
This API returns monthly time series (last trading day of each month, monthly open, monthly high, monthly low, monthly close, monthly volume) of the global equity specified, covering 20+ years of historical data.
*/
func FetchMonthlyTSData(apiKey, ticker string) (*TimeSeriesData, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY&symbol=%s&outputsize=full&apikey=%s", ticker, apiKey)

	var data TimeSeriesData

	content, err := DataPull(url)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)

	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %s", err)
	}

	return &data, nil
}
