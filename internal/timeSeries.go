package internal

// Intraday Data

type Metadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	Size          string `json:"5. Output Size"`
	TZ            string `json:"6. Time Zone"`
}

type IntradayTSD struct {
}

func intradayData(apiKey, ticker, interval string) {

}
