package main

import (
	"ale-abel/sdk-tbk-go/controller/oneclick"
	oneclick_model "ale-abel/sdk-tbk-go/model/oneclick"

	"github.com/gin-gonic/gin"
)

func main() {

}

func CreateInscription(username string, email string, responseURL string) (interface{}, error) {
	// origen := ctx.FullPath()
	init := oneclick_model.CreateInscriptionRequest{
		Username:    username,
		Email:       email,
		ResponseURL: responseURL,
	}
	return oneclick.CreateInscription(&init)
}

func ConfirmInscription(tbkToken string) (interface{}, error) {
	confirm := oneclick_model.ConfirmInscriptionRequest{
		Token: tbkToken,
	}
	return oneclick.ConfirmInscription(&confirm)
}
func Authorize(username string, tbkUser string, buyOrder string, commerceCode string, childBuyOrder string, amount int, installmentsNumber int, ctx *gin.Context) (interface{}, error) {
	authorize := oneclick_model.AuthorizeTransactionRequest{
		UserName: username,
		TbkUser:  tbkUser,
		BuyOrder: buyOrder,
		Details: []oneclick_model.AuthorizeTransactionDetail{
			{
				CommerceCode:       commerceCode,
				BuyOrder:           childBuyOrder,
				Amount:             amount,
				InstallmentsNumber: installmentsNumber,
			},
		},
	}
	return oneclick.Authorize(&authorize)
}
