package main

func main() {
	//apiKey, _ := internal.LoadEnv()
	/*
			intradayStockInfo, err := api.IntradayDataPull(apiKey, "AAPL", "5min")
			if err != nil {
				fmt.Println("intraday broken", err)
			}

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

			fmt.Println(intradayStockInfo.TimeSeries)
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
				fmt.Printf("%s (%s): %s| ",active, active.Ticker, active.Volume)
			}
	*/
}
