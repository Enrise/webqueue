package webqueue

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ProcessorMiddlewareHttpClient struct {
	ProcessorMiddlewareImplementation
	next ProcessorMiddleware
}

func (self *ProcessorMiddlewareHttpClient) HandleMessage(request **http.Request, response **http.Response) {
	reqBody := RogueRead(&(*request).Body)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(*request)
	(*request).Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		err = errors.New(fmt.Sprintf("Sending message to Target failed: %s", err))
		Log.Warning(err.Error())
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))
	*response = resp

	if err != nil {
		err = errors.New(fmt.Sprintf("Could not read response body from Target: %s", err))
		Log.Warning(err.Error())
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = errors.New(fmt.Sprintf("Target sent negative response code (%d). Response body: %s", resp.StatusCode, respBody))
		Log.Warning(err.Error())
		return
	}

	if self.next != nil {
		self.next.HandleMessage(request, response)
	}
}
