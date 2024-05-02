package main

import (
	"fmt"
	"mangomarkets/internal"
	"mangomarkets/internal/api"
)

func main() {
	apiKey, _ := internal.LoadEnv()
	/*
							intradayStockInfo, err := api.IntradayDataPull(apiKey, "AAPL", "5min")
							if err != nil {
								fmt.Println("intraday broken", err)
							}
							fmt.Println(intradayStockInfo.TimeSeries)

							dailyStockInfo, err := api.DailyDataPull(apiKey, "AAPL")
							if err != nil {
								fmt.Println("daily broken", err)
							}

							weeklyStockInfo, err := api.WeeklyDataPull(apiKey, "NVDA")
							if err != nil {
								fmt.Println("weekly broken", err)
							}

							monthlyStockInfo, err := api.MonthlyDataPull(apiKey, "NVDA")
							if err != nil {
								fmt.Println("monthly broken", err)
							}

							fmt.Print("\n\n\n\n\n\n\n")
							fmt.Println(dailyStockInfo.TimeSeries)
							fmt.Print("\n\n\n\n\n\n\n")
							fmt.Println(weeklyStockInfo.TimeSeries)
							fmt.Print("\n\n\n\n\n\n\n")
							fmt.Println(monthlyStockInfo.TimeSeries)
							fmt.Print("\n\n\n\n\n\n\n")

							news, err := api.FetchNewsSentimentData(apiKey, "NVDA")
							if err != nil {
								fmt.Println("news broken", err)
							}

							for _, inf := range news.Feed {
								fmt.Printf("|%s|\n- %s (%s)\n", inf.Title, inf.Summary, inf.URL)
							}

							gainLoss, err := api.FetchTopGainLossData(apiKey)

							if err != nil {
								fmt.Println("gain loss broken", err)
							}

							for _, active := range gainLoss.MostActive {
								fmt.Printf("%s (%s): %s| ", active, active.Ticker, active.Volume)
							}


						companyInfo, _ := api.FetchCompanyInfo("AAPL", apiKey)

						fmt.Println(companyInfo)

					incomeStatement, _ := api.FetchIncomeStatement("AAPL", apiKey)

					for _, x := range incomeStatement.AnnualReports {
						fmt.Println(x.CostOfRevenue)
					}

				income,_ := api.FetchIncomeStatement("AAPL",apiKey)

				for _, a := range income.QuarterlyIncomeStatement{
					fmt.Printf("%s: %s\n",a.FiscalDateEnding,a.GrossProfit)
				}

			bal, _ := api.FetchBalanceSheet("AAPL", apiKey)
			for _, x := range bal.AnnualBalanceSheet {
				fmt.Printf("%s - %s: Current Debt: %s\n", bal.Symbol, x.FiscalDateEnding, x.CurrentDebt)
			}

		data, _ := api.FetchCashFlow("AMD", apiKey)

		for _, x := range data.QuarterlyCashFlow {
			fmt.Printf("%s (%s) = %s\n", x.FiscalDateEnding, data.Symbol, x.NetIncome)
		}
	*/

	data, _ := api.FetchEarnings("AMD", apiKey)

	for _, e := range data.QuarterlyEarnings {
		fmt.Printf("%s (%s): Estimated = %s | Reported = %s | Surprise = %s\n", e.ReportedDate, data.Symbol, e.EstimatedEPS, e.ReportedEPS, e.SurprisePercentage)
	}
}
