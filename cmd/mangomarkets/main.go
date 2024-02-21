// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jp-mango/mangomarkets/internal/api"
	"github.com/jp-mango/mangomarkets/internal/config"
	"github.com/jp-mango/mangomarkets/util"
)

func main() {
	//	charm.Start()
	//}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nWelcome to Mango Markets!🥭\n\n")
		fmt.Println("1. Stock Market")
		fmt.Println("2. Financial News")
		fmt.Println("3. Forex Market")
		fmt.Println("4. Cryptocurrency Market")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			apiKey := config.LoadEnv()
			var ticker string
			fmt.Print("\nEnter ticker: ")
			fmt.Scanln(&ticker)

			var interval int
			fmt.Println("Time interval?\n [1]:daily  [2]:weekly  [3]monthly")
			fmt.Scanln(&interval)

			switch interval {
			case 1:
				fmt.Printf("\nDaily prices for %v:\n\n", strings.ToUpper(ticker))
				tsDataDaily, err := api.FetchTimeSeriesDaily(apiKey, ticker)
				if err != nil {
					fmt.Print("Unable to load daily time series data for", ticker)
				}
				util.PrintTimeSeriesData(tsDataDaily)
			case 2:
				fmt.Printf("\nWeekly prices for %v:\n\n", strings.ToUpper(ticker))
				tsDataWeekly, err := api.FetchTimeSeriesWeekly(apiKey, ticker)
				if err != nil {
					fmt.Print("Unable to load daily time series data for", ticker)
				}
				util.PrintTimeSeriesData(tsDataWeekly)
			case 3:
				fmt.Printf("\nMonthly prices for %v:\n\n", strings.ToUpper(ticker))
				tsDataMonthly, err := api.FetchTimeSeriesMonthly(apiKey, ticker)
				if err != nil {
					fmt.Print("Unable to load daily time series data for", ticker)
				}
				util.PrintTimeSeriesData(tsDataMonthly)
			default:
				fmt.Println("Invalid interval")
			}
		case "5":
			fmt.Println("\nmay your profits be HIGH and your risk LOW!")
			return
		}
	}
}
