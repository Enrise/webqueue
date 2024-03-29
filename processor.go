package webqueue

import (
	"net/http"
	"strings"
)

type Processor struct {
	middleware ProcessorMiddleware
	config     LineConfig
}

func (self *Processor) Init(lineConf LineConfig) {
	processorMiddlewareImplementation := ProcessorMiddlewareImplementation{config: lineConf}

	self.config = lineConf
	self.middleware = &ProcessorMiddlewareMongoDBLog{
		processorMiddlewareImplementation,
		&ProcessorMiddlewareHttpClient{processorMiddlewareImplementation, nil},
	}
}

func (self *Processor) HandleMessage(payload string) bool {
	request, err := http.NewRequest("POST", self.config.Target, strings.NewReader(payload))
	if err != nil {
		Log.Error("Could not create new http.Request instance: %v", err)
		return false
	}

	response := &http.Response{}
	self.middleware.HandleMessage(&request, &response)

	return response.StatusCode >= 200 || response.StatusCode < 300
}
