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
	reader := bufio.NewReader(os.Stdin)
	apiKey := config.LoadEnv()

mainLoop:
	for {
		fmt.Println("\nWelcome to Mango Markets!🥭")
		fmt.Println("1. Stock Market")
		fmt.Println("2. Financial News")
		fmt.Println("3. Forex Market")
		fmt.Println("4. Cryptocurrency Market")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		choiceMain, _ := reader.ReadString('\n')
		choiceMain = strings.TrimSpace(choiceMain)

		switch choiceMain {
		case "1":
		stockMarket:
			for {
				fmt.Println("\nStock Market Data:🏦")
				fmt.Println("1. Top Gainers and Losers")
				fmt.Println("2. Ticker Search")
				fmt.Println("3. Global Market Status")
				fmt.Println("4. Return to Main Menu")
				fmt.Print("Enter a choice: ")
				choiceStock, _ := reader.ReadString('\n')
				choiceStock = strings.TrimSpace(choiceStock)

				switch choiceStock {
				case "1":
					// TODO: Implementation for top gainers and losers
					fmt.Println("Retrieving top gainers and losers...")
					gainLoss, err := api.FetchGainLoss(apiKey)
					if err != nil {
						fmt.Println("Unable to load the top winners and losers")
					} else {
						util.PrintTopGainersAndLosers(gainLoss)
					}

				case "2":
				tickerSearch:
					for {
						var ticker string
						fmt.Print("\nEnter ticker: ")
						ticker, _ = reader.ReadString('\n')
						ticker = strings.TrimSpace(ticker)

						fmt.Printf("\nWhat data would you like to see for %s?\n", strings.ToUpper(ticker))
						fmt.Println("1. Stock Price")
						fmt.Println("2. Company Overview")
						fmt.Println("3. Income Statement")
						fmt.Println("4. Balance Sheet")
						fmt.Println("5. Cash Flow")
						fmt.Println("6. Earnings")
						fmt.Println("7. Return To Stock Market Data")
						fmt.Print("Enter a choice: ")
						tickerChoice, _ := reader.ReadString('\n')
						tickerChoice = strings.TrimSpace(tickerChoice)

						switch tickerChoice {
						case "1": // stock price
							fmt.Println("Time interval - [1]:daily [2]:weekly [3]:monthly")
							var interval string
							interval, _ = reader.ReadString('\n')
							interval = strings.TrimSpace(interval)

							switch interval {
							case "1": // daily prices
								fmt.Printf("\nDaily prices for %v:\n\n", strings.ToUpper(ticker))
								tsDataDaily, err := api.FetchTimeSeriesDaily(apiKey, ticker)
								if err != nil {
									fmt.Println("Unable to load daily time series data for", ticker)
								} else {
									util.PrintTimeSeriesData(tsDataDaily)
								}
							case "2": // weekly prices
								fmt.Printf("\nWeekly prices for %v:\n\n", strings.ToUpper(ticker))
								tsDataWeekly, err := api.FetchTimeSeriesWeekly(apiKey, ticker)
								if err != nil {
									fmt.Println("Unable to load weekly time series data for", ticker)
								} else {
									util.PrintTimeSeriesData(tsDataWeekly)
								}
							case "3": // monthly prices
								fmt.Printf("\nMonthly prices for %v:\n\n", strings.ToUpper(ticker))
								tsDataMonthly, err := api.FetchTimeSeriesMonthly(apiKey, ticker)
								if err != nil {
									fmt.Println("Unable to load monthly time series data for", ticker)
								} else {
									util.PrintTimeSeriesData(tsDataMonthly)
								}
							default:
								fmt.Println("Invalid interval")
							}

						case "2": // company overview
							companyInfo, err := api.FetchCompanyOverview(ticker, apiKey)
							if err != nil {
								fmt.Printf("Unable to fetch company info for %s", ticker)
							} else {
								util.PrintCompanyInfo(companyInfo)
							}
						case "3":
							fmt.Printf("[A]nnual or [Q]uarterly Income Statement for %s?\n", strings.ToUpper(ticker))
							incomeTimeFrame, _ := reader.ReadString('\n')
							incomeTimeFrame = strings.ToUpper(strings.TrimSpace(incomeTimeFrame))
							switch incomeTimeFrame {
							case "A":
								incomeStatement, err := api.FetchIncomeStatement(ticker, apiKey)
								if err != nil {
									fmt.Printf("Unable to fetch annual income statement for %s", ticker)
								} else {
									util.PrintAnnualIncomeStatement(incomeStatement)
									fmt.Println()
								}
							case "Q":
								incomeStatement, err := api.FetchIncomeStatement(ticker, apiKey)
								if err != nil {
									fmt.Printf("Unable to fetch annual income statement for %s", ticker)
								} else {
									util.PrintQuarterlyIncomeStatement(incomeStatement)
									fmt.Println()
								}
							}

						case "4":
							// TODO: Balance Sheet
							fmt.Println("Balance sheet functionality to be implemented.")

						case "5":
							// TODO: Cash Flow
							fmt.Println("Cash flow functionality to be implemented.")

						case "6":
							// TODO: Earnings
							fmt.Println("Earnings functionality to be implemented.")

						case "7": // return to stock market data
							continue stockMarket
						}

						if tickerChoice != "7" {
							fmt.Println("Press any key to return to Ticker Search")
							_, _ = reader.ReadString('\n')
							continue tickerSearch
						}
					}
				case "3":
					// TODO: Global Market Status
					fmt.Println("Global market status:")
					marketHours, err := api.FetchMarketStatus(apiKey)
					if err != nil {
						fmt.Println("Unable to fetch market hours", err)
					} else {
						util.PrintMarketStatus(marketHours)
					}
				case "4": // return to main loop
					continue mainLoop
				}
			}
		case "2":
			// TODO: Financial News
			fmt.Println("Financial news functionality to be implemented.")

		case "3":
			// TODO: Forex Market
			fmt.Println("Forex market functionality to be implemented.")

		case "4":
			// TODO: Crypto Market
			fmt.Println("crypto market functionality to be implemented.")

		case "5":
			fmt.Println("\nMay your profits be HIGH and your risks LOW! 🥭")
			break mainLoop
		}
	}
}
