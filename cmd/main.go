package main

import (
	"fmt"
	"mangomarkets/internal"
	"mangomarkets/internal/api"
)

func main() {
	apiKey, _ := internal.LoadEnv()

	intradayStockInfo, err := api.IntradayDataPull(apiKey, "AAPL", "5min")
	if err != nil {
		fmt.Println("broken", err)
	}

	dailyStockInfo, err := api.DailyDataPull(apiKey, "AAPL")
	if err != nil {
		fmt.Println("broken", err)
	}

	weeklyStockInfo, err := api.WeeklyDataPull(apiKey, "NVDA")
	if err != nil {
		fmt.Println("broken", err)
	}

	monthlyStockInfo, err := api.MonthlyDataPull(apiKey, "NVDA")
	if err != nil {
		fmt.Println("broken", err)
	}

	fmt.Println(intradayStockInfo.TimeSeries)
	fmt.Print("\n\n\n\n\n\n\n")
	fmt.Println(dailyStockInfo.TimeSeries)
	fmt.Print("\n\n\n\n\n\n\n")
	fmt.Println(weeklyStockInfo.TimeSeries)
	fmt.Print("\n\n\n\n\n\n\n")
	fmt.Println(monthlyStockInfo.TimeSeries)
	fmt.Print("\n\n\n\n\n\n\n")

	news, err := api.FetchNewsSentimentData(apiKey, "NVDA")
	fmt.Println(news.Feed)
}
