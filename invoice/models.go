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

type ovoDetail struct {
	PaymentDetails []ovoPaymentDetail `json:"paymentDetails"`
}

type ovoPaymentDetail struct {
	PaymentType string `json:"paymentType"`
	Amount      int    `json:"amount"`
}

type shopeeDetail struct {
	PromoIds string `json:"promo_ids"`
	UseCoin  bool   `json:"useCoin"`
}

type accountLink struct {
	CredentialCode string       `json:"credentialCode"`
	Ovo            ovoDetail    `json:"ovo"`
	Shopee         shopeeDetail `json:"shopee"`
}

type CreateInvoiceRequest struct {
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
	AccountLink      accountLink          `json:"accountLink,omitempty"`
}

type CreateInvoiceResponse struct {
	MerchantCode  string `json:"merchantCode"`
	Reference     string `json:"reference"`
	PaymentURL    string `json:"paymentUrl"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}
