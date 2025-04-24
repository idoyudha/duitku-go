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

func (s *TransactionService) GetStatus(ctx context.Context, req GetTransactionStatusRequest) (GetTransactionStatusResponse, *http.Response, error) {
	res := &GetTransactionStatusResponse{}
	path := "/merchant/transactionStatus"
	headerParams := make(map[string]string)

	baseUrl := common.SandboxV2BaseURL
	if s.client.Cfg.Environment == common.ProductionEnv {
		baseUrl = common.ProductionV2BaseURL
	}

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

func (s *TransactionService) generateTransactionSignature(parameter string) string {
	combinedStr := s.client.Cfg.MerchantCode + parameter + s.client.Cfg.APIKey
	hash := md5.Sum([]byte(combinedStr))
	return hex.EncodeToString(hash[:])
}
