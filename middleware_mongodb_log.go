package webqueue

import (
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"time"
)

type ProcessorMiddlewareMongoDBLog struct {
	ProcessorMiddlewareImplementation
	next ProcessorMiddleware
}

type messageResultIn struct {
	Request  string
	Response string
	Status   int
	Duration float64
}

type extraResultInfo struct {
	Duration time.Duration
}

func (self *ProcessorMiddlewareMongoDBLog) HandleMessage(request **http.Request, response **http.Response) {
	startTime := time.Now()
	self.next.HandleMessage(request, response)

	writeMessageResult(*request, *response, extraResultInfo{time.Since(startTime)})
}

var mongoSession *mgo.Session

func writeMessageResult(request *http.Request, response *http.Response, extraInfo extraResultInfo) {
	var err error
	if mongoSession == nil {
		mongoSession, err = mgo.DialWithTimeout(AppMongoConfig.Host, time.Duration(AppMongoConfig.Timeout)*time.Second)
		if err != nil {
			Log.Critical("Could not open connection to MongoDB: %s", err)
			return
		}
	}

	reqBody, _ := ioutil.ReadAll(request.Body)
	respBody, _ := ioutil.ReadAll(response.Body)

	// defer session.Close()

	c := mongoSession.DB(AppMongoConfig.Database).C("messagelog")
	err = c.Insert(&messageResultIn{
		Request:  string(reqBody),
		Response: string(respBody),
		Status:   response.StatusCode,
		Duration: extraInfo.Duration.Seconds(),
	})

	if err != nil {
		Log.Warning("Unable to write message result %s", err)
		return
	}
}
