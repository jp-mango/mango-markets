package util

import (
	"fmt"
	"sort"

	"github.com/jp-mango/mangomarkets/internal/api"
)

func PrintTimeSeriesData(provider api.TimeSeriesData) {
	tsData := provider.GetTimeSeriesData()

	var dates []string
	for date := range tsData {
		dates = append(dates, date)
	}
	// Sorts in ascending order
	sort.Slice(dates, func(i, j int) bool {
		return dates[i] < dates[j]
	})
	for _, date := range dates {
		fmt.Printf("Date: %s\n", date)
		// Collect keys of the inner map
		var keys []string
		for key := range tsData[date] {
			keys = append(keys, key)
		}
		// Sort the keys to ensure consistent order
		sort.Strings(keys)
		// Now iterate over the sorted keys
		for _, key := range keys {
			value := tsData[date][key]
			fmt.Printf("%s: %s\n", key, value)
		}
		fmt.Println("----------------")
	}
}

func PrintTopGainersAndLosers(marketData *api.TopGainLoss) {
	fmt.Println("Top Gainers:")

	for _, gainer := range marketData.TopGainers {
		fmt.Printf("Ticker: %s, Price: %s, Change: %s%%, Volume: %s\n", gainer.Ticker, gainer.Price, gainer.ChangePercentage, gainer.Volume)
	}
	fmt.Println("\nTop Losers:")

	for _, loser := range marketData.TopLosers {
		fmt.Printf("Ticker: %s, Price: %s, Change: %s%%, Volume: %s\n", loser.Ticker, loser.Price, loser.ChangePercentage, loser.Volume)
	}
}

func PrintMarketStatus(marketStatus *api.MarketHours) {
	fmt.Println(marketStatus.Endpoint)
	for _, market := range marketStatus.Markets {
		fmt.Printf("\nMarket Type: %s\nRegion: %s\nPrimary Exchanges: %s\nOpen: %s\nClose: %s\nStatus: %s\nNotes: %s\n",
			market.MarketType, market.Region, market.PrimaryExchanges, market.LocalOpen, market.LocalClose, market.CurrentStatus, market.Notes)
	}
}

