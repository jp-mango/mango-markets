package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TimeSeriesProvider is an interface for fetching time series data.
type TimeSeriesData interface {
	GetTimeSeriesData() map[string]map[string]string
}

// ! Time series daily
func (tsd TimeSeriesDaily) GetTimeSeriesData() map[string]map[string]string {
	return tsd.TimeSeries
}

type TimeSeriesDaily struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		OutputSize    string `json:"4. Output Size"`
		TimeZone      string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Time Series (Daily)"`
}

func FetchTimeSeriesDaily(apiKey, ticker string) (TimeSeriesDaily, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var tsdData TimeSeriesDaily
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &tsdData); err != nil {
		return TimeSeriesDaily{}, fmt.Errorf("error parsing API content:%e", err)
	}
	return tsdData, err
}

// ! Time series weekly
func (tsw TimeSeriesWeekly) GetTimeSeriesData() map[string]map[string]string {
	return tsw.TimeSeries
}

type TimeSeriesWeekly struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		TimeZone      string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Weekly Adjusted Time Series"`
}

func FetchTimeSeriesWeekly(apiKey, ticker string) (TimeSeriesWeekly, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY_ADJUSTED&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var tswData TimeSeriesWeekly
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &tswData); err != nil {
		return TimeSeriesWeekly{}, fmt.Errorf("error parsing API content:%e", err)
	}
	return tswData, err
}

// ! Time series monthly
func (tsm TimeSeriesMonthly) GetTimeSeriesData() map[string]map[string]string {
	return tsm.TimeSeries
}

type TimeSeriesMonthly struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		TimeZone      string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Monthly Adjusted Time Series"`
}

func FetchTimeSeriesMonthly(apiKey, ticker string) (TimeSeriesMonthly, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY_ADJUSTED&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to hit endpoint:", err)
	}
	defer resp.Body.Close()

	var tsmData TimeSeriesMonthly
	content, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(content, &tsmData); err != nil {
		return TimeSeriesMonthly{}, fmt.Errorf("error parsing API content:%e", err)
	}
	return tsmData, err
}

// ! Top Gainers and Losers
type TopGainLoss struct {
	Metadata           string   `json:"metadata"`
	LastUpdated        string   `json:"last_updated"`
	TopGainers         []Ticker `json:"top_gainers"`
	TopLosers          []Ticker `json:"top_losers"`
	MostActivelyTraded []Ticker `json:"most_actively_traded"`
}

type Ticker struct {
	Ticker           string `json:"ticker"`
	Price            string `json:"price"`
	ChangeAmount     string `json:"change_amount"`
	ChangePercentage string `json:"change_percentage"`
	Volume           string `json:"volume"`
}

func FetchGainLoss(apiKey string) (*TopGainLoss, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TOP_GAINERS_LOSERS&apikey=%s", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	var marketData TopGainLoss
	err = json.Unmarshal(body, &marketData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %s", err)
	}

	return &marketData, nil
}

// ! Global Market Hours
type MarketHours struct {
	Endpoint string   `json:"endpoint"`
	Markets  []Market `json:"markets"`
}

type Market struct {
	MarketType       string `json:"market_type"`
	Region           string `json:"region"`
	PrimaryExchanges string `json:"primary_exchanges"`
	LocalOpen        string `json:"local_open"`
	LocalClose       string `json:"local_close"`
	CurrentStatus    string `json:"current_status"`
	Notes            string `json:"notes"`
}

func FetchMarketStatus(apiKey string) (*MarketHours, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=MARKET_STATUS&apikey=%s", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	var marketStatus MarketHours
	if err := json.Unmarshal(body, &marketStatus); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}

	return &marketStatus, nil
}

//! Company Overview

