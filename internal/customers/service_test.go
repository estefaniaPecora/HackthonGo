package customers

import (
	"testing"

	"github.com/estefaniaPecora/HackthonGo/internal/models"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
}

func TestStoresServiceOK(t *testing.T) {
	newCustomer := models.Customer{
		LastName:  "Costanza",
		FirstName: "George",
		Condition: "active",
	}

	repo := NewRepository()
	service := NewService(repo)

	customerCreated, _ := service.SaveCustomer(newCustomer.LastName, newCustomer.FirstName, newCustomer.Condition)

	assert.Equal(t, newCustomer.LastName, customerCreated.LastName)

}