func PrintCompanyInfo(companyInfo *api.CompanyOverview) {
	fmt.Printf("Overview for %s(%s)\n\n", companyInfo.Name, companyInfo.Symbol)
	fmt.Printf("Description: %s\n", companyInfo.Description)
	fmt.Printf("Central Index Key (CIK): %s\n", companyInfo.CIK)
	fmt.Printf("Exchange: %s\n", companyInfo.Exchange)
	fmt.Printf("Currency: %s\n", companyInfo.Currency)
	fmt.Printf("Country: %s\n", companyInfo.Country)
	fmt.Printf("Sector: %s\n", companyInfo.Sector)
	fmt.Printf("Industry %s\n", companyInfo.Industry)
	fmt.Printf("Address: %s\n", companyInfo.Address)
	fmt.Printf("Fiscal Year End: %s\n", companyInfo.FiscalYearEnd)
	fmt.Printf("Latest Quarter: %s\n", companyInfo.LatestQuarter)
	fmt.Printf("Market Cap: %s\n", companyInfo.MarketCap)
	fmt.Printf("EBITDA: %s\n", companyInfo.EBITDA)      // Earnings Before Interest, Taxes, Depreciation, and Amortization. A measure of a company's overall financial performance.
	fmt.Printf("PE Ratio: %s\n", companyInfo.PERatio)   // (Price-to-Earnings Ratio): A valuation ratio of a company's current share price compared to its per-share earnings.
	fmt.Printf("PEG Ratio: %s\n", companyInfo.PEGRatio) // (Price/Earnings to Growth Ratio): A stock's price-to-earnings ratio divided by the growth rate of its earnings for a specified time period.
	fmt.Printf("Book Value: %s\n", companyInfo.BookVal) // The net asset value of a company calculated as total assets minus intangible assets (patents, goodwill) and liabilities.
	fmt.Printf("Dividend Per Share: %s\n", companyInfo.Dividend)
	fmt.Printf("Dividend Yield: %s\n", companyInfo.DividendYield) // A financial ratio that shows how much a company pays out in dividends each year relative to its stock price.
	fmt.Printf("Earnings Per Share (TTM): %s\n", companyInfo.EPS)
	fmt.Printf("Revenue Per Share (TTM): %s\n", companyInfo.ShareRevenue)
	fmt.Printf("Profit Margin: %s\n", companyInfo.ProfitMargin)
	fmt.Printf("Operating Margin (TTM): %s\n", companyInfo.OperatingMargin)
	fmt.Printf("Return on Assets (TTM): %s\n", companyInfo.ROA)
	fmt.Printf("Return on Equity (TTM): %s\n", companyInfo.ROE)
	fmt.Printf("Revenue (TTM): %s\n", companyInfo.Revenue)
	fmt.Printf("Gross Profit (TTM): %s\n", companyInfo.GrossProfit)
	fmt.Printf("Diluted Earnings Per Share: %s\n", companyInfo.DilutedEPSTTM)
	fmt.Printf("Quarterly Earnings Growth YOY: %s\n", companyInfo.QuarterlyEarningsGrowthYOY)
	fmt.Printf("Quarterly Revenue Growth YOY: %s\n", companyInfo.QuarterlyRevenueGrowthYOY)
	fmt.Printf("Analyst Target Price: %s\n", companyInfo.AnalystTargetPrice)
	fmt.Printf("Trailing Price to Earnings: %s\n", companyInfo.TrailingPE)
	fmt.Printf("Forward Price to Earnings: %s\n", companyInfo.ForwardPE)
	fmt.Printf("Price to Sales Ratio: %s\n", companyInfo.PriceToSales)
	fmt.Printf("Price to Book Ratio: %s\n", companyInfo.PriceToBook)
	fmt.Printf("Enterprise Value to Revenue: %s\n", companyInfo.EVToRevenue)
	fmt.Printf("Enterprise Value to EBITDA:  %s\n", companyInfo.EVToEBITDA)
	fmt.Printf("Beta: %s\n", companyInfo.Beta)
	fmt.Printf("52-Week-High: %s\n", companyInfo.YearHigh)
	fmt.Printf("52-Week-Low: %s\n", companyInfo.YearLow)
	fmt.Printf("50 Day Moving Average: %s\n", companyInfo.FiddyDayMA)
	fmt.Printf("200 Day Moving Average: %s\n", companyInfo.TwoHunnaDayMA)
	fmt.Printf("Shares Outstanding: %s\n", companyInfo.SharesOutstanding)
	fmt.Printf("Dividend Date: %s\n", companyInfo.DividendDate)
	fmt.Printf("ExDividend Date: %s\n", companyInfo.ExDividendDate)
}

func PrintAnnualIncomeStatement(incomeStatement *api.IncomeStatement) {
	for _, incomeStatement := range incomeStatement.AnnualReport {
		fmt.Printf(`
		Fiscal Year End: %s
		Report Currency: %s
		Gross Profit: %s
		Total Revenue: %s
		Cost of Revenue: %s
		Cost of Goods and Services Sold: %s
		Operating Income: %s
		Selling General and Administrative: %s
		Research and Development: %s
		Operating Expenses: %s
		Investment Income: %s
		Net Interest Income: %s
		Interest Income: %s
		Interest Expense: %s
		Non-Interest Income: %s
		Other Non-Operative Income: %s
		Deprecation: %s
		Deprecation and Amortization: %s
		Income Before Tax: %s
		Income Tax Expenses: %s
		Interest and Debt Expenses: %s
		Net Income From Continuing Operation: %s
		Comprehensive Income: %s
		EBIT: %s
		EBITDA: %s
		Net Income: %s
		`, incomeStatement.FiscalEndDate, incomeStatement.Currency, incomeStatement.GrossProfit,
			incomeStatement.TotalRevenue, incomeStatement.CostOfRevenue, incomeStatement.CostOfGSSold,
			incomeStatement.OperatingIncome, incomeStatement.SellingGenAdmin, incomeStatement.RnD,
			incomeStatement.OperatingExpenses, incomeStatement.InvestmentIncome, incomeStatement.NetInterestIncome,
			incomeStatement.InterestIncome, incomeStatement.InterestExpense, incomeStatement.NonInterestIncome,
			incomeStatement.NonOperatingIncome, incomeStatement.Deprecation, incomeStatement.DepreciationAndAmortization,
			incomeStatement.IncomeBeforeTax, incomeStatement.IncomeTaxExpense, incomeStatement.InterestAndDebt,
			incomeStatement.NetIncomeFromContinuingOperations, incomeStatement.ComprehensiveIncomeNetOfTax,
			incomeStatement.EBIT, incomeStatement.EBITDA, incomeStatement.NetIncome)

	}
}

