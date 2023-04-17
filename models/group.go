package models

type Group struct {
	GroupId                 string                        `json:"group_id"`
	Users                   []User                        `json:"user_ids"`
	Name                    string                        `json:"name"`
	GroupTransactionDetails map[string]map[string]float64 `json:"group_transaction_details"`
}
