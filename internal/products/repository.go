package products

import (
	"database/sql"

	"log"

	"github.com/estefaniaPecora/HackthonGo/internal/models"
	"github.com/estefaniaPecora/HackthonGo/pkg/db"
)

type Repository interface {
	SaveProducts(product models.Product) (models.Product, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) SaveProducts(product models.Product) (models.Product, error) {
	DbStarted := db.StorageDB
	stmt, err := DbStarted.Prepare("INSERT INTO products(description, price) VALUES(?, ?)")

	if err != nil {
		log.Fatal(err)
		return models.Product{}, err
	}
	defer stmt.Close()

	var result sql.Result

	result, err = stmt.Exec(product.Description, product.Price)

	if err != nil {
		log.Fatal(err)
		return models.Product{}, err
	}

	id_Creado, _ := result.LastInsertId()
	product.ID = int(id_Creado)

	return product, nil

}
