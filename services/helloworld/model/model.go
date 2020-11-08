package model

// Advice :
type Advice struct {
	ID         int64  `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

// AdviceListQueryObj :
type AdviceListQueryObj struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Start     int64  `json:"start"`
	Size      int64  `json:"size"`
}
