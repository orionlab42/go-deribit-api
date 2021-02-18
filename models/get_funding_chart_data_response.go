package models

type GetFundingChartData struct {
	IndexPrice float64 `json:"index_price"`
	Interest8H float64 `json:"interest_8h"`
	Timestamp  int64   `json:"timestamp"`
}

type GetFundingChartDataResponse struct {
	CurrentInterest float64               `json:"current_interest"`
	Data            []GetFundingChartData `json:"data"`
	Interest8H      float64               `json:"interest_8h"`
}
