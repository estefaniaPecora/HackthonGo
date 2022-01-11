package sales

import (
	"github.com/estefaniaPecora/HackthonGo/internal/models"
)

type Service interface {
	SaveSales(quantity float64, idinvoice, idproduct int) (models.Sale, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (serv *service) SaveSales(quantity float64, idinvoice, idproduct int) (models.Sale, error) {

	newSale := models.Sale{Quantity: quantity, Id_invoice: idinvoice, Id_product: idproduct}
	saleCreated, err := serv.repository.SaveSales(newSale)
	if err != nil {
		return models.Sale{}, err
	}
	return saleCreated, nil

}