func PrintQuarterlyIncomeStatement(incomeStatement *api.IncomeStatement) {
	for _, incomeStatement := range incomeStatement.QuarterlyReport {
		fmt.Printf(`
		Fiscal Date End: %s
		Report Currency: %s
		Gross Profit: %s
		Total Revenue: %s
		Cost of Revenue: %s
		Cost of Goods and Services Sold: %s
		Operating Income: %s
		Selling General and Administrative: %s
		Research and Development: %s
		Operating Expenses: %s
		Investment Income: %s
		Net Interest Income: %s
		Interest Income: %s
		Interest Expense: %s
		Non-Interest Income: %s
		Other Non-Operative Income: %s
		Deprecation: %s
		Deprecation and Amortization: %s
		Income Before Tax: %s
		Income Tax Expenses: %s
		Interest and Debt Expenses: %s
		Net Income From Continuing Operation: %s
		Comprehensive Income: %s
		EBIT: %s
		EBITDA: %s
		Net Income: %s
		`, incomeStatement.FiscalEndDate, incomeStatement.Currency, incomeStatement.GrossProfit,
			incomeStatement.TotalRevenue, incomeStatement.CostOfRevenue, incomeStatement.CostOfGSSold,
			incomeStatement.OperatingIncome, incomeStatement.SellingGenAdmin, incomeStatement.RnD,
			incomeStatement.OperatingExpenses, incomeStatement.InvestmentIncome, incomeStatement.NetInterestIncome,
			incomeStatement.InterestIncome, incomeStatement.InterestExpense, incomeStatement.NonInterestIncome,
			incomeStatement.NonOperatingIncome, incomeStatement.Deprecation, incomeStatement.DepreciationAndAmortization,
			incomeStatement.IncomeBeforeTax, incomeStatement.IncomeTaxExpense, incomeStatement.InterestAndDebt,
			incomeStatement.NetIncomeFromContinuingOperations, incomeStatement.ComprehensiveIncomeNetOfTax,
			incomeStatement.EBIT, incomeStatement.EBITDA, incomeStatement.NetIncome)

	}
}

func PrintAnnualBalanceSheet(balanceSheet *api.BalanceSheet) {
	for _, balanceSheet := range balanceSheet.AnnualReport {
		fmt.Printf(`
		Fiscal Year End: %s
		Currency: %s
		Total Assets: %s
		Total Current Assets: %s
		Cash And Cash Equivalents At Carrying Value: %s
		Cash And Short Term Investments: %s
		Inventory: %s
		Current Net Receivables: %s
		Total Non-Current Assets: %s
		Property Plant Equipment: %s
		Accumulated Depreciation Amortization PPE: %s
		Intangible Assets: %s
		Intangible Assets Excluding Goodwill: %s
		Goodwill: %s
		Investments: %s
		Long Term Investments: %s
		Short Term Investments: %s
		Other Current Assets: %s
		Other Non Current Assets: %s
		Total Liabilities: %s
		Total Current Liabilities: %s
		Current Accounts Payable: %s
		Deferred Revenue: %s
		Current Debt: %s
		Short Term Debt: %s
		Total Non-Current Liabilities: %s
		Capital Lease Obligations: %s
		Long Term Debt: %s
		Current Long Term Debt: %s
		Long Term Debt Non-Current: %s
		Total Short Long Term Debt: %s
		Other Current Liabilities: %s
		Other Non-Current Liabilities: %s
		Total Shareholder Equity: %s
		Treasury Stock: %s
		Retained Earnings: %s
		Common Stock: %s
		Common Stock Shares Outstanding: %s 
		`, balanceSheet.FiscalEndDate, balanceSheet.Currency, balanceSheet.TotalAssets, balanceSheet.TotalCurrentAssets,
			balanceSheet.CashAndCashEquivalentsAtCarryingValue, balanceSheet.CashAndShortTermInvestments, balanceSheet.Inventory,
			balanceSheet.CurrentNetReceivables, balanceSheet.TotalNonCurrentAssets, balanceSheet.PropertyPlantEquipment,
			balanceSheet.AccumulatedDepreciationAmortizationPPE, balanceSheet.IntangibleAssets, balanceSheet.IntangibleAssetsExcludingGoodwill,
			balanceSheet.Goodwill, balanceSheet.Investments, balanceSheet.LongTermInvestments, balanceSheet.ShortTermInvestments,
			balanceSheet.OtherCurrentAssets, balanceSheet.OtherCurrentAssets, balanceSheet.TotalLiabilities, balanceSheet.TotalCurrentLiabilities,
			balanceSheet.CurrentAccountsPayable, balanceSheet.DeferredRevenue, balanceSheet.CurrentDebt, balanceSheet.ShortTermDebt,
			balanceSheet.TotalNonCurrentLiabilities, balanceSheet.CapitalLeaseObligations, balanceSheet.LongTermDebt,
			balanceSheet.CurrentLongTermDebt, balanceSheet.LongTermDebtNoncurrent, balanceSheet.ShortLongTermDebtTotal,
			balanceSheet.OtherCurrentLiabilities, balanceSheet.OtherNonCurrentLiabilities, balanceSheet.TotalShareholderEquity,
			balanceSheet.TreasuryStock, balanceSheet.RetainedEarnings, balanceSheet.CommonStock, balanceSheet.CommonStockSharesOutstanding)
	}
}

