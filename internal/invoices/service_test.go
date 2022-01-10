package invoices

import (
	"testing"

	"github.com/estefaniaPecora/HackthonGo/internal/models"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
}

func TestStoresServiceOK(t *testing.T) {
	newInvoice := models.Invoice{
		DateTime:   "2019",
		Total:      100.50,
		IdCustomer: 1,
	}

	repo := NewRepository()
	service := NewService(repo)

	invoiceCreated, _ := service.SaveData(newInvoice.DateTime, newInvoice.Total, newInvoice.IdCustomer)

	assert.Equal(t, newInvoice.Total, invoiceCreated.Total)

}
