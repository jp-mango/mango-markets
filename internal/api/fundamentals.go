package api

import (
	"encoding/json"
	"fmt"
)

type CompanyInfo struct {
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
	MarketCapitalization       string `json:"MarketCapitalization"`
	EBITDA                     string `json:"EBITDA"`
	PERatio                    string `json:"PERatio"`
	PEGRatio                   string `json:"PEGRatio"`
	BookValue                  string `json:"BookValue"`
	DividendPerShare           string `json:"DividendPerShare"`
	DividendYield              string `json:"DividendYield"`
	EPS                        string `json:"EPS"`
	RevenuePerShareTTM         string `json:"RevenuePerShareTTM"`
	ProfitMargin               string `json:"ProfitMargin"`
	OperatingMarginTTM         string `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          string `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          string `json:"ReturnOnEquityTTM"`
	RevenueTTM                 string `json:"RevenueTTM"`
	GrossProfitTTM             string `json:"GrossProfitTTM"`
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
	PriceToSalesRatioTTM       string `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           string `json:"PriceToBookRatio"`
	EVToRevenue                string `json:"EVToRevenue"`
	EVToEBITDA                 string `json:"EVToEBITDA"`
	Beta                       string `json:"Beta"`
	FiftyTwoWeekHigh           string `json:"52WeekHigh"`
	FiftyTwoWeekLow            string `json:"52WeekLow"`
	FiftyDayMovingAverage      string `json:"50DayMovingAverage"`
	TwoHundredDayMovingAverage string `json:"200DayMovingAverage"`
	SharesOutstanding          string `json:"SharesOutstanding"`
	DividendDate               string `json:"DividendDate"`
	ExDividendDate             string `json:"ExDividendDate"`
}

func FetchCompanyInfo(ticker, apiKey string) (*CompanyInfo, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=OVERVIEW&symbol=%s&apikey=%s", ticker, apiKey)

	var companyInfo CompanyInfo

	data, err := DataPull(url)
	if err != nil {
		return nil, fmt.Errorf("err: %s", err)
	}

	err = json.Unmarshal(data, &companyInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %v", err)

	}

	return &companyInfo, nil
}

type CompanyFinancials struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []AnnualReport    `json:"annualReports"`
	QuarterlyReports []QuarterlyReport `json:"quarterlyReports"`
}

type AnnualReport struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       string `json:"grossProfit"`
	TotalRevenue                      string `json:"totalRevenue"`
	CostOfRevenue                     string `json:"costOfRevenue"`
	CostofGoodsAndServicesSold        string `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   string `json:"operatingIncome"`
	SellingGeneralAndAdministrative   string `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            string `json:"researchAndDevelopment"`
	OperatingExpenses                 string `json:"operatingExpenses"`
	InvestmentIncomeNet               string `json:"investmentIncomeNet"`
	NetInterestIncome                 string `json:"netInterestIncome"`
	InterestIncome                    string `json:"interestIncome"`
	InterestExpense                   string `json:"interestExpense"`
	NonInterestIncome                 string `json:"nonInterestIncome"`
	OtherNonOperatingIncome           string `json:"otherNonOperatingIncome"`
	Depreciation                      string `json:"depreciation"`
	DepreciationAndAmortization       string `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   string `json:"incomeBeforeTax"`
	IncomeTaxExpense                  string `json:"incomeTaxExpense"`
	InterestAndDebtExpense            string `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations string `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       string `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              string `json:"ebit"`
	Ebitda                            string `json:"ebitda"`
	NetIncome                         string `json:"netIncome"`
}

type QuarterlyReport struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       string `json:"grossProfit"`
	TotalRevenue                      string `json:"totalRevenue"`
	CostOfRevenue                     string `json:"costOfRevenue"`
	CostofGoodsAndServicesSold        string `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   string `json:"operatingIncome"`
	SellingGeneralAndAdministrative   string `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            string `json:"researchAndDevelopment"`
	OperatingExpenses                 string `json:"operatingExpenses"`
	InvestmentIncomeNet               string `json:"investmentIncomeNet"`
	NetInterestIncome                 string `json:"netInterestIncome"`
	InterestIncome                    string `json:"interestIncome"`
	InterestExpense                   string `json:"interestExpense"`
	NonInterestIncome                 string `json:"nonInterestIncome"`
	OtherNonOperatingIncome           string `json:"otherNonOperatingIncome"`
	Depreciation                      string `json:"depreciation"`
	DepreciationAndAmortization       string `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   string `json:"incomeBeforeTax"`
	IncomeTaxExpense                  string `json:"incomeTaxExpense"`
	InterestAndDebtExpense            string `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations string `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       string `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              string `json:"ebit"`
	Ebitda                            string `json:"ebitda"`
	NetIncome                         string `json:"netIncome"`
}

func FetchIncomeStatement(ticker, apiKey string) (*CompanyFinancials, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=INCOME_STATEMENT&symbol=%s&apikey=%s", ticker, apiKey)

	var incomeStatement CompanyFinancials

	data, err := DataPull(url)
	if err != nil {
		return nil, fmt.Errorf("err: %s", err)
	}

	err = json.Unmarshal(data, &incomeStatement)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &incomeStatement, nil
}
