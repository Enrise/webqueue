package webqueue

import (
	"net/http"
)

type ProcessorMiddlewareImplementation struct {
	config LineConfig
}

type ProcessorMiddleware interface {
	HandleMessage(request **http.Request, response **http.Response)
}
