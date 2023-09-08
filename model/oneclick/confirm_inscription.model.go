package oneclick_model

type ConfirmInscriptionRequest struct {
	Token string `json:"token"`
}

type ConfirmInscriptionResponse struct {
	ResponseCode      string `json:"response_code"`
	TbkUser           string `json:"tbk_user"`
	AuthorizationCode string `json:"authorization_code"`
	CardType          string `json:"card_type"`
	CardNumber        string `json:"card_number"`
}
