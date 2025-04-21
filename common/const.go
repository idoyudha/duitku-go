package common

type Environment string

const (
	SandboxEnv        Environment = "SANDBOX"
	SandboxBaseURL    string      = "https://sandbox.duitku.com/webapi/api"
	ProductionEnv     Environment = "PRODUCTION"
	ProductionBaseURL string      = "https://passport.duitku.com/webapi/api"
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
