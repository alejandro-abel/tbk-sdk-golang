package oneclick

import (
	oneclick_model "ale-abel/sdk-tbk-go/model/oneclick"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"ale-abel/sdk-tbk-go/constant"

	"ale-abel/sdk-tbk-go/util"

	log "github.com/sirupsen/logrus"
)

func Authorize(request *oneclick_model.AuthorizeTransactionRequest) (interface{}, error) {
	post, _ := json.Marshal(&request)
	url := fmt.Sprintf(`%s/%s`, constant.Host, constant.PathOneClick+"transactions")
	newReq, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(post))
	log.Infof("[REQUEST][TRANSBANK][Authorize][%s]", string(post))
	responseService, err := util.InvokeAPI(newReq)
	if err != nil {
		log.Errorf("[ERROR HTTP][TRANSBANK][%s][Authorize][%s]", request.TbkUser, err.Error())
		return nil, errors.New("error calling transbank init transaction api")
	}
	defer responseService.Body.Close()
	if responseService.StatusCode != http.StatusOK {
		log.Errorf("[ERROR TRANSBANK][TRANSBANK][%s][Authorize][%s]", request.TbkUser, "hubo un error al procesar el servicio en TRANSBANK")
		return nil, errors.New("hubo un error al procesar el servicio en TRANSBANK")
	}
	var resp oneclick_model.AuthorizeResopnse
	e := json.NewDecoder(responseService.Body).Decode(&resp)
	if e != nil {
		log.Errorf("[ERROR DECODING BODY TO TRANSBANK STRUCT][TRANSBANK][%s][Authorize][%s]", request.TbkUser, e.Error())
		return nil, errors.New("error decodificando body")
	}
	log.Infof("[RESPONSE] Authorize [%v]", resp)
	return resp, nil
}
