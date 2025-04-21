package duitku

import (
	"github.com/idoyudha/duitku-go/common"
	"github.com/idoyudha/duitku-go/invoice"
)

type APIClient struct {
	*common.ServiceClient
	// API Services
	InvoiceService *invoice.InvoiceService
}

func NewClient(cfg *common.Config) *APIClient {
	c := &APIClient{}
	c.Cfg = cfg

	// Add API Services
	c.InvoiceService = invoice.NewInvoiceService(c.ServiceClient)

	return c
}
