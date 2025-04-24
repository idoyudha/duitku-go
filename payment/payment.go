package payment

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/idoyudha/duitku-go/common"
)

type PaymentService struct {
	client *common.ServiceClient
}

// NewPaymentService returns a new instance of PaymentService with the given service client.
func NewPaymentService(service *common.ServiceClient) *PaymentService {
	return &PaymentService{
		client: service,
	}
}

func (s *PaymentService) GetMethods(ctx context.Context, req GetPaymentMethodRequest) (GetPaymentMethodResponse, *http.Response, error) {
	res := &GetPaymentMethodResponse{}
	path := "/merchant/paymentmethod/getpaymentmethod"

	baseUrl := common.SandboxV2BaseURL
	if s.client.Cfg.Environment == common.ProductionEnv {
		baseUrl = common.ProductionV2BaseURL
	}

	req.Signature = s.generatePaymentSignature(strconv.Itoa(req.Amount) + req.Datetime)

	httpRes, err := common.SendAPIRequest(
		ctx,
		s.client,
		req,
		res,
		http.MethodPost,
		baseUrl+path,
		nil,
	)

	return *res, httpRes, err
}

func (s *PaymentService) generatePaymentSignature(parameter string) string {
	combinedStr := s.client.Cfg.MerchantCode + parameter + s.client.Cfg.APIKey
	hash := sha256.Sum256([]byte(combinedStr))
	return hex.EncodeToString(hash[:])
}
