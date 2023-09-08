package oneclick

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"ale-abel/sdk-tbk-go/constant"
	oneclick_model "ale-abel/sdk-tbk-go/model/oneclick"
	"ale-abel/sdk-tbk-go/util"

	log "github.com/sirupsen/logrus"
)

func ConfirmInscription(request *oneclick_model.ConfirmInscriptionRequest) (interface{}, error) {
	post, _ := json.Marshal(&request)
	url := fmt.Sprintf(`%s/%s`, constant.Host, constant.PathOneClick+"inscriptions/"+request.Token)
	newReq, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(post))
	log.Infof("[REQUEST][TRANSBANK][ConfirmInscription] [%s]", string(post))
	responseService, err := util.InvokeAPI(newReq)
	if err != nil {
		log.Errorf("[ERROR HTTP][TRANSBANK][%s][ConfirmInscription][%s]", request.Token, err.Error())
		return nil, errors.New("error calling transbank init transaction api")
	}
	defer responseService.Body.Close()
	if responseService.StatusCode != http.StatusOK {
		log.Errorf("[ERROR TRANSBANK][TRANSBANK][%s][ConfirmInscription][%s]", request.Token, "hubo un error al procesar el servicio en TRANSBANK")
		return nil, errors.New("hubo un error al procesar el servicio en TRANSBANK")
	}
	var resp oneclick_model.CreateInscriptionResponse
	e := json.NewDecoder(responseService.Body).Decode(&resp)
	if e != nil {
		log.Errorf("[ERROR DECODING BODY TO TRANSBANK STRUCT][TRANSBANK][%s][CreateInscription][%s]", request.Token, e.Error())
		return nil, errors.New("error decodificando body")
	}

	log.Infof("[RESPONSE] CreateInscription [%s]", resp)
	return resp, nil
}
