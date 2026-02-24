package dto

type SendDailyAssignmentRequest struct {
	PositionID   int64  `json:"position_id"`
	PositionName string `json:"position_name"`
	Description  string `json:"description"`
}

type SendDailyAssignmentResponse struct {
	SentCount int `json:"sent_count"`
}
