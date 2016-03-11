package main

import (
	"flag"
	. "github.com/Enrise/webqueue"
)

func main() {
	SetupLogging()

	configFile := flag.String("c", "webqueue.yml", "Configuration file")

	flag.Parse()

	config := Config{}
	config.Load(*configFile)
	Webqueue(config)
}
