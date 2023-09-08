package oneclick_model

type CreateInscriptionRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	ResponseURL string `json:"response_url"`
}

type CreateInscriptionResponse struct {
	Token     string `json:"token"`
	URLWebPay string `json:"url_webpay"`
}
