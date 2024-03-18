package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jp-mango/mangomarkets/internal/api"
	"github.com/jp-mango/mangomarkets/internal/config"
	"github.com/jp-mango/mangomarkets/internal/db"
	"github.com/jp-mango/mangomarkets/util"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	apiKey := config.LoadEnv()

	// Connect to MongoDB
	client, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Select the database
	database := client.Database("mangomarkets")

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
		case "1": // Stock Market
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
					fmt.Println("Retrieving top gainers and losers...")
					gainLoss, err := api.FetchGainLoss(apiKey)
					if err != nil {
						fmt.Println("Unable to load the top winners and losers")
					} else {
						util.PrintTopGainersAndLosers(gainLoss)
					}

				case "2":
					var ticker string
				tickerEntry:
					for {
						fmt.Print("\nEnter ticker: ")
						ticker, _ = reader.ReadString('\n')
						ticker = strings.ToUpper(strings.TrimSpace(ticker))
						status, err := util.CheckTickerStatus(ticker)
						if err != nil {
							fmt.Print("error validating entry", err)
						} else if !status {
							continue tickerEntry
						} else {
							break
						}
					}
				tickerSearch:
					for {
						fmt.Printf("\nWhat data would you like to see for %s?\n", ticker)
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
								//TODO: check if most recent data is in table (current day, week, month), then pull from api
								results, err := api.FetchSavedData(client, "daily_stock_price_data", ticker)
								if err != nil || len(results) == 0 {
									dailyData, err := api.SaveStockDataDaily(apiKey, ticker, database.Collection("daily_stock_price_data"))
									if err != nil {
										fmt.Printf("Unable to save daily time series data for %s: %v\n", ticker, dailyData)
									} else {
										fmt.Println("Daily prices saved to the database.")
									}
								} else if len(results) != 0 {
									fmt.Println(results)
								}

							case "2": // weekly prices
								//TODO: check if most recent data is in table, then pull
								results, err := api.FetchSavedData(client, "weekly_stock_price_data", ticker)
								if err != nil || len(results) == 0 {
									weeklyData, err := api.SaveStockDataWeekly(apiKey, ticker, database.Collection("weekly_stock_price_data"))
									if err != nil {
										fmt.Printf("Unable to save weekly time series data for %s: %v\n", ticker, weeklyData)
									} else {
										fmt.Println("Weekly prices saved to the database.")
									}
								} else if len(results) != 0 {
									fmt.Println(results)
								}
							case "3": // monthly prices
								//TODO: check if most recent data is in table, then pull
								results, err := api.FetchSavedData(client, "monthly_stock_price_data", ticker)
								if err != nil || len(results) == 0 {
									monthlyData, err := api.SaveStockDataMonthly(apiKey, ticker, database.Collection("monthly_stock_price_data"))
									if err != nil {
										fmt.Printf("Unable to save monthly time series data for %s: %v\n", ticker, monthlyData)
									} else {
										fmt.Println("Monthly prices saved to the database.")
									}
								} else if len(results) != 0 {
									fmt.Println(results)
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

						case "3": //income statements
							fmt.Printf("[A]nnual or [Q]uarterly Income Statement for %s?\n", strings.ToUpper(ticker))
							incomeTimeFrame, _ := reader.ReadString('\n')
							incomeTimeFrame = strings.ToUpper(strings.TrimSpace(incomeTimeFrame))
							switch incomeTimeFrame {
							case "A":
								incomeStatement, err := api.FetchIncomeStatement(ticker, apiKey)
								if err != nil {
									fmt.Printf("Unable to fetch annual income statement for %s\n", ticker)
								} else {
									util.PrintAnnualIncomeStatement(incomeStatement)
									fmt.Println()
								}
							case "Q":
								incomeStatement, err := api.FetchIncomeStatement(ticker, apiKey)
								if err != nil {
									fmt.Printf("Unable to fetch quarterly income statement for %s\n", ticker)
								} else {
									util.PrintQuarterlyIncomeStatement(incomeStatement)
									fmt.Println()
								}
							}

						case "4": // balance sheet
							fmt.Printf("[A]nnual or [Q]uarterly Balance Sheet for %s?\n", strings.ToUpper(ticker))
							balanceSheet, _ := reader.ReadString('\n')
							balanceSheet = strings.ToUpper(strings.TrimSpace(balanceSheet))
							switch balanceSheet {
							case "A":
								balanceSheet, err := api.FetchBalanceSheet(ticker, apiKey)
								if err != nil {
									fmt.Printf("Unable to fetch annual income statement for %s\n", ticker)
								} else {
									util.PrintAnnualBalanceSheet(balanceSheet)
									fmt.Println()
								}
							case "Q":
								balanceSheet, err := api.FetchBalanceSheet(ticker, apiKey)
								if err != nil {
									fmt.Printf("Unable to fetch quarterly income statement for %s\n", ticker)
								} else {
									util.PrintQuarterlyBalanceSheet(balanceSheet)
									fmt.Println()
								}
							}

						case "5": //cashflow
							fmt.Printf("[A]nnual or [Q]uarterly Cashflow for %s?\n", strings.ToUpper(ticker))
							cashFlow, _ := reader.ReadString('\n')
							cashFlow = strings.ToUpper(strings.TrimSpace(cashFlow))
							switch cashFlow {
							case "A":
								cashFlow, err := api.FetchCashflow(ticker, apiKey)
								if err != nil {
									fmt.Printf("Unable to retrieve cashflow data for %s\n", ticker)
								} else {
									util.PrintAnnualCashflow(cashFlow)
								}
							case "Q":
								cashFlow, err := api.FetchCashflow(ticker, apiKey)
								if err != nil {
									fmt.Printf("Unable to retrieve cashflow data for %s\n", ticker)
								} else {
									util.PrintQuarterlyCashflow(cashFlow)
								}
							}

						case "6":
							fmt.Printf("[A]nnual or [Q]uarterly Earnings Report for %s?\n", strings.ToUpper(ticker))
							earnings, _ := reader.ReadString('\n')
							earnings = strings.ToUpper(strings.TrimSpace(earnings))
							switch earnings {
							case "A":
								earnings, err := api.FetchEarnings(ticker, apiKey)
								if err != nil {
									fmt.Printf("Unable to retrieve cashflow data for %s\n", ticker)
								} else {
									util.PrintAnnualEarnings(earnings)
								}
							case "Q":
								earnings, err := api.FetchEarnings(ticker, apiKey)
								if err != nil {
									fmt.Printf("Unable to retrieve cashflow data for %s\n", ticker)
								} else {
									util.PrintQuarterlyEarnings(earnings)
								}
							}

						case "7": // return to stock market data
							continue stockMarket

						default:
							fmt.Println("Invalid choice. Please enter a valid number or press '7' to return.")
							continue tickerSearch
						}

						if tickerChoice != "7" {
							fmt.Printf("\nContinue looking at %s data? [y]/[n]\n", strings.ToUpper(ticker))
							r, _ := reader.ReadString('\n')
							r = strings.ToLower(strings.TrimSpace(r))
							if r == "y" {
								continue tickerSearch
							} else if r == "n" {
								break tickerSearch // This will exit the tickerSearch loop and proceed to the next iteration of the enclosing loop
							} else {
								fmt.Println("Invalid response. Returning to ticker search.")
								continue tickerSearch
							}
						}
					}

				case "3": // global market status
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
		case "2": // News
		newsSearch:
			for {
				fmt.Printf("\nFinancial News:📰\n")
				fmt.Println("1. Search By Ticker")
				fmt.Println("2. Search By Topic")
				fmt.Println("3. Return To Main Menu")
				fmt.Print("Enter your choice: ")

				newsChoice, _ := reader.ReadString('\n')
				newsChoice = strings.TrimSpace(newsChoice)
				switch newsChoice {
				case "1":
					var ticker string
					fmt.Print("\nEnter ticker: ")
					ticker, _ = reader.ReadString('\n')
					ticker = strings.TrimSpace(ticker)

					news, err := api.FetchNewsByTicker(ticker, apiKey)
					if err != nil {
						fmt.Println("Unable to fetch news for", strings.ToUpper(ticker))
					} else {
						util.PrintNews(news)
					}
					continue newsSearch
				case "2":
					fmt.Printf("\nAvailable Topics:\n\n")
					fmt.Println("0. Return To Previous Menu")
					fmt.Println("1. Blockchain")
					fmt.Println("2. Earnings")
					fmt.Println("3. IPO")
					fmt.Println("4. Mergers & Acquisitions")
					fmt.Println("5. Financial Markets")
					fmt.Println("6. Economy - Fiscal Policy (e.g., tax reform, government spending)")
					fmt.Println("7. Economy - Monetary Policy (e.g., interest rates, inflation)")
					fmt.Println("8. Economy - Macro/Overall")
					fmt.Println("9. Energy & Transportation")
					fmt.Println("10. Finance")
					fmt.Println("11. Life Science")
					fmt.Println("12. Manufacturing")
					fmt.Println("13. Real Estate & Construction")
					fmt.Println("14. Retail & Wholesale")
					fmt.Println("15. Technology")
					fmt.Print("\nEnter choice: ")

					topics, _ := reader.ReadString('\n')
					topics = strings.TrimSpace(topics)

					if topics == "0" {
						continue newsSearch
					} else {
						userTopics := strings.Split(topics, " ")
						news, err := api.FetchNewsByTopic(util.ConstructTopicsURL(apiKey, userTopics), apiKey)
						if err != nil {
							fmt.Println("Unable to fetch news for those topics")
						} else {
							util.PrintNews(news)
						}
						continue newsSearch
					}

				case "3":
					continue mainLoop
				}
			}

		case "3": // Forex
		forexMenu:
			for {
				fmt.Println("\n Forex Data:💱")
				fmt.Println("1. Search Currency Pair")
				fmt.Println("2. Return To Main Menu")
				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(choice)

			forexSearch:
				for {
					switch choice {

					case "1":
						fmt.Print("\nEnter base currency: ")
						base, _ := reader.ReadString('\n')
						base = strings.ToUpper(strings.TrimSpace(base))

						fmt.Print("Enter the quote currency: ")
						quote, _ := reader.ReadString('\n')
						quote = strings.ToUpper(strings.TrimSpace(quote))

						if util.ForexVerification(base) && util.ForexVerification(quote) && (base != quote) {
							fmt.Printf("\n1. (%s/%s) Exchange Rate\n", base, quote)
							fmt.Printf("2. (%s/%s) Time Series Data\n", base, quote)
							fmt.Println("3. Return To Pair Search")
							choice, _ := reader.ReadString('\n')
							choice = strings.TrimSpace(choice)
							switch choice {
							case "1":
								fmt.Printf("Exchange rate between %s and %s", base, quote)
								exchangeRate, err := api.FetchExchangeRate(apiKey, base, quote)
								if err != nil {
									fmt.Printf("Unable to  retrieve exchange rate for (%s/%s): %s", base, quote, err)
								}
								fmt.Println(exchangeRate.ExchangeRate)
							case "2":
								fmt.Print("Time interval?\n[1] = daily, [2] = weekly, [3] = monthly: ")
								choice, _ := reader.ReadString('\n')
								choice = strings.TrimSpace(choice)
								switch choice {
								case "1":
									fmt.Printf("\nDaily prices for (%s/%s):\n", base, quote)
									tsDataDaily, err := api.FetchForexTimeSeriesDaily(apiKey, base, quote)
									if err != nil {
										fmt.Printf("Unable to load daily time series data for (%s/%s): %s\n", quote, base, err)
									} else {
										util.PrintTimeSeriesData(tsDataDaily, tsDataDaily.MetaData)
									}
								case "2":
									fmt.Printf("\nWeekly prices for (%s/%s):\n", base, quote)
									tsDataWeekly, err := api.FetchForexTimeSeriesWeekly(apiKey, base, quote)
									if err != nil {
										fmt.Printf("Unable to load weekly time series data for (%s/%s)\n", quote, base)
									} else {
										util.PrintTimeSeriesData(tsDataWeekly, tsDataWeekly.MetaData)
									}
								case "3":
									fmt.Printf("\nMonthly prices for (%s/%s):\n", base, quote)
									tsDataMonthly, err := api.FetchForexTimeSeriesMonthly(apiKey, base, quote)
									if err != nil {
										fmt.Printf("Unable to load monthly time series data for (%s/%s)\n", quote, base)
									} else {
										util.PrintTimeSeriesData(tsDataMonthly, tsDataMonthly.MetaData)
									}
								}

							case "3":
								continue forexSearch
							}
						} else if util.ForexVerification(base) && !util.ForexVerification(quote) {
							fmt.Println("Quote Not Found")
						} else if !util.ForexVerification(base) && util.ForexVerification(quote) {
							fmt.Println("Base Not Found")
						} else if base == quote {
							fmt.Println("Please enter different values for the base and quote currencies")
						} else {
							continue forexMenu
						}

					case "2":
						continue mainLoop
					}
				}
			}

		case "4": // Crypto
			// TODO: Crypto Market
			fmt.Println("crypto market functionality to be implemented.")

		case "5": // Exit
			fmt.Println("\nMay your profits be HIGH and your risks LOW! 🥭")
			break mainLoop
		}
	}
}
