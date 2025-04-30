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

// GetMethods gets the available payment methods for the given amount and datetime.
// The method returns GetPaymentMethodResponse, HTTP response, and error.
// The GetPaymentMethodResponse contains the available payment methods.
//
// The request body should match the GetPaymentMethodRequest struct.
// The method will return error if the request body is empty or if the request
// body is invalid.
//
// The method will also return error if the request to the server fails or if
// the response from the server is invalid.
//
// The method will not return error if the request to the server is successful
// and the response from the server is valid.
//
// The method will send a POST request to the server with the given request
// body and headers.
//
// The method will use the sandbox base URL if the environment in the config
// is set to sandbox. Otherwise, it will use the production base URL.
func (s *PaymentService) GetMethods(ctx context.Context, req GetPaymentMethodRequest) (GetPaymentMethodResponse, *http.Response, error) {
	res := &GetPaymentMethodResponse{}
	path := "/merchant/paymentmethod/getpaymentmethod"

	baseUrl := common.SandboxV2BaseURL
	if s.client.Cfg.Environment == common.ProductionEnv {
		baseUrl = common.ProductionV2BaseURL
	}

	req.MerchantCode = s.client.Cfg.MerchantCode
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

// generatePaymentSignature generates an HMAC SHA-256 signature for the given
// parameter. It uses the API key from the config as the secret key to create
// the signature. The result is returned as a hexadecimal-encoded string. This
// signature is used for authenticating requests to the Duitku API.
func (s *PaymentService) generatePaymentSignature(parameter string) string {
	combinedStr := s.client.Cfg.MerchantCode + parameter + s.client.Cfg.APIKey
	hash := sha256.Sum256([]byte(combinedStr))
	return hex.EncodeToString(hash[:])
}
