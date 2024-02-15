// main.go
package main

import (
	"fmt"

	"github.com/jp-mango/mangomarkets/internal/api"
	"github.com/jp-mango/mangomarkets/internal/config"
	"github.com/jp-mango/mangomarkets/util"
)

func main() {
	apiKey := config.LoadEnv()

	var ticker string
	fmt.Print("Enter ticker: ")
	fmt.Scanln(&ticker)

	var interval int
	fmt.Println("Time interval?\n [1]:daily  [2]:weekly  [3]monthly")
	fmt.Scanln(&interval)

	switch interval {
	case 1:
		tsDataDaily := api.FetchTimeSeriesDaily(apiKey, ticker)
		util.PrintTimeSeriesData(tsDataDaily)
	case 2:
		tsDataWeekly := api.FetchTimeSeriesWeekly(apiKey, ticker)
		util.PrintTimeSeriesData(tsDataWeekly)
	case 3:
		tsDataMonthly := api.FetchTimeSeriesMonthly(apiKey, ticker)
		util.PrintTimeSeriesData(tsDataMonthly)
	default:
		fmt.Println("Invalid interval")
	}
}
