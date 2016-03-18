package webqueue

import (
	"fmt"
)

func Webqueue(config Config) {
	Log.Notice("Starting webqueue")

	AppMongoConfig = config.MongoDB

	for line := range config.Lines {
		Log.Notice("Starting line %d", line)

		go StartLine(config.Rabbitmq, config.Lines[line])
	}

	Log.Notice("Starting dashboard")
	go StartDashboard(config)

	forever := make(chan bool)
	<-forever
}

func panicOnError(err error, msg string) {
	if err != nil {
		message := fmt.Sprintf("%s: %s", msg, err)
		Log.Fatal(message)
	}
}
