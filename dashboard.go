package webqueue

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var dashboardConfig Config

func StartDashboard(config Config) {
	dashboardConfig = config
	dashAddress := fmt.Sprintf("%s:%d", config.Dashboard.BindAddress, config.Dashboard.Port)
	go func() {
		http.HandleFunc("/", servePage)
		http.HandleFunc("/api/status", serveStatus)
		http.HandleFunc("/api/latest-jobs", serveLatestJobs)
		http.HandleFunc("/api/queue-info", serveQueueInfo)
		http.HandleFunc("/api/create-job", serveCreateJob)
		err := http.ListenAndServe(dashAddress, nil)
		if err != nil {
			Log.Fatal("Cannot start dashboard on %s: %s", dashAddress, err)
		}
	}()
	Log.Info("Dashboard available at http://%s", dashAddress)
	forever := make(chan bool)
	<-forever
}

func checkMongoStatus(config MongoConfig) statusReport {
	session, err := mgo.DialWithTimeout(config.Host, time.Second)
	if err != nil {
		return statusReport{Program: "MongoDB", Healthy: false, Error: err.Error()}
	}
	defer session.Close()

	return statusReport{Program: "MongoDB", Healthy: true}
}

func servePage(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/html")

	requestFile := req.URL.Path
	if requestFile == "/" {
		requestFile = "/index.html"
	}
	filename := fmt.Sprintf("dashboard%s", requestFile)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		Log.Info("Could not read file \"%s\": %s", filename, err)
		w.WriteHeader(http.StatusNotFound)
	}

	switch {
	case strings.HasSuffix(filename, ".css"):
		w.Header().Set("Content-Type", "text/css")
	case strings.HasSuffix(filename, ".svg"):
		w.Header().Set("Content-Type", "image/svg+xml")
	}
	w.Write(file)
}

type statusReport struct {
	Program string
	Error   string
	Healthy bool
}

func serveStatus(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	status := make([]statusReport, 0)
	mongoStatus := checkMongoStatus(dashboardConfig.MongoDB)
	status = append(status, mongoStatus)

	result, err := json.Marshal(status)
	if err != nil {
		w.WriteHeader(500)
	}
	w.Write(result)

}

func serveLatestJobs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jobs := GetLatestJobs()

	for index, element := range jobs {
		id, ok := element.Id.(bson.ObjectId)
		if !ok {
			Log.Warning("Could not convert mongo ObjectId to ObjectId: %v", id)
			continue
		}
		jobs[index].Timestamp = id.Time()
	}

	result, err := json.Marshal(jobs)
	if err != nil {
		w.WriteHeader(500)
	}
	w.Write(result)
}

func serveQueueInfo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rabbitUrl := fmt.Sprintf(
		"http://%s:%s@%s:%d/api/queues/%%2F/%s?lengths_age=600&lengths_incr=5&msg_rates_age=600&msg_rates_incr=5",
		dashboardConfig.Rabbitmq.User,
		dashboardConfig.Rabbitmq.Password,
		dashboardConfig.Rabbitmq.Host,
		dashboardConfig.Rabbitmq.ManagementPort,
		dashboardConfig.Lines[0].Queue,
	)

	resp, err := http.Get(rabbitUrl)
	if err != nil {
		Log.Warning("Could not fetch info from queue: %v", err)
		w.WriteHeader(500)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		Log.Warning("Could not fetch info from queue: %v", err)
		w.WriteHeader(500)
		return
	}

	w.Write(body)
}

type CreateJobRequest struct {
	Payload string `json:"payload"`
}

func serveCreateJob(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		return
	}

	var reqJson CreateJobRequest
	reqBody, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(reqBody, &reqJson)
	if err != nil {
		Log.Error("Could not decode JSON: %v", err)
		w.WriteHeader(500)
		return
	}

	err = PublishToLine(reqJson.Payload, dashboardConfig.Lines[0], dashboardConfig.Rabbitmq)
	if err != nil {
		Log.Error("Could not publish message to line: %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}
