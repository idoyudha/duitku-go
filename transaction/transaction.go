package transaction

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/idoyudha/duitku-go/common"
)

type TransactionService struct {
	client *common.ServiceClient
}

// NewTransactionService returns a new instance of TransactionService with the given service client.
func NewTransactionService(service *common.ServiceClient) *TransactionService {
	return &TransactionService{
		client: service,
	}
}

func (s *TransactionService) Create(ctx context.Context, req CreateTransactionRequest) (CreateTransactionResponse, *http.Response, error) {
	res := &CreateTransactionResponse{}
	path := "/merchant/v2/inquiry"
	headerParams := make(map[string]string)

	baseUrl := common.SandboxV2BaseURL
	if s.client.Cfg.Environment == common.ProductionEnv {
		baseUrl = common.ProductionV2BaseURL
	}

	req.MerchantCode = s.client.Cfg.MerchantCode
	req.Signature = s.generateTransactionSignature(req.MerchantOrderId + strconv.Itoa(req.PaymentAmount))

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

// GetStatus gets the status of a transaction on duitku's server.
// The method returns GetTransactionStatusResponse, HTTP response, and error.
// The GetTransactionStatusResponse contains the status of the transaction.
//
// The request body should match the GetTransactionStatusRequest struct.
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
func (s *TransactionService) GetStatus(ctx context.Context, req GetTransactionStatusRequest) (GetTransactionStatusResponse, *http.Response, error) {
	res := &GetTransactionStatusResponse{}
	path := "/merchant/transactionStatus"
	headerParams := make(map[string]string)

	baseUrl := common.SandboxV2BaseURL
	if s.client.Cfg.Environment == common.ProductionEnv {
		baseUrl = common.ProductionV2BaseURL
	}

	req.MerchantCode = s.client.Cfg.MerchantCode
	req.Signature = s.generateTransactionSignature(req.MerchantOrderId)

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

// generateTransactionSignature generates an MD5 signature for the given
// parameter. It uses the merchant code and API key from the config to
// create a combined string, which is then hashed. The result is returned
// as a hexadecimal-encoded string. This signature is used for authenticating
// transaction requests to the Duitku API.
func (s *TransactionService) generateTransactionSignature(parameter string) string {
	combinedStr := s.client.Cfg.MerchantCode + parameter + s.client.Cfg.APIKey
	hash := md5.Sum([]byte(combinedStr))
	return hex.EncodeToString(hash[:])
}
