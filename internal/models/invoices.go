package models

type Invoice struct {
	ID         int     `json:"id"`
	DateTime   string  `json:"date_time"`
	Total      float64 `json:"total"`
	IdCustomer int     `json:"id_customer"`
}
