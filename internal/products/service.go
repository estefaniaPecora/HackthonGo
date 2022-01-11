package products

import (
	"github.com/estefaniaPecora/HackthonGo/internal/models"
)

type Service interface {
	SaveProducts(description string, price float64) (models.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (serv *service) SaveProducts(description string, price float64) (models.Product, error) {

	newProduct := models.Product{Description: description, Price: price}
	productCreated, err := serv.repository.SaveProducts(newProduct)
	if err != nil {
		return models.Product{}, err
	}
	return productCreated, nil

}
