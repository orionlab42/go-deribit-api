package models

type GetFundingRateValueParams struct {
	InstrumentName string `json:"instrument_name"`
	StartTimestamp int64  `json:"start_timestamp"`
	EndTimestamp   int64  `json:"end_timestamp"`
}
