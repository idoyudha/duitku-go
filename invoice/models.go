package invoice

import "github.com/idoyudha/duitku-go/common"

type itemDetails struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type customerDetail struct {
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

type creditCardDetail struct {
	Acquirer     string   `json:"acquirer"`
	BinWhitelist []string `json:"binWhitelist"`
}

type CreateInvoiceRequest struct {
	PaymentAmount    int                  `json:"paymentAmount"`
	MerchantOrderId  string               `json:"merchantOrderId"`
	ProductDetails   string               `json:"productDetails"`
	Email            string               `json:"email"`
	AdditionalParam  string               `json:"additionalParam,omitempty"`
	MerchantUserInfo string               `json:"merchantUserInfo,omitempty"`
	CustomerVaName   string               `json:"customerVaName"`
	PhoneNumber      string               `json:"phoneNumber,omitempty"`
	ItemDetails      []itemDetails        `json:"itemDetails,omitempty"`
	CustomerDetail   customerDetail       `json:"customerDetail,omitempty"`
	CallbackURL      string               `json:"callbackUrl"`
	ReturnURL        string               `json:"returnUrl"`
	ExpiryPeriod     int                  `json:"expiryPeriod,omitempty"`
	PaymentMethod    common.PaymentMethod `json:"paymentMethod"`
	CreditCardDetail creditCardDetail     `json:"creditCardDetail"`
}

type CreateInvoiceResponse struct {
	MerchantCode  string `json:"merchantCode"`
	Reference     string `json:"reference"`
	PaymentURL    string `json:"paymentUrl"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}
