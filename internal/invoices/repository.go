package invoices

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/estefaniaPecora/HackthonGo/internal/models"
	"github.com/estefaniaPecora/HackthonGo/pkg/db"
)

type Repository interface {
	SaveData(invoice models.Invoice) (models.Invoice, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) SaveData(invoice models.Invoice) (models.Invoice, error) {
	DbStarted := db.StorageDB
	stmt, err := DbStarted.Prepare("INSERT INTO invoices(datetime, total, idcustomer) VALUES(?, ?, ?)")

	if err != nil {

		log.Fatal(err)
		return models.Invoice{}, err
	}

	var result sql.Result

	result, err = stmt.Exec(invoice.DateTime, invoice.Total, invoice.IdCustomer)

	if err != nil {
		fmt.Println("ACA ENTRO REPO EROOR")
		log.Fatal(err)
		return models.Invoice{}, err
	}
	defer stmt.Close()

	id_Creado, _ := result.LastInsertId()
	invoice.ID = int(id_Creado)

	return invoice, nil

}
