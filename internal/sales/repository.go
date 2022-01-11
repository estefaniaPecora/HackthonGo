package sales

import (
	"database/sql"

	"log"

	"github.com/estefaniaPecora/HackthonGo/internal/models"
	"github.com/estefaniaPecora/HackthonGo/pkg/db"
)

type Repository interface {
	SaveSales(sale models.Sale) (models.Sale, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) SaveSales(sale models.Sale) (models.Sale, error) {
	DbStarted := db.StorageDB
	stmt, err := DbStarted.Prepare("INSERT INTO sales(idinvoice, idproduct, quantity) VALUES(?, ?, ?)")

	if err != nil {
		log.Fatal(err)
		return models.Sale{}, err
	}
	defer stmt.Close()

	var result sql.Result

	result, err = stmt.Exec(sale.Id_invoice, sale.Id_product, sale.Quantity)

	if err != nil {
		log.Fatal(err)
		return models.Sale{}, err
	}

	id_Creado, _ := result.LastInsertId()
	sale.ID = int(id_Creado)

	return sale, nil

}
