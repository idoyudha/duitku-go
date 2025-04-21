package invoice

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/idoyudha/duitku-go"
	"github.com/idoyudha/duitku-go/common"
	"github.com/idoyudha/duitku-go/invoice"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvoiceService(t *testing.T) {
	godotenv.Load("./../../.env")

	client := duitku.NewClient(&common.Config{
		MerchantCode: os.Getenv("MERCHANT_CODE"),
		APIKey:       os.Getenv("API_KEY"),
		Environment:  common.SandboxEnv,
	})

	t.Run("CreateInvoice", func(t *testing.T) {
		t.Run("Failed Bad Request", func(t *testing.T) {
			invReq := invoice.CreateInvoiceRequest{}
			res, httpResp, err := client.InvoiceService.Create(context.Background(), invReq)
			require.NotNil(t, res)
			require.NotNil(t, httpResp)
			require.Nil(t, err)
			assert.Equal(t, http.StatusBadRequest, httpResp.StatusCode)
		})
		t.Run("Success", func(t *testing.T) {
			invReq := invoice.CreateInvoiceRequest{
				PaymentAmount:   10001,
				MerchantOrderId: time.Now().Format("20060102150405"),
				ProductDetails:  "test pay integration",
				Email:           "test@duitku.com",
				CallbackURL:     "https://duitku.com/callback",
				ReturnURL:       "https://duitku.com",
			}
			res, httpResp, err := client.InvoiceService.Create(context.Background(), invReq)
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
			assert.Contains(t, string(respBody), "statusCode")
			assert.Contains(t, string(respBody), "statusMessage")
		})
	})
}
