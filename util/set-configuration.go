package util

import (
	"ale-abel/sdk-tbk-go/constant"

	log "github.com/sirupsen/logrus"
)

func SetIntegrationConfiguration() {
	constant.Host = constant.HostIntegration
	constant.MallCommerceCode = constant.MallCommerceCodeIntegration
	constant.SecretKey = constant.SecretKeyIntegration
}

func SetProductionConfiguration(mallComerceCode string, secretkey string) {
	constant.Host = constant.HostProduction
	constant.MallCommerceCode = mallComerceCode
	constant.SecretKey = secretkey
}

func ValidateConfiguration() bool {
	if constant.Host == "" {
		log.Info("=============================")
		log.Info("NO CONFIGURATION SET")
		log.Info("=============================")
		return false
	}
	return true
}
