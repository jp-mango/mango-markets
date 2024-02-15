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
	//sorts in descending order
	sort.Slice(dates, func(i, j int) bool {
		return dates[i] > dates[j]
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
