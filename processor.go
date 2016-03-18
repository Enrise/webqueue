package webqueue

import (
	"bytes"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
)

type MessageResultIn struct {
	Request  string
	Response string
	Status   int
}

func processMessage(line LineConfig, payload string) (respBody []byte, err error) {
	resp, err := http.Post(line.Target, "", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		err = errors.New(fmt.Sprintf("Sending message to Target failed: %s", err))
		Log.Warning(err.Error())
		writeMessageResult(payload, err.Error(), 0)
		return
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	writeMessageResult(payload, string(respBody), resp.StatusCode)

	if err != nil {
		err = errors.New(fmt.Sprintf("Could not read response body from Target: %s", err))
		Log.Warning(err.Error())
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = errors.New(fmt.Sprintf("Target sent negative response code (%d). Response body: %s", resp.StatusCode, respBody))
		Log.Warning(err.Error())
		return nil, err
	}

	return respBody, err
}

func writeMessageResult(reqBody string, respBody string, status int) {
	session, err := mgo.Dial(AppMongoConfig.Host)
	if err != nil {
		Log.Warning("Could not log message result: %s", err)
		return
	}

	defer session.Close()

	c := session.DB(AppMongoConfig.Database).C("messagelog")
	err = c.Insert(&MessageResultIn{Request: reqBody, Response: respBody, Status: status})
	if err != nil {
		Log.Warning("Unable to write message result %s", err)
		return
	}
}
