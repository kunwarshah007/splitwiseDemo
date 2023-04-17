package models

type Expense struct {
	GroupId             string    `json:"group_id"`
	PaidByUserId        string    `json:"paid_by_user_id"`
	Amount              float64   `json:"amount"`
	UserIdsThatOwsShare []string  `json:"user_ids_that_ows_share"`
	Type                string    `json:"type"` //Could either be EQUAL, EXACT or PERCENT
	Share               []float64 `json:"share"`
}
