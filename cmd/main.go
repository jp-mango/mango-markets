package main

import (
	"fmt"
	"os"

	"mangomarkets/internal/api"
	"mangomarkets/internal/load"

	_ "github.com/lib/pq"
)

func main() {
	//! Set up logging
	logger, logFile, err := load.Logging()
	if err != nil {
		fmt.Println(fmt.Errorf("error loading logging: %v", err))
		os.Exit(1)
	}
	defer logFile.Close()

	//! Load env variables
	apiKey, DB_CONN, ACTIVE_STOCKS, err := load.Env()
	if err != nil {
		logger.Error(err.Error())
	}

	//! Connect to DB
	db, err := load.DB(DB_CONN)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	//! bizniz logic

	intradayData, err := api.FetchIntradayTSData(apiKey, "AMD", "30min")
	if err != nil || intradayData == nil {
		logger.Error(err.Error())
	}

	dailyData, err := api.FetchDailyTSData(apiKey, "SMCI")
	if err != nil || dailyData == nil {
		logger.Error(err.Error())
	}

	weeklyData, err := api.FetchWeeklyTSData(apiKey, "AMD")
	if err != nil || weeklyData == nil {
		logger.Error(err.Error())
	}

	monthlyData, err := api.FetchMonthlyTSData(apiKey, "AMD")
	if err != nil || monthlyData == nil {
		logger.Error(err.Error())
	}

	companyInfo, err := api.FetchCompanyInfo("AMD", apiKey)
	if err != nil || companyInfo == nil {
		logger.Error(err.Error())
	}

	incomeStatement, err := api.FetchIncomeStatement("AMD", apiKey)
	if err != nil || incomeStatement == nil {
		logger.Error(err.Error())
	}

	balanceSheet, err := api.FetchBalanceSheet("NVDA", apiKey)
	if err != nil || balanceSheet == nil {
		logger.Error(err.Error())
	}

	cashFlow, err := api.FetchCashFlow("AAPL", apiKey)
	if err != nil || cashFlow == nil {
		logger.Error(err.Error())
	}

	earnings, err := api.FetchEarnings("AAPL", apiKey)
	if err != nil || earnings == nil {
		logger.Error(err.Error())
	}

	listings, err := api.FetchActiveListings(apiKey)
	if err != nil {
		logger.Error(err.Error())
	} else {
		path := ACTIVE_STOCKS
		err = os.WriteFile(path, []byte(listings), 0644)
		if err != nil {
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("wrote listings to %s", path))
	}

	var ticker string
	fmt.Print("Enter a ticker: ")
	fmt.Scan(&ticker)
	ticker = api.SanitizeTicker(ticker)

	if api.FoundTickerInput(ticker) {
		fmt.Printf("%s was found\n", ticker)
	}

}
