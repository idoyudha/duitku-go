package invoice

import (
	"context"

	"github.com/idoyudha/duitku-go/common"
)

type InvoiceService struct {
	config *common.Config
}

func NewInvoiceService(config *common.Config) InvoiceService {
	return InvoiceService{
		config: config,
	}
}

func (s *InvoiceService) Create(ctx context.Context, req CreateInvoiceRequest) (CreateInvoiceResponse, error) {
	var response CreateInvoiceResponse
	return response, nil
}
