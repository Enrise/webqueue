package webqueue

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type RabbitMQConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}
type MongoConfig struct {
	Host     string
	Database string
	Timeout  int
}

type DashboardConfig struct {
	BindAddress string `yaml:"bind_address"`
	Port        int
}

type LineConfig struct {
	Queue         string
	Target        string
	MaxConcurrent int `yaml:"max_concurrent"`
}

type Config struct {
	Rabbitmq  RabbitMQConfig
	Lines     []LineConfig
	Dashboard DashboardConfig
	MongoDB   MongoConfig
}

func (c *Config) Load(filename string) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		Log.Fatalf("Could not load %s: %s", filename, err)
	}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		Log.Fatalf("Could not load %s: %s", filename, err)
	}

	Log.Debug("Loaded configuration from %s", filename)
}