func PrintQuarterlyBalanceSheet(balanceSheet *api.BalanceSheet) {
	for i := len(balanceSheet.AnnualReport) - 1; i >= 0; i-- {
		balanceSheet := balanceSheet.AnnualReport[i]
		fmt.Printf(`
		Fiscal Date End: %s
		Currency: %s
		Total Assets: %s
		Total Current Assets: %s
		Cash And Cash Equivalents At Carrying Value: %s
		Cash And Short Term Investments: %s
		Inventory: %s
		Current Net Receivables: %s
		Total Non-Current Assets: %s
		Property Plant Equipment: %s
		Accumulated Depreciation Amortization PPE: %s
		Intangible Assets: %s
		Intangible Assets Excluding Goodwill: %s
		Goodwill: %s
		Investments: %s
		Long Term Investments: %s
		Short Term Investments: %s
		Other Current Assets: %s
		Other Non Current Assets: %s
		Total Liabilities: %s
		Total Current Liabilities: %s
		Current Accounts Payable: %s
		Deferred Revenue: %s
		Current Debt: %s
		Short Term Debt: %s
		Total Non-Current Liabilities: %s
		Capital Lease Obligations: %s
		Long Term Debt: %s
		Current Long Term Debt: %s
		Long Term Debt Non-Current: %s
		Total Short Long Term Debt: %s
		Other Current Liabilities: %s
		Other Non-Current Liabilities: %s
		Total Shareholder Equity: %s
		Treasury Stock: %s
		Retained Earnings: %s
		Common Stock: %s
		Common Stock Shares Outstanding: %s 
		`, balanceSheet.FiscalEndDate, balanceSheet.Currency, balanceSheet.TotalAssets, balanceSheet.TotalCurrentAssets,
			balanceSheet.CashAndCashEquivalentsAtCarryingValue, balanceSheet.CashAndShortTermInvestments, balanceSheet.Inventory,
			balanceSheet.CurrentNetReceivables, balanceSheet.TotalNonCurrentAssets, balanceSheet.PropertyPlantEquipment,
			balanceSheet.AccumulatedDepreciationAmortizationPPE, balanceSheet.IntangibleAssets, balanceSheet.IntangibleAssetsExcludingGoodwill,
			balanceSheet.Goodwill, balanceSheet.Investments, balanceSheet.LongTermInvestments, balanceSheet.ShortTermInvestments,
			balanceSheet.OtherCurrentAssets, balanceSheet.OtherCurrentAssets, balanceSheet.TotalLiabilities, balanceSheet.TotalCurrentLiabilities,
			balanceSheet.CurrentAccountsPayable, balanceSheet.DeferredRevenue, balanceSheet.CurrentDebt, balanceSheet.ShortTermDebt,
			balanceSheet.TotalNonCurrentLiabilities, balanceSheet.CapitalLeaseObligations, balanceSheet.LongTermDebt,
			balanceSheet.CurrentLongTermDebt, balanceSheet.LongTermDebtNoncurrent, balanceSheet.ShortLongTermDebtTotal,
			balanceSheet.OtherCurrentLiabilities, balanceSheet.OtherNonCurrentLiabilities, balanceSheet.TotalShareholderEquity,
			balanceSheet.TreasuryStock, balanceSheet.RetainedEarnings, balanceSheet.CommonStock, balanceSheet.CommonStockSharesOutstanding)
	}
}