type CompanyOverview struct {
	Symbol                     string `json:"Symbol"`
	AssetType                  string `json:"AssetType"`
	Name                       string `json:"Name"`
	Description                string `json:"Description"`
	CIK                        string `json:"CIK"`
	Exchange                   string `json:"Exchange"`
	Currency                   string `json:"Currency"`
	Country                    string `json:"Country"`
	Sector                     string `json:"Sector"`
	Industry                   string `json:"Industry"`
	Address                    string `json:"Address"`
	FiscalYearEnd              string `json:"FiscalYearEnd"`
	LatestQuarter              string `json:"LatestQuarter"`
	MarketCap                  string `json:"MarketCapitalization"`
	EBITDA                     string `json:"EBITDA"`
	PERatio                    string `json:"PERatio"`
	PEGRatio                   string `json:"PEGRatio"`
	BookVal                    string `json:"BookValue"`
	Dividend                   string `json:"DividendPerShare"`
	DividendYield              string `json:"DividendYield"`
	EPS                        string `json:"EPS"`
	ShareRevenue               string `json:"RevenuePerShareTTM"`
	ProfitMargin               string `json:"ProfitMargin"`
	OperatingMargin            string `json:"OperatingMarginTTM"`
	ROA                        string `json:"ReturnOnAssetsTTM"`
	ROE                        string `json:"ReturnOnEquityTTM"`
	Revenue                    string `json:"RevenueTTM"`
	GrossProfit                string `json:"GrossProfitTTM"`
	DilutedEPSTTM              string `json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY string `json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  string `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         string `json:"AnalystTargetPrice"`
	AnalystRatingStrongBuy     string `json:"AnalystRatingStrongBuy"`
	AnalystRatingBuy           string `json:"AnalystRatingBuy"`
	AnalystRatingHold          string `json:"AnalystRatingHold"`
	AnalystRatingSell          string `json:"AnalystRatingSell"`
	AnalystRatingStrongSell    string `json:"AnalystRatingStrongSell"`
	TrailingPE                 string `json:"TrailingPE"`
	ForwardPE                  string `json:"ForwardPE"`
	PriceToSales               string `json:"PriceToSalesRatioTTM"`
	PriceToBook                string `json:"PriceToBookRatio"`
	EVToRevenue                string `json:"EVToRevenue"`
	EVToEBITDA                 string `json:"EVToEBITDA"`
	Beta                       string `json:"Beta"`
	YearHigh                   string `json:"52WeekHigh"`
	YearLow                    string `json:"52WeekLow"`
	FiddyDayMA                 string `json:"50DayMovingAverage"`
	TwoHunnaDayMA              string `json:"200DayMovingAverage"`
	SharesOutstanding          string `json:"SharesOutstanding"`
	DividendDate               string `json:"DividendDate"`
	ExDividendDate             string `json:"ExDividendDate"`
}

func FetchCompanyOverview(ticker, apiKey string) (*CompanyOverview, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=OVERVIEW&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	var companyInfo CompanyOverview
	if err := json.Unmarshal(body, &companyInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}
	return &companyInfo, nil
}

//! Income statement

type IncomeStatement struct {
	Symbol          string         `json:"symbol"`
	AnnualReport    []IncomeReport `json:"annualReports"`
	QuarterlyReport []IncomeReport `json:"quarterlyReports"`
}

type IncomeReport struct {
	FiscalEndDate                     string `json:"fiscalDateEnding"`
	Currency                          string `json:"reportedCurrency"`
	GrossProfit                       string `json:"grossProfit"`
	TotalRevenue                      string `json:"totalRevenue"`
	CostOfRevenue                     string `json:"costOfRevenue"`
	CostOfGSSold                      string `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   string `json:"operatingIncome"`
	SellingGenAdmin                   string `json:"sellingGeneralAndAdministrative"`
	RnD                               string `json:"researchAndDevelopment"`
	OperatingExpenses                 string `json:"operatingExpenses"`
	InvestmentIncome                  string `json:"investmentIncomeNet"`
	NetInterestIncome                 string `json:"netInterestIncome"`
	InterestIncome                    string `json:"interestIncome"`
	InterestExpense                   string `json:"interestExpense"`
	NonInterestIncome                 string `json:"nonInterestIncome"`
	NonOperatingIncome                string `json:"otherNonOperatingIncome"`
	Deprecation                       string `json:"depreciation"`
	DepreciationAndAmortization       string `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   string `json:"incomeBeforeTax"`
	IncomeTaxExpense                  string `json:"incomeTaxExpense"`
	InterestAndDebt                   string `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations string `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       string `json:"comprehensiveIncomeNetOfTax"`
	EBIT                              string `json:"ebit"`
	EBITDA                            string `json:"ebitda"`
	NetIncome                         string `json:"netIncome"`
}

func FetchIncomeStatement(ticker, apiKey string) (*IncomeStatement, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=INCOME_STATEMENT&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	var incomeStatement IncomeStatement
	if err := json.Unmarshal(body, &incomeStatement); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}
	return &incomeStatement, nil
}

//! Balance Sheet

type BalanceSheet struct {
	Symbol          string               `json:"symbol"`
	AnnualReport    []BalanceSheetReport `json:"annualReports"`
	QuarterlyReport []BalanceSheetReport `json:"quarterlyReports"`
}

type BalanceSheetReport struct {
	FiscalEndDate                          string `json:"fiscalDateEnding"`
	Currency                               string `json:"reportedCurrency"`
	TotalAssets                            string `json:"totalAssets"`
	TotalCurrentAssets                     string `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  string `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            string `json:"cashAndShortTermInvestments"`
	Inventory                              string `json:"inventory"`
	CurrentNetReceivables                  string `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  string `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 string `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE string `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       string `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      string `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               string `json:"goodwill"`
	Investments                            string `json:"investments"`
	LongTermInvestments                    string `json:"longTermInvestments"`
	ShortTermInvestments                   string `json:"shortTermInvestments"`
	OtherCurrentAssets                     string `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  string `json:"otherNonCurrentAssets"`
	TotalLiabilities                       string `json:"totalLiabilities"`
	TotalCurrentLiabilities                string `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 string `json:"currentAccountsPayable"`
	DeferredRevenue                        string `json:"deferredRevenue"`
	CurrentDebt                            string `json:"currentDebt"`
	ShortTermDebt                          string `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             string `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                string `json:"capitalLeaseObligations"`
	LongTermDebt                           string `json:"longTermDebt"`
	CurrentLongTermDebt                    string `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 string `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 string `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                string `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             string `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 string `json:"totalShareholderEquity"`
	TreasuryStock                          string `json:"treasuryStock"`
	RetainedEarnings                       string `json:"retainedEarnings"`
	CommonStock                            string `json:"commonStock"`
	CommonStockSharesOutstanding           string `json:"commonStockSharesOutstanding"`
}

