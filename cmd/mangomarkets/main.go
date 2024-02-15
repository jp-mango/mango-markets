// main.go
package main

import (
	"fmt"
	"sort"

	"github.com/jp-mango/mangomarkets/internal/api"
	"github.com/jp-mango/mangomarkets/internal/config"
)

func main() {
	apiKey := config.LoadEnv()

	var ticker string
	fmt.Print("Enter ticker: ")
	fmt.Scanln(&ticker)

	tsData := api.FetchTimeSeriesDaily(apiKey, ticker)

	// Sort and print as before
	var dates []string
	for date := range tsData.TimeSeries {
		dates = append(dates, date)
	}
	sort.Strings(dates) // Adjust sorting as needed

	for _, date := range dates {
		fmt.Printf("Date: %s\n", date)
		for key, value := range tsData.TimeSeries[date] {
			fmt.Printf("%s: %s\n", key, value)
		}
		fmt.Println("----------------")
	}
}
