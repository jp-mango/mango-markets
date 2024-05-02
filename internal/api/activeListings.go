package api

import (
	"fmt"
	"mangomarkets/internal"
)

/*
! Listing Status:
This API returns a list of active or delisted US stocks and ETFs, either as of the latest trading day or at a specific time in history. The endpoint is positioned to facilitate equity research on asset lifecycle and survivorship.
*/
func FetchActiveListings(apiKey string) (string, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=LISTING_STATUS&apikey=%s", apiKey)

	listings, err := DataPull(url)
	if err != nil {
		return "", internal.ErrDataPull(err)
	}

	return string(listings), nil
}
