package plain

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

//Plain dispatch will dispatch raw message
type Plain struct {
	urlPath string
	client  *http.Client
}

//New return Plain dispatch instance
func New(urlPath string, client *http.Client) *Plain {
	return &Plain{
		urlPath: urlPath,
		client:  client,
	}
}

//Process is main logic
func (p *Plain) Process(logger *zap.Logger, bodys []byte) []byte {
	defer logger.Sync()
	logger.Debug("we get the request from scf and deal it with plain dispatch")
	req, err := http.NewRequest("POST", p.urlPath, bytes.NewReader(bodys))
	if err != nil {
		logger.Panic("api dispatch create request failed", zap.Errors("error", []error{err}))
	}
	resp, err := p.client.Do(req)
	if err != nil {
		logger.Panic("api dispatch create request failed", zap.Errors("error", []error{err}))
	}
	defer resp.Body.Close()
	respbody, err := ioutil.ReadAll(resp.Body)
	return respbody
}
