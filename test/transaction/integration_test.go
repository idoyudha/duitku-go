package transaction

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/idoyudha/duitku-go"
	"github.com/idoyudha/duitku-go/common"
	"github.com/idoyudha/duitku-go/transaction"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionServiceCC(t *testing.T) {
	godotenv.Load("./../../.env")

	client := duitku.NewClient(&common.Config{
		MerchantCode: os.Getenv("MERCHANT_CODE"),
		APIKey:       os.Getenv("API_KEY"),
		Environment:  common.SandboxEnv,
	})

	merchantOrderId := time.Now().Format("20060102150405")

	t.Run("CreateTransaction Credit/Debit Card", func(t *testing.T) {
		t.Run("Success Credit/Debit Card", func(t *testing.T) {
			trxReq := transaction.CreateTransactionRequest{
				PaymentAmount:   10001,
				MerchantOrderId: merchantOrderId,
				ProductDetails:  "test pay integration",
				Email:           "test@duitku.com",
				PaymentMethod:   common.CC,
				CustomerVaName:  "Test Transaction",
				CallbackURL:     "https://duitku.com/callback",
				ReturnURL:       "https://duitku.com",
			}
			res, httpResp, err := client.TransactionService.Create(context.Background(), trxReq)
			require.NotNil(t, res)
			require.NotNil(t, httpResp)
			require.Nil(t, err)
			assert.Equal(t, http.StatusOK, httpResp.StatusCode)
			respBody, readErr := io.ReadAll(httpResp.Body)
			require.NoError(t, readErr)
			assert.Contains(t, string(respBody), "merchantCode")
			assert.Contains(t, string(respBody), "reference")
			assert.Contains(t, string(respBody), "paymentUrl")
			assert.Contains(t, string(respBody), "statusCode")
			assert.Contains(t, string(respBody), "statusMessage")
		})

		t.Run("Get Transaction Status", func(t *testing.T) {
			trxStatusReq := transaction.GetTransactionStatusRequest{
				MerchantOrderId: merchantOrderId,
			}
			res, httpResp, err := client.TransactionService.GetStatus(context.Background(), trxStatusReq)
			require.NotNil(t, res)
			require.NotNil(t, httpResp)
			require.Nil(t, err)
			assert.Equal(t, http.StatusOK, httpResp.StatusCode)
			respBody, readErr := io.ReadAll(httpResp.Body)
			require.NoError(t, readErr)
			assert.Contains(t, string(respBody), "merchantOrderId")
			assert.Contains(t, string(respBody), "reference")
			assert.Contains(t, string(respBody), "amount")
			assert.Contains(t, string(respBody), "statusCode")
			assert.Contains(t, string(respBody), "statusMessage")
		})
	})
}

