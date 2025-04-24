package duitku

import (
	"github.com/idoyudha/duitku-go/common"
	"github.com/idoyudha/duitku-go/invoice"
	"github.com/idoyudha/duitku-go/payment"
	"github.com/idoyudha/duitku-go/transaction"
)

type APIClient struct {
	*common.ServiceClient
	// API Services
	InvoiceService     *invoice.InvoiceService
	TransactionService *transaction.TransactionService
	PaymentService     *payment.PaymentService
}

func NewClient(cfg *common.Config) *APIClient {
	c := &APIClient{
		ServiceClient: &common.ServiceClient{
			Cfg: cfg,
		},
	}

	// Add API Services
	c.InvoiceService = invoice.NewInvoiceService(c.ServiceClient)
	c.TransactionService = transaction.NewTransactionService(c.ServiceClient)
	c.PaymentService = payment.NewPaymentService(c.ServiceClient)

	return c
}
