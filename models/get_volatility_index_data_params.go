package models

type GetVolatilityIndexDataParams struct {
	Currency       string `json:"currency"`
	StartTimestamp int64  `json:"start_timestamp"`
	EndTimestamp   int64  `json:"end_timestamp"`
	Resolution     string `json:"resolution"`
}
