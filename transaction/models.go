package transaction

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

type ovoPaymentDetails struct {
	PaymentType string `json:"paymentType"`
	Amount      int    `json:"amount"`
}

type ovo struct {
	PaymentDetails []ovoPaymentDetails `json:"paymentDetails"`
}

type shopee struct {
	PromoIDs string `json:"promo_ids"`
	UseCoin  string `json:"useCoin"`
}

type accountLink struct {
	CredentialCode string `json:"credentialCode"`
	Ovo            ovo    `json:"ovo"`
	Shopee         shopee `json:"shopee"`
}

type creditCardDetail struct {
	Acquirer     string   `json:"acquirer"`
	BinWhitelist []string `json:"binWhitelist"`
}

type CreateTransactionRequest struct {
	MerchantCode     string               `json:"merchantCode"`
	PaymentAmount    int                  `json:"paymentAmount"`
	MerchantOrderId  string               `json:"merchantOrderId"`
	ProductDetails   string               `json:"productDetails"`
	Email            string               `json:"email"`
	AdditionalParam  string               `json:"additionalParam,omitempty"`
	PaymentMethod    common.PaymentMethod `json:"paymentMethod"`
	MerchantUserInfo string               `json:"merchantUserInfo,omitempty"`
	CustomerVaName   string               `json:"customerVaName"`
	PhoneNumber      string               `json:"phoneNumber,omitempty"`
	ItemDetails      []itemDetails        `json:"itemDetails,omitempty"`
	CustomerDetail   customerDetail       `json:"customerDetail,omitempty"`
	ReturnURL        string               `json:"returnUrl"`
	CallbackURL      string               `json:"callbackUrl"`
	Signature        string               `json:"signature"`
	ExpiryPeriod     int                  `json:"expiryPeriod,omitempty"`
	AccountLink      accountLink          `json:"accountLink"`
	CreditCardDetail creditCardDetail     `json:"creditCardDetail"`
}

type CreateTransactionResponse struct {
	MerchantCode string `json:"merchantCode"`
	Reference    string `json:"reference"`
	PaymentURL   string `json:"paymentUrl"`
	VaNumber     string `json:"vaNumber"`
	Amount       string `json:"amount"`
	QRString     string `json:"qrString"`
}

type GetTransactionStatusRequest struct {
	MerchantCode    string `json:"merchantCode"`
	MerchantOrderId string `json:"merchantOrderId"`
	Signature       string `json:"signature"`
}

type GetTransactionStatusResponse struct {
	MerchantOrderId string `json:"merchantOrderId"`
	Reference       string `json:"reference"`
	Amount          string `json:"amount"`
	Fee             string `json:"fee"`
	StatusCode      string `json:"statusCode"`
	StatusMessage   string `json:"statusMessage"`
}
