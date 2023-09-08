package util

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"

	"ale-abel/sdk-tbk-go/constant"

	log "github.com/sirupsen/logrus"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

func InvokeAPI(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Tbk-Api-Key-Id", constant.MallCommerceCode)
	req.Header.Set("Tbk-Api-Key-Secret", constant.SecretKey)
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}
	log.Infof("[TRANSBANK][HEADERS]")
	for _, h := range req.Header {
		for _, hh := range h {
			log.Infof("[HEADER][%s]", hh)
		}
	}

	httpclient := &http.Client{Transport: tr}
	httpclient.Timeout = time.Minute * 2
	httpclient = httptrace.WrapClient(httpclient)
	resService, err := httpclient.Do(req.WithContext(req.Context()))
	defer httpclient.CloseIdleConnections()
	if err != nil {
		return nil, err
	}
	if resService.StatusCode != http.StatusOK {
		rawBody, err := ioutil.ReadAll(resService.Body)
		if err != nil {
			resService.Body.Close()
			log.Errorf("[WP-ERROR] error reading response %s", err.Error())
			return nil, err
		}
		log.Infof("[ERROR STATUS CODE][%d][%s][&s]", resService.StatusCode, req.URL, string(rawBody))
		resService.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
	} else {
		log.Infof("[OK][%d][%s]", resService.StatusCode, req.URL)
	}
	return resService, nil
}