func TestTransactionServiceQR(t *testing.T) {
	godotenv.Load("./../../.env")

	client := duitku.NewClient(&common.Config{
		MerchantCode: os.Getenv("MERCHANT_CODE"),
		APIKey:       os.Getenv("API_KEY"),
		Environment:  common.SandboxEnv,
	})

	merchantOrderId := time.Now().Format("20060102150405")

	t.Run("CreateTransaction QR", func(t *testing.T) {
		t.Run("Success QR", func(t *testing.T) {
			trxReq := transaction.CreateTransactionRequest{
				MerchantCode:    client.Cfg.MerchantCode,
				PaymentAmount:   10001,
				MerchantOrderId: merchantOrderId,
				ProductDetails:  "test pay integration",
				Email:           "test@duitku.com",
				PaymentMethod:   common.SPAYQR,
				CustomerVaName:  "Test Transaction",
				CallbackURL:     "https://duitku.com/callback",
				ReturnURL:       "https://duitku.com",
			}
			res, httpResp, err := client.TransactionService.Create(context.Background(), trxReq)
			require.NotNil(t, res)
			require.NotNil(t, httpResp)
			require.Nil(t, err)
			assert.Equal(t, http.StatusOK, httpResp.StatusCode)
			respBody, readErr := io.ReadAll(httpResp.Body)
			require.NoError(t, readErr)
			assert.Contains(t, string(respBody), "merchantCode")
			assert.Contains(t, string(respBody), "reference")
			assert.Contains(t, string(respBody), "paymentUrl")
			assert.Contains(t, string(respBody), "amount")
			assert.Contains(t, string(respBody), "qrString")
		})

		t.Run("Get Transaction Status", func(t *testing.T) {
			trxStatusReq := transaction.GetTransactionStatusRequest{
				MerchantOrderId: merchantOrderId,
			}
			res, httpResp, err := client.TransactionService.GetStatus(context.Background(), trxStatusReq)
			require.NotNil(t, res)
			require.NotNil(t, httpResp)
			require.Nil(t, err)
			assert.Equal(t, http.StatusOK, httpResp.StatusCode)
			respBody, readErr := io.ReadAll(httpResp.Body)
			require.NoError(t, readErr)
			assert.Contains(t, string(respBody), "merchantOrderId")
			assert.Contains(t, string(respBody), "reference")
			assert.Contains(t, string(respBody), "amount")
			assert.Contains(t, string(respBody), "statusCode")
			assert.Contains(t, string(respBody), "statusMessage")
		})
	})
}

func TestTransactionServiceVA(t *testing.T) {
	godotenv.Load("./../../.env")

	client := duitku.NewClient(&common.Config{
		MerchantCode: os.Getenv("MERCHANT_CODE"),
		APIKey:       os.Getenv("API_KEY"),
		Environment:  common.SandboxEnv,
	})

	merchantOrderId := time.Now().Format("20060102150405")

	t.Run("CreateTransaction Virtual Account", func(t *testing.T) {
		t.Run("Success Virtual Account", func(t *testing.T) {
			trxReq := transaction.CreateTransactionRequest{
				MerchantCode:    client.Cfg.MerchantCode,
				PaymentAmount:   10001,
				MerchantOrderId: merchantOrderId,
				ProductDetails:  "test pay integration",
				Email:           "test@duitku.com",
				PaymentMethod:   common.BCAVA,
				CustomerVaName:  "Test Transaction",
				CallbackURL:     "https://duitku.com/callback",
				ReturnURL:       "https://duitku.com",
			}
			res, httpResp, err := client.TransactionService.Create(context.Background(), trxReq)
			require.NotNil(t, res)
			require.NotNil(t, httpResp)
			require.Nil(t, err)
			assert.Equal(t, http.StatusOK, httpResp.StatusCode)
			respBody, readErr := io.ReadAll(httpResp.Body)
			require.NoError(t, readErr)
			assert.Contains(t, string(respBody), "merchantCode")
			assert.Contains(t, string(respBody), "reference")
			assert.Contains(t, string(respBody), "paymentUrl")
			assert.Contains(t, string(respBody), "vaNumber")
			assert.Contains(t, string(respBody), "amount")
		})

		t.Run("Get Transaction Status", func(t *testing.T) {
			trxStatusReq := transaction.GetTransactionStatusRequest{
				MerchantOrderId: merchantOrderId,
			}
			res, httpResp, err := client.TransactionService.GetStatus(context.Background(), trxStatusReq)
			require.NotNil(t, res)
			require.NotNil(t, httpResp)
			require.Nil(t, err)
			assert.Equal(t, http.StatusOK, httpResp.StatusCode)
			respBody, readErr := io.ReadAll(httpResp.Body)
			require.NoError(t, readErr)
			assert.Contains(t, string(respBody), "merchantOrderId")
			assert.Contains(t, string(respBody), "reference")
			assert.Contains(t, string(respBody), "amount")
			assert.Contains(t, string(respBody), "statusCode")
			assert.Contains(t, string(respBody), "statusMessage")
		})
	})
}
