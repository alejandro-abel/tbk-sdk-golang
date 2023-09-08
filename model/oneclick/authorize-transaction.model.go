package oneclick_model

type AuthorizeTransactionRequest struct {
	UserName string                       `json:"username"`
	TbkUser  string                       `json:"tbk_user"`
	BuyOrder string                       `json:"buy_order"`
	Details  []AuthorizeTransactionDetail `json:"details"`
}

type AuthorizeTransactionDetail struct {
	CommerceCode       string `json:"commerceCode"`
	BuyOrder           string `json:"buy_order"`
	Amount             int    `json:"amount"`
	InstallmentsNumber int    `json:"installments_number"`
}

type AuthorizeResopnse struct {
	BuyOrder        string                              `json:"buy_order"`
	CardDetail      CardDetail                          `json:"card_detail"`
	AccountingDate  string                              `json:"accounting_date"`
	TransactionDate string                              `json:"transactio_date"`
	Details         AuthorizeTransactionDetailsResponse `json:"details"`
}

type CardDetail struct {
	CardNumber string `json:"card_number"`
}

type AuthorizeTransactionDetailsResponse struct {
	Amount             int    `json:"amount"`
	Status             string `json:"status"`
	AuthorizationCode  string `json:"authorization_code"`
	PaymentTypeCode    string `json:"payment_type_code"`
	ResponseCode       int    `json:"response_code"`
	InstallmentsNumber int    `json:"installments_number"`
	CommerceCode       int    `json:"commerce_code"`
	BuyOrder           string `json:"buy_order"`
}
