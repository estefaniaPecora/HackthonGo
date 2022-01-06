package models

type Sales struct {
	ID         int     `json:"id"`
	DateTime   string  `json:"date_time"`
	IdCustomer int     `json:"id_customer"`
	Total      float64 `json:"total"`
}
