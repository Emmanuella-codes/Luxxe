package types

type PaymentClientFetchSingleResponse struct {
	Paystack    PaystackStandardVerificationResponseData 
}

type PaymentClientFetchAllResponse struct {
	Paystack    []PaystackStandardVerificationResponseData
}
