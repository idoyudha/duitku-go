# Payment Service

## Get Payment Method

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/idoyudha/duitku-go"
	"github.com/idoyudha/duitku-go/common"
	"github.com/idoyudha/duitku-go/payment"
)

func main() {
    client := duitku.NewClient(&common.Config{
		MerchantCode: "YOUR_MERCHANT_CODE",
		APIKey:       "YOUR_API_KEY",
		Environment:  common.SandboxEnv,
	})

    getPaymentMethodRequest := payment.GetPaymentMethodRequest{
		MerchantCode: client.Cfg.MerchantCode,
		Amount:       10000,
		Datetime:     time.Now().Format("2006-01-02 15:04:05"),
	}
	paymentMethodRes, httpResponse, err := client.PaymentService.GetMethods(context.Background(), getPaymentMethodRequest)
	if err != nil {
		log.Println("PaymentService.GetMethods found error => ", err)
	}

	log.Println("PaymentService.GetMethods httpResponse => ", httpResponse)
	log.Println("PaymentService.GetMethods paymentMethodRes => ", paymentMethodRes)
}
```