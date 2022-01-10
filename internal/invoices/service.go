package invoices

import (
	"github.com/estefaniaPecora/HackthonGo/internal/models"
)

type Service interface {
	SaveData(date_time string, total float64, id_customer int) (models.Invoice, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (serv *service) SaveData(date_time string, total float64, id_customer int) (models.Invoice, error) {
	newInvoice := models.Invoice{DateTime: date_time, Total: total, IdCustomer: id_customer}
	invoiceCreated, err := serv.repository.SaveData(newInvoice)
	if err != nil {
		return models.Invoice{}, err
	}
	return invoiceCreated, nil

}
