package invoice

import (
	"context"
	"net/http"

	"github.com/idoyudha/duitku-go/common"
)

type InvoiceService struct {
	client *common.ServiceClient
}

func NewInvoiceService(service *common.ServiceClient) *InvoiceService {
	return &InvoiceService{
		client: service,
	}
}

func (s *InvoiceService) Create(ctx context.Context, req CreateInvoiceRequest) (CreateInvoiceResponse, *http.Response, error) {
	res := &CreateInvoiceResponse{}
	path := "/merchant/createInvoice"
	headerParams := make(map[string]string)
	headerParams[common.MerchantCodeHeader] = s.client.Cfg.MerchantCode
	headerParams[common.TimeStampHeader] = s.client.GetCurrentTimestamp()
	headerParams[common.SignatureHeader] = s.client.CreateSignature(headerParams[common.TimeStampHeader])

	baseUrl := common.SandboxPOPBaseURL
	if s.client.Cfg.Environment == common.ProductionEnv {
		baseUrl = common.ProductionPOPBaseURL
	}

	httpRes, err := common.SendAPIRequest(
		ctx,
		s.client,
		req,
		res,
		http.MethodPost,
		baseUrl+path,
		headerParams,
	)
	return *res, httpRes, err
}
