package payment

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/idoyudha/duitku-go"
	"github.com/idoyudha/duitku-go/common"
	"github.com/idoyudha/duitku-go/payment"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPaymentMethods(t *testing.T) {
	godotenv.Load("./../../.env")

	client := duitku.NewClient(&common.Config{
		MerchantCode: os.Getenv("MERCHANT_CODE"),
		APIKey:       os.Getenv("API_KEY"),
		Environment:  common.SandboxEnv,
	})

	t.Run("Get Payment Methods", func(t *testing.T) {
		t.Run("Bad Request", func(t *testing.T) {
			getPaymentMethodReq := payment.GetPaymentMethodRequest{}
			res, httpResp, err := client.PaymentService.GetMethods(context.Background(), getPaymentMethodReq)
			require.NotNil(t, res)
			require.NotNil(t, httpResp)
			require.Nil(t, err)
			assert.Equal(t, http.StatusBadRequest, httpResp.StatusCode)
			respBody, readErr := io.ReadAll(httpResp.Body)
			require.NoError(t, readErr)
			assert.Contains(t, string(respBody), "Message")
		})

		t.Run("Success", func(t *testing.T) {
			getPaymentMethodReq := payment.GetPaymentMethodRequest{
				MerchantCode: client.Cfg.MerchantCode,
				Amount:       10001,
				Datetime:     time.Now().Format("2006-01-02 15:04:05"),
			}
			res, httpResp, err := client.PaymentService.GetMethods(context.Background(), getPaymentMethodReq)
			require.NotNil(t, res)
			require.NotNil(t, httpResp)
			require.Nil(t, err)
			assert.Equal(t, http.StatusOK, httpResp.StatusCode)
			respBody, readErr := io.ReadAll(httpResp.Body)
			require.NoError(t, readErr)
			assert.Contains(t, string(respBody), "paymentFee")
			assert.Contains(t, string(respBody), "paymentMethod")
			assert.Contains(t, string(respBody), "paymentName")
			assert.Contains(t, string(respBody), "paymentImage")
			assert.Contains(t, string(respBody), "totalFee")
			assert.Contains(t, string(respBody), "responseCode")
			assert.Contains(t, string(respBody), "responseMessage")
		})
	})
}
