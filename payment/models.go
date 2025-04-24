package payment

type GetPaymentMethodRequest struct {
	MerchantCode string `json:"merchantCode"`
	Amount       int    `json:"amount"`
	Datetime     string `json:"datetime"` // format: YYYY-MM-DD HH:mm:ss
	Signature    string `json:"signature"`
}

type GetPaymentMethodResponse struct {
	PaymentFee      []paymentFee `json:"paymentFee"`
	ResponseCode    string       `json:"responseCode"`
	ResponseMessage string       `json:"responseMessage"`
}

type paymentFee struct {
	PaymentMethod string `json:"paymentMethod"`
	PaymentName   string `json:"paymentName"`
	PaymentImage  string `json:"paymentImage"`
	TotalFee      string `json:"totalFee"`
}
