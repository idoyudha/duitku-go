package common

type Environment string

const (
	SandboxEnv           Environment = "SANDBOX"
	ProductionEnv        Environment = "PRODUCTION"
	SandboxV2BaseURL     string      = "https://sandbox.duitku.com/webapi/api"
	ProductionV2BaseURL  string      = "https://passport.duitku.com/webapi/api"
	SandboxPOPBaseURL    string      = "https://api-sandbox.duitku.com/api"
	ProductionPOPBaseURL string      = "https://api-prod.duitku.com/api"
)

type PaymentMethod string

const (
	// Credit or Debit Card
	CC PaymentMethod = "VC"
	// Virtual Account
	BCAVA     PaymentMethod = "BC"
	MANDIRIVA PaymentMethod = "M2"
	BNIVA     PaymentMethod = "I1"
	// Ritel
	FT        PaymentMethod = "FT"
	Indomaret PaymentMethod = "IM"
	// E-wallet
	OVO         PaymentMethod = "OV"
	SPAYAPP     PaymentMethod = "SP"
	DANAEWALLET PaymentMethod = "DA"
	// Qris
	SPAYQR    PaymentMethod = "SP"
	DANAQR    PaymentMethod = "DQ"
	NUSAPAYQR PaymentMethod = "SQ"
	// Paylater
	INDODANA PaymentMethod = "DN"
	ATOME    PaymentMethod = "AT"
)

const (
	SignatureHeader    string = "x-duitku-signature"
	TimeStampHeader    string = "x-duitku-timestamp"
	MerchantCodeHeader string = "x-duitku-merchantcode"
)
