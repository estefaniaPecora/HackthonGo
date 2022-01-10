package customers

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/estefaniaPecora/HackthonGo/internal/models"
)

type Service interface {
	SaveCustomer(last_name string, first_name string, condition string) (models.Customer, error)
	LoadAllCustomers(file string) error
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

func (serv *service) LoadAllCustomers(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()

	readItems := bufio.NewScanner(file)

	for readItems.Scan() {
		line := readItems.Text()
		items := strings.Split(line, "#$%#")

		customer := models.Customer{LastName: items[1], FirstName: items[2], Condition: items[3]}
		_, err := serv.repository.SaveCustomer(customer)
		if err != nil {
			log.Fatal(err)
		}

	}

	return nil

}
