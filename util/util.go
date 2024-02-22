package util

import (
	"fmt"
	"sort"

	"github.com/jp-mango/mangomarkets/internal/api"
)

func PrintTimeSeriesData(provider api.TimeSeriesProvider) {
	tsData := provider.GetTimeSeriesData()

	var dates []string
	for date := range tsData {
		dates = append(dates, date)
	}
	// Sorts in ascending order
	sort.Slice(dates, func(i, j int) bool {
		return dates[i] < dates[j]
	})
	for _, date := range dates {
		fmt.Printf("Date: %s\n", date)
		// Collect keys of the inner map
		var keys []string
		for key := range tsData[date] {
			keys = append(keys, key)
		}
		// Sort the keys to ensure consistent order
		sort.Strings(keys)
		// Now iterate over the sorted keys
		for _, key := range keys {
			value := tsData[date][key]
			fmt.Printf("%s: %s\n", key, value)
		}
		fmt.Println("----------------")
	}
}

func PrintTopGainersAndLosers(marketData api.TopGainLoss) {
	fmt.Println("Top Gainers:")

	for _, gainer := range marketData.TopGainers {
		fmt.Printf("Ticker: %s, Price: %s, Change: %s%%, Volume: %s\n", gainer.Ticker, gainer.Price, gainer.ChangePercentage, gainer.Volume)
	}
	fmt.Println("\nTop Losers:")

	for _, loser := range marketData.TopLosers {
		fmt.Printf("Ticker: %s, Price: %s, Change: %s%%, Volume: %s\n", loser.Ticker, loser.Price, loser.ChangePercentage, loser.Volume)
	}
}

func PrintMarketStatus(marketStatus api.MarketHours) {
	fmt.Println(marketStatus.Endpoint)
	for _, market := range marketStatus.Markets {
		fmt.Printf("\nMarket Type: %s\nRegion: %s\nPrimary Exchanges: %s\nOpen: %s\nClose: %s\nStatus: %s\nNotes: %s\n",
			market.MarketType, market.Region, market.PrimaryExchanges, market.LocalOpen, market.LocalClose, market.CurrentStatus, market.Notes)
	}
}
