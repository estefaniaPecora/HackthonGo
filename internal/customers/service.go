package customers

import (
	"github.com/estefaniaPecora/HackthonGo/internal/models"
)

type Service interface {
	SaveCustomer(last_name string, first_name string, condition string) (models.Customer, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (serv *service) SaveCustomer(last_name string, first_name string, condition string) (models.Customer, error) {
	newCustomer := models.Customer{LastName: last_name, FirstName: first_name, Condition: condition}
	customerCreated, err := serv.repository.SaveCustomer(newCustomer)
	if err != nil {
		return models.Customer{}, err
	}
	return customerCreated, nil

}
