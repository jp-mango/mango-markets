package util

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jp-mango/mangomarkets/internal/api"
)

func PrintTimeSeriesData(provider api.TimeSeriesProvider) string {
	tsData := provider.GetTimeSeriesData()

	var builder strings.Builder

	var dates []string
	for date := range tsData {
		dates = append(dates, date)
	}
	// Sorts in descending order
	sort.Slice(dates, func(i, j int) bool {
		return dates[i] > dates[j]
	})

	for _, date := range dates {
		builder.WriteString(fmt.Sprintf("Date: %s\n", date))
		// Collect keys of the inner map
		var keys []string
		for key := range tsData[date] {
			keys = append(keys, key)
		}
		// Sort the keys to ensure consistent order and iterate
		sort.Strings(keys)
		for _, key := range keys {
			value := tsData[date][key]
			builder.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}
		builder.WriteString("----------------\n")
	}

	return builder.String()
}

func FormatTimeSeriesData(provider api.TimeSeriesProvider) []string {
	data := provider.GetTimeSeriesData()

	// We need to sort the dates to display them in order.
	var dates []string
	for date := range data {
		dates = append(dates, date)
	}

	// Sorting dates in descending order. You might adjust sorting depending on your preference.
	sort.Slice(dates, func(i, j int) bool {
		return dates[i] > dates[j]
	})

	var formattedData []string
	for _, date := range dates {
		// Adding the date as the first line for each entry.
		formattedData = append(formattedData, fmt.Sprintf("Date: %s", date))

		// Ensure the metrics are displayed in a consistent order.
		// Assuming the metrics are "open", "high", "low", "close", "volume".
		metrics := []string{"open", "high", "low", "close", "volume"}
		for _, metric := range metrics {
			// Constructing each line as "metric: value".
			// The keys in the inner map might need adjustments to match the actual JSON structure.
			formattedData = append(formattedData, fmt.Sprintf("%s: %s", metric, data[date][fmt.Sprintf("1. %s", metric)]))
		}

		// Adding a separator after each date's data for readability.
		formattedData = append(formattedData, "----------------")
	}

	return formattedData
}
