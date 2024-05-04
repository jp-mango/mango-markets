package api

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"mangomarkets/internal/errors"
	"mangomarkets/internal/load"
	"os"
)

/*
! Listing Status:
This API returns a list of active or delisted US stocks and ETFs, either as of the latest trading day or at a specific time in history. The endpoint is positioned to facilitate equity research on asset lifecycle and survivorship.
*/

func FetchActiveListings(apiKey string) (map[string]string, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=LISTING_STATUS&apikey=%s", apiKey)

	listings, err := DataPull(url)
	if err != nil {
		return nil, errors.ErrDataPull(err)
	}

	_, _, stockLocation, err := load.Env()
	if err != nil {
		slog.Error(err.Error())
	}

	err = os.WriteFile(stockLocation, []byte(listings), 0644)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info(fmt.Sprintf("wrote listings to %s", stockLocation))

	activeListings, err := os.Open(stockLocation)
	if err != nil {
		slog.Error(err.Error())
	}
	defer activeListings.Close()

	reader := csv.NewReader(activeListings)

	records, err := reader.ReadAll()
	if err != nil {
		slog.Error(err.Error())
	}

	activeTickers := make(map[string]string)
	for _, eachrecord := range records {
		if len(eachrecord) >= 2 {
			ticker := eachrecord[0]
			companyName := eachrecord[1]
			activeTickers[ticker] = companyName
		}
	}

	return activeTickers, nil
}
