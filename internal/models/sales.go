package models

type Sale struct {
	ID         int     `json:"id"`
	Id_product int     `json:"id_product"`
	Id_invoice int     `json:"id_invoice"`
	Quantity   float64 `json:"quantity"`
}
