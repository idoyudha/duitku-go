# Duitku Golang API Client Library
[![Go Reference](https://pkg.go.dev/badge/github.com/idoyudha/duitku-go)](https://pkg.go.dev/badge/github.com/idoyudha/duitku-go)
[![Build Status](https://github.com/idoyudha/duitku-go/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/idoyudha/duitku-go/actions/workflows/go.yml?query=branch%3Amaster)

Duitku API Library for Go

## Supported Feature
|        Feature         |              Function                |                HTTP Request                   |              Description              |
|------------------------|--------------------------------------|-----------------------------------------------|---------------------------------------|
| Get Payment Method     | client.PaymentService.GetMethods     | POST /merchant/paymentmethod/getpaymentmethod | Get list of available payment methods |
| Create New Transaction | client.TransactionService.Create     | POST /merchant/v2/inquiry                     | Create Transaction via V2 API         |
| Get Transaction        | client.TransactionService.GetStatus  | POST /merchant/transactionStatus              | Get Transaction via V2 API            |
| Craete New Invoice     | client.InvoiceService.Create         | POST /merchant/createInvoice                  | Create Transaction via POP API        |

## Requirements
- Go 1.24 or later
- Duitku account, [register here](https://dashboard.duitku.com/Account/Register)
- [API Key](https://docs.duitku.com/en/account/#account-integration--getting-api-key)

## Documentation
- https://docs.duitku.com/

## Installation

Get this library, add to your project

```bash
go get https://github.com/idoyudha/duitku-go
```

## Usage
```go
import (
	"context"

	"github.com/idoyudha/duitku-go"
	"github.com/idoyudha/duitku-go/common"
	"github.com/idoyudha/duitku-go/invoice"
)

client := duitku.NewClient(&common.Config{
	MerchantCode: "YOUR MERCHANT CODE",
	APIKey:       "YOUR API KEY",
	Environment:  common.SandboxEnv,
})

invoiceRequest := invoice.CreateInvoiceRequest{
    PaymentAmount:   10001,
	MerchantOrderId: "YOUR UNIQUE ORDER ID",
	ProductDetails:  "YOUR PRODUCT DETAILS",
	Email:           "admin@yourcompany.com",
	CallbackURL:     "https://yourcompany.com/callback",
	ReturnURL:       "https://yourcompany.com",
}

res, httpResponse, err := client.InvoiceService.Create(context.Background(), invoiceRequest)
```

## Support
If you have a feature request or spotted a bug or a techical problem, [create an issue here](https://github.com/idoyudha/duitku-go/issues/new/choose).
For other questions, please contact duitku through their live chat on your dashboard.

## License
MIT license. For more information, see the LICENSE file.

