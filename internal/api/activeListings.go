package api

import (
	"fmt"
	"mangomarkets/internal"
)

func FetchActiveListings(apiKey string) (string, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=LISTING_STATUS&apikey=%s", apiKey)

	listings, err := DataPull(url)
	if err != nil {
		return "", internal.ErrDataPull(err)
	}

	return string(listings), nil
}