func FetchBalanceSheet(ticker, apiKey string) (*BalanceSheet, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=BALANCE_SHEET&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	var balanceSheet BalanceSheet
	if err := json.Unmarshal(body, &balanceSheet); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}
	return &balanceSheet, nil
}

// ! Cashflow
type CashFlow struct {
	Symbol          string           `json:"symbol"`
	AnnualReport    []CashFlowReport `json:"annualReports"`
	QuarterlyReport []CashFlowReport `json:"quarterlyReports"`
}

type CashFlowReport struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	Currency                                                  string `json:"reportedCurrency"`
	OperatingCashFlow                                         string `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            string `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           string `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              string `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   string `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      string `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       string `json:"capitalExpenditures"`
	ChangeInReceivables                                       string `json:"changeInReceivables"`
	ChangeInInventory                                         string `json:"changeInInventory"`
	ProfitLoss                                                string `json:"profitLoss"`
	CashflowFromInvestment                                    string `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     string `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     string `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        string `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             string `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     string `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            string `json:"dividendPayout"`
	DividendPayoutCommonStock                                 string `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              string `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         string `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet string `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      string `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            string `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           string `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            string `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      string `json:"changeInExchangeRate"`
	NetIncome                                                 string `json:"netIncome"`
}

func FetchCashflow(ticker, apiKey string) (*CashFlow, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=CASH_FLOW&symbol=%s&apikey=%s", ticker, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	var cashFlow CashFlow
	if err := json.Unmarshal(body, &cashFlow); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}
	return &cashFlow, nil
}
