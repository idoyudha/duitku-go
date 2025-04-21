package duitku

import (
	"github.com/idoyudha/duitku-go/common"
	"github.com/idoyudha/duitku-go/invoice"
)

type APIClient struct {
	client *common.Config
	// API Services
	InvoiceService invoice.InvoiceService
}

func NewClient(cfg *common.Config) *APIClient {
	c := &APIClient{}
	c.client = cfg

	// Add API Services
	c.InvoiceService = invoice.NewInvoiceService(cfg)

	return c
}
