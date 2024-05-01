package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ! Custom unmarshal to handle dynamic json requests
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

// ! Pull the data from the api
func dataPull(url string) (*TimeSeriesData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("no info returned from url: %s", err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}

	var data TimeSeriesData

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %s", err)
	}

	return &data, nil
}
