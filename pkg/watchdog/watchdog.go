package watchdog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/TencentCloudMiddleWare/ScfWatchDog/pkg/dispatch"
	"go.uber.org/zap"
)

//Watchdog provide the main server
type Watchdog struct {
	logger     *zap.Logger
	client     *http.Client
	dispatcher dispatch.Dispatcher
}

//New return the instance of watchdog
func New(logger *zap.Logger) *Watchdog {
	client := &http.Client{}
	workmode := os.Getenv("WATCHDOG_DSIPATCH_MODE")
	disparhpath := os.Getenv("WATCHDOG_DSIPATCH_PATH")
	disp := dispatch.Choose(workmode, disparhpath, client)
	logger.Info("dispatch init", zap.String("workmode", workmode), zap.String("dispatchpath", disparhpath))
	if disp == nil {
		logger.Panic("WATCHDOG_DSIPATCH_MODE must be set or it's not correct")
	}
	return &Watchdog{
		logger:     logger,
		client:     client,
		dispatcher: disp,
	}
}

//Run will start the main loop
func (w *Watchdog) Run() {
	defer w.logger.Sync()
	w.logger.Info("start watchdog",
		zap.String("version", "version 0.1.0"),
	)
	urladdr := fmt.Sprintf("http://%s:%s", os.Getenv("SCF_RUNTIME_API"), os.Getenv("SCF_RUNTIME_API_PORT"))
	defer func() {
		if r := recover(); r != nil {
			w.client.Post(fmt.Sprintf("%s/runtime/invocation/error", urladdr), "application/json", strings.NewReader("bad watchdog"))
		}
	}()
	execmd := strings.Split(os.Getenv("WATCHDOG_RUN_PATH"), " ")
	if len(execmd) <= 0 {
		w.logger.Panic("you must set the env WATCHDOG_RUN_PATH")
	}
	cmd := exec.Command(execmd[0], execmd[1:]...)
	if cmd.Start() != nil {
		w.logger.Panic("we start the server error")
	}
	//send ready to scf
	w.client.Post(fmt.Sprintf("%s/runtime/init/ready", urladdr), "application/json", nil)
	//get req from scf
	for {
		res, err := w.client.Get(fmt.Sprintf("%s/runtime/invocation/next", urladdr))
		if err != nil {
			w.logger.Panic(err.Error())
		}
		bodys, err := ioutil.ReadAll(res.Body)
		if err != nil {
			w.logger.Panic(err.Error())
		}
		res.Body.Close()
		//dispatch by mode
		w.logger.Debug(string(bodys))
		respbody := w.dispatcher.Process(w.logger, bodys)
		w.client.Post(fmt.Sprintf("%s/runtime/invocation/response", urladdr), "application/json", bytes.NewReader(respbody))
	}
}
