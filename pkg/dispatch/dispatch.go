package dispatch

import (
	"net/http"

	"github.com/TencentCloudMiddleWare/ScfWatchDog/pkg/dispatch/api"
	"go.uber.org/zap"
)

const (
	//API mode
	API = "api"
	//Plain mode
	Plain = "plain"
)

//Dispatcher is interface for process scf message
type Dispatcher interface {
	Process(logger *zap.Logger, bodys []byte) []byte
}

//Choose get mode string and return dispatch instance
func Choose(mode string, urlPath string, client *http.Client) Dispatcher {
	if mode == API {
		return api.New(urlPath, client)
	}
	return nil
}
