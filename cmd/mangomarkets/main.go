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

	tsDataDaily := api.FetchTimeSeriesDaily(apiKey, ticker)
	util.PrintTimeSeriesData(tsDataDaily)

	tsDataWeekly := api.FetchTimeSeriesWeekly(apiKey, ticker)
	util.PrintTimeSeriesData(tsDataWeekly)

	tsDataMonthly := api.FetchTimeSeriesMonthly(apiKey, ticker)
	util.PrintTimeSeriesData(tsDataMonthly)
}
