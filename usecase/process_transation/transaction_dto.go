package processtransation

type TransactionDTOInput struct {
	ID                         string  `json:"id"`
	AccountID                  string  `json:"acount_id"`
	CreditCardNumber           string  `json:"credit_card_number"`
	CreditCardName             string  `json:"credit_card_name"`
	CreditCardExpirationMonth  int     `json:"credit_card_expiration_month"`
	CreditCardExpirationYear   int     `json:"credit_card_expiration_year"`
	CreditCardVerificationCode int     `json:"credit_card_verification_code"`
	CreditCardAmount           float64 `json:"credit_card_amount"`
}

type TransactionDTOOutput struct {
	ID           string `json:id`
	Status       string `json:status`
	ErrorMessage string `json:error_message`
}
