# Payment Service

## Create Transaction

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/idoyudha/duitku-go"
	"github.com/idoyudha/duitku-go/common"
	"github.com/idoyudha/duitku-go/transaction"
)

func main() {
    client := duitku.NewClient(&common.Config{
		MerchantCode: "YOUR_MERCHANT_CODE",
		APIKey:       "YOUR_API_KEY",
		Environment:  common.SandboxEnv,
	})

    trxRequest := transaction.CreateTransactionRequest{
		MerchantCode:    client.Cfg.MerchantCode,
		PaymentAmount:   10001,
		MerchantOrderId: time.Now().Format("20060102150405"),
		ProductDetails:  "test pay integration",
		Email:           "test@duitku.com",
		PaymentMethod:   common.CC,
		CustomerVaName:  "Test Transaction",
		CallbackURL:     "https://duitku.com/callback",
		ReturnURL:       "https://duitku.com",
	}
	trxRes, httpResponse, err := client.TransactionService.Create(context.Background(), trxRequest)
	if err != nil {
		log.Println("TransactionService.Create found error => ", err)
	}

	log.Println("TransactionService.Create httpResponse => ", httpResponse)
	log.Println("TransactionService.Create trxRes => ", trxRes)
}
```

## Get Transaction Status

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

    trxStatusReq := transaction.GetTransactionStatusRequest{
		MerchantCode:    client.Cfg.MerchantCode,
		MerchantOrderId: time.Now().Format("20060102150405"),
	}
	trxStatusRes, httpResponse, err := client.TransactionService.GetStatus(context.Background(), trxStatusReq)
	if err != nil {
		log.Println("TransactionService.GetStatus found error => ", err)
	}

	log.Println("TransactionService.GetStatus httpResponse => ", httpResponse)
	log.Println("TransactionService.GetStatus trxStatusRes => ", trxStatusRes)
}
```