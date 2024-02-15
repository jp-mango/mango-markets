// main.go
package main

import (
	"fmt"

	"github.com/jp-mango/mangomarkets/internal/api"
	"github.com/jp-mango/mangomarkets/internal/config"
	"github.com/jp-mango/mangomarkets/util"
)

func main() {
	// charm.Start()
	apiKey := config.LoadEnv()

	var ticker string
	fmt.Print("Enter ticker: ")
	fmt.Scanln(&ticker)

	var interval int
	fmt.Println("Time interval?\n [1]:daily  [2]:weekly  [3]monthly")
	fmt.Scanln(&interval)

	switch interval {
	case 1:
		tsDataDaily, err := api.FetchTimeSeriesDaily(apiKey, ticker)
		if err != nil {
			fmt.Print("Unable to load daily time series data for", ticker)
		}
		util.PrintTimeSeriesData(tsDataDaily)
	case 2:
		tsDataWeekly, err := api.FetchTimeSeriesWeekly(apiKey, ticker)
		if err != nil {
			fmt.Print("Unable to load daily time series data for", ticker)
		}
		util.PrintTimeSeriesData(tsDataWeekly)
	case 3:
		tsDataMonthly, err := api.FetchTimeSeriesMonthly(apiKey, ticker)
		if err != nil {
			fmt.Print("Unable to load daily time series data for", ticker)
		}
		util.PrintTimeSeriesData(tsDataMonthly)
	default:
		fmt.Println("Invalid interval")
	}
}
