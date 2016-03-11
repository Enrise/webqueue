package webqueue

import (
	"github.com/op/go-logging"
	"os"
)

var Log = logging.MustGetLogger("webqueue")

func SetupLogging() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	var format = logging.MustStringFormatter(`%{time:2006/01/02 15:04:05.000} [%{level}] %{message}`)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	logging.SetBackend(backendLeveled)
	logging.SetLevel(logging.DEBUG, "")
}
