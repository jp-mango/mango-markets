package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mangomarkets/internal/errors"
	"mangomarkets/internal/load"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

// - Custom unmarshal to handle dynamic json requests for time series data
func (t *TimeSeriesData) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	// Unmarshal Metadata
	err = json.Unmarshal(raw["Meta Data"], &t.Metadata)
	if err != nil {
		return err
	}

	// Unmarshal TimeSeries
	for key, value := range raw {
		if strings.HasPrefix(key, "Time Series") {
			err = json.Unmarshal(value, &t.TimeSeries)
			if err != nil {
				return err
			}
			t.Frequency = strings.TrimPrefix(key, "Time Series (")
			t.Frequency = strings.TrimSuffix(t.Frequency, ")")
			break
		} else if strings.HasSuffix(key, "Time Series") {
			err = json.Unmarshal(value, &t.TimeSeries)
			if err != nil {
				return err
			}
			t.Frequency = strings.TrimSuffix(key, " Time Series")
			break
		}
	}

	return nil
}

// - Pull the data from the api, allowing you to unmarshal json into desired struct
func DataPull(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("no info returned from url: %s", err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}

	if strings.Contains(string(content), "Error Message") {
		return nil, errors.ErrInvalidAPIRequest(string(content))
	}

	return content, nil
}

// Pull and unmarshal to reduce boilerplate in timeSeries.go
func pullUnmarshTSD(url string) (*TimeSeriesData, error) {
	var data TimeSeriesData

	content, err := DataPull(url)
	if err != nil {
		return nil, errors.ErrDataPull(err)
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, errors.ErrUnmarshalJSON(err)
	}

	return &data, nil
}

// Prints time series data
func PrintTimeSeries(data *TimeSeriesData) string {
	var sb strings.Builder

	keys := make([]string, 0, len(data.TimeSeries))
	for k := range data.TimeSeries {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for i, k := range keys {
		v := data.TimeSeries[k]
		sb.WriteString(fmt.Sprintf("| %d | %s | %s | Open: %s | High: %s | Low: %s | Close: %s | Volume: %s | Current Time: %s |\n", i, data.Metadata.Symbol, k, v.Open, v.High, v.Low, v.Close, v.Volume, time.Now()))
	}

	return sb.String()
}

// Cleans up user input for tickers
func SanitizeTicker(ticker string) string {
	return strings.ToUpper(strings.TrimSpace(ticker))
}

// helper for verifying user ticker input
func FoundTickerInput(ticker string) bool {
	_, _, ACTIVE_STOCKS, _ := load.Env()
	activeListings, err := os.Open(ACTIVE_STOCKS)
	if err != nil {
		slog.Error(err.Error())
	}
	defer activeListings.Close()

	reader := csv.NewReader(activeListings)

	records, err := reader.ReadAll()
	if err != nil {
		slog.Error(err.Error())
	}

	for _, eachrecord := range records {
		if eachrecord[0] == ticker {
			return true
		}
	}

	return false
}
