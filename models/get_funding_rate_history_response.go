package models

type GetFundingRateHistoryData struct {
	Timestamp      int64   `json:"timestamp"`
	IndexPrice     float64 `json:"index_price"`
	PrevIndexPrice float64 `json:"prev_index_price"`
	Interest1H     float64 `json:"interest_1h"`
	Interest8H     float64 `json:"interest_8h"`
}

type GetFundingRateHistoryResponse []GetFundingRateHistoryData
