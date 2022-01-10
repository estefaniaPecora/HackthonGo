package customers

import (
	"database/sql"
	"fmt"

	"log"

	"github.com/estefaniaPecora/HackthonGo/internal/models"
	"github.com/estefaniaPecora/HackthonGo/pkg/db"
)

type Repository interface {
	SaveCustomer(customer models.Customer) (models.Customer, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) SaveCustomer(customer models.Customer) (models.Customer, error) {
	DbStarted := db.StorageDB
	stmt, err := DbStarted.Prepare("INSERT INTO customers(last_name, first_name, customer_state) VALUES(?, ?, ?)")

	if err != nil {
		fmt.Println("ERRRORRRR REPOOOOOOOOOOO")
		log.Fatal(err)
		return models.Customer{}, err
	}
	defer stmt.Close()

	var result sql.Result

	result, err = stmt.Exec(customer.LastName, customer.FirstName, customer.Condition)

	if err != nil {
		log.Fatal(err)
		return models.Customer{}, err
	}

	id_Creado, _ := result.LastInsertId()
	customer.ID = int(id_Creado)

	return customer, nil

}
