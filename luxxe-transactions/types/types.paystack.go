package types

type PaystackAuthorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpiryMonth       string `json:"exp_month"`
	ExpiryYear        string `json:"exp_year"`
	Channel           string `json:"channel"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	Reusable          bool   `json:"reusable"`
	Signature         string `json:"signature"`
	AccountName       string `json:"account_name"`
}

type PaystackCustomer struct {
	ID                       int         `json:"id"`
	FirstName                string      `json:"first_name"`
	LastName                 string      `json:"last_name"`
	Email                    string      `json:"email"`
	CustomerCode             string      `json:"customer_code"`
	Phone                    string      `json:"phone"`
	MetaData                 interface{} `json:"metadata"`
	RiskAction               string      `json:"risk_action"`
	InternationalFormatPhone string      `json:"international_format_phone"`
}

type PaystackLogHistory struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

type PaystackLog struct {
	StartTime int                `json:"start_time"`
	TimeSpent int16              `json:"time_spent"`
	Attempts  int16              `json:"attempts"`
	Errors    int16              `json:"errors"`
	Success   bool               `json:"success"`
	Mobile    bool               `json:"mobile"`
	Input     interface{}        `json:"input"`
	History   PaystackLogHistory `json:"history"`
}

type PaystackStandardVerificationResponseData struct {
	ID                 int                       `json:"id"`
	Domain             string                    `json:"domain"`
	Status             TransactionStatusPaystack `json:"status"`
	Reference          string                    `json:"reference"`
	Amount             int                       `json:"amount"`
	Message            string                    `json:"message"`
	GatewayResponse    string                    `json:"gateway_response"`
	Channel            string                    `json:"channel"`
	Currency           string                    `json:"currency"`
	IpAddress          string                    `json:"ip_address"`
	Metadata           interface{}               `json:"metadata"`
	Log                interface{}               `json:"log"`
	Fees               interface{}               `json:"fees"`
	FeesSplit          interface{}               `json:"fees_split"`
	Authorization      interface{}               `json:"authorization"`
	Customer           interface{}               `json:"customer"`
	Plan               interface{}               `json:"plan"`
	Splut              interface{}               `json:"split"`
	OrderId            interface{}               `json:"order_id"`
	PaidAt             string                    `json:"paidAt"`
	CreatedAt          string                    `json:"created_at"`
	RequestedAmount    int                       `json:"requested_amount"`
	PosTransactionData interface{}               `json:"pos_transaction_data"`
	Source             interface{}               `json:"source"`
	FeesBreakdown      interface{}               `json:"fees_breakdown"`
	TransactionDate    string                    `json:"transaction_date"`
	PlanObject         interface{}               `json:"plan_object"`
	Subaccount         interface{}               `json:"subaccount"`
}

type PaystackCreateSubaccountResponseData struct {
	BusinessName         string  `json:"business_name"`
	AccountNumber        string  `json:"account_number"`
	PercentageCharge     float32 `json:"percentage_charge"`
	SettlementBank       string  `json:"settlement_bank"`
	Currency             string  `json:"currency"`
	Bank                 int     `json:"bank"`
	Integration          int     `json:"integration"`
	Domain               string  `json:"domain"`
	AccountName          string  `json:"account_name"`
	Product              string  `json:"product"`
	ManagedByIntegration int     `json:"managed_by_integration"`
	SubaccountCode       string  `json:"subaccount_code"`
	IsVerified           bool    `json:"is_verified"`
	SettlementSchedule   string  `json:"settlement_schedule"`
	Active               bool    `json:"active"`
	Migrate              bool    `json:"migrate"`
	Id                   int     `json:"id"`
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
}

type PaystackInitializeTransactionResponseData struct {
	AuthorizationUrl string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type ResolveAccountNumberData struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankId        int    `json:"bank_id"`
}

type PaystackResponse[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type PaystackStandardWebhookVerificationResponse struct {
	Event WebHookEventPaystack                     `json:"event"`
	Data  PaystackStandardVerificationResponseData `json:"data"`
}

type WebHookEventPaystack string

const WebHookEventPaystackSuccessfulCharge WebHookEventPaystack = "charge.success"

type TransactionStatusPaystack string

const (
	TransactionStatusPaystackSuccess TransactionStatusPaystack = "success"
	TransactionStatusPaystackPending TransactionStatusPaystack = "pending"
	TransactionStatusPaystackFailed  TransactionStatusPaystack = "failed"
)

type MessagePaystack string

const (
	MessagePaystackVerificationSuccessful  MessagePaystack = "Verification successful"
	MessagePaystackWebhookEventSuccessful  MessagePaystack = "charge.success"
	MessagePaystackInvalidKey              MessagePaystack = "Invalid key"
	MessagePaystackAuthorizationURLCreated MessagePaystack = "Authorization URL created"
)
