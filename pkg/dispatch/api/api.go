package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

//API dispatch will dispatch apigateway message
type API struct {
	urlPath string
	client  *http.Client
}

//New return API dispatch instance
func New(urlPath string, client *http.Client) *API {
	return &API{
		urlPath: urlPath,
		client:  client,
	}
}

//Process is main logic
func (a *API) Process(logger *zap.Logger, bodys []byte) []byte {
	defer logger.Sync()
	apigw := gjson.ParseBytes(bodys)
	reqpath := apigw.Get("path").String()
	reqmethod := apigw.Get("httpMethod").String()
	sourceIP := apigw.Get("requestContext.sourceIp").String()
	reqid := apigw.Get("requestContext.requestId").String()
	reqbody := apigw.Get("body").String()
	logger.Debug(reqbody)
	logger.Debug("we get the request from scf and deal it with api dispatch", zap.String("requestId", reqid))
	req, err := http.NewRequest(reqmethod, fmt.Sprintf("%s%s", a.urlPath, reqpath), strings.NewReader(reqbody))
	if err != nil {
		logger.Panic("api dispatch create request failed", zap.Errors("error", []error{err}))
	}
	req.Header.Add("X-Forwarded-For", sourceIP)
	//proxy headers
	for key, header := range apigw.Get("headers").Map() {
		req.Header.Add(key, header.String())
	}
	//proxy query
	q := req.URL.Query()
	for key, query := range apigw.Get("queryString").Map() {
		q.Add(key, query.String())
	}
	req.URL.RawQuery = q.Encode()

	resp, err := a.client.Do(req)
	if err != nil {
		logger.Panic("api dispatch create request failed", zap.Errors("error", []error{err}))
	}
	defer resp.Body.Close()
	respbody, err := ioutil.ReadAll(resp.Body)
	return respbody
}
