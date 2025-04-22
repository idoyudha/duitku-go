package invoice

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/idoyudha/duitku-go/common"
)

type InvoiceService struct {
	client *common.ServiceClient
}

// NewInvoiceService returns a new instance of InvoiceService with the given service client.
func NewInvoiceService(service *common.ServiceClient) *InvoiceService {
	return &InvoiceService{
		client: service,
	}
}

// Create creates a new invoice on duitku's server.
// The method returns CreateInvoiceResponse, HTTP response, and error.
// The CreateInvoiceResponse contains the detail of the created invoice.
//
// The request body should match the CreateInvoiceRequest struct.
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
// The method will add the following headers to the request:
//   - X-Duitku-Merchantcode: the merchant code from the config.
//   - X-Duitku-Timestamp: the current timestamp in milliseconds.
//   - X-Duitku-Signature: the signature of the request using the API key from
//     the config and the timestamp as the input.
//
// The method will use the sandbox base URL if the environment in the config
// is set to sandbox. Otherwise, it will use the production base URL.
func (s *InvoiceService) Create(ctx context.Context, req CreateInvoiceRequest) (CreateInvoiceResponse, *http.Response, error) {
	res := &CreateInvoiceResponse{}
	path := "/merchant/createInvoice"
	headerParams := make(map[string]string)
	headerParams[common.MerchantCodeHeader] = s.client.Cfg.MerchantCode
	headerParams[common.TimeStampHeader] = s.client.GetCurrentTimestamp()
	headerParams[common.SignatureHeader] = s.GenerateInvoiceSignature(headerParams[common.TimeStampHeader])

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

// GenerateInvoiceSignature generates an HMAC SHA-256 signature for the given timestamp.
// It uses the API key from the config as the secret key to create the signature.
// The result is returned as a hexadecimal-encoded string. This signature is used
// for authenticating requests to the Duitku POP API.
func (s *InvoiceService) GenerateInvoiceSignature(timestamp string) string {
	h := hmac.New(sha256.New, []byte(s.client.Cfg.APIKey))
	h.Write([]byte(s.client.Cfg.MerchantCode + timestamp))

	return hex.EncodeToString(h.Sum(nil))
}
