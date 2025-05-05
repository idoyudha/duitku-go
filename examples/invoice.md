# Invoice Service

## Create Invoice

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/idoyudha/duitku-go"
	"github.com/idoyudha/duitku-go/common"
	"github.com/idoyudha/duitku-go/invoice"
)

func main() {
    client := duitku.NewClient(&common.Config{
		MerchantCode: "YOUR_MERCHANT_CODE",
		APIKey:       "YOUR_API_KEY",
		Environment:  common.SandboxEnv,
	})

    invoiceRequest := invoice.CreateInvoiceRequest{
		PaymentAmount:   10001,
		MerchantOrderId: time.Now().Format("20060102150405"),
		ProductDetails:  "test go library",
		Email:           "admin@yourcompany.com",
		CallbackURL:     "https://yourcompany.com/callback",
		ReturnURL:       "https://yourcompany.com",
	}
	createInvoiceRes, httpResponse, err := client.InvoiceService.Create(context.Background(), invoiceRequest)
	if err != nil {
		log.Printf("Found error InvoiceService.Create => %v", err)
	}

	log.Printf("Full HTTP Response from InvoiceService.Create => %v", httpResponse)
	log.Printf("Response body from InvoiceService.Create => %v", createInvoiceRes)
}
```