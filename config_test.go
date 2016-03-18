package webqueue_test

import (
	. "github.com/enrise/webqueue"
	check "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type ConfigSuite struct{}

var _ = check.Suite(&ConfigSuite{})

func (s *ConfigSuite) TestLoadConfig(c *check.C) {
	config := Config{}
	config.Load("fixtures/config.1.yml")
	c.Assert(
		config,
		check.DeepEquals,
		Config{
			Rabbitmq: RabbitMQConfig{
				Host: "127.0.0.1", Port: 5672, User: "guest", Password: "guest"},
			Lines: []LineConfig{
				{Queue: "foobar", Target: "http://localhost:1234/job", MaxConcurrent: 4},
			},
			Dashboard: DashboardConfig{BindAddress: "0.0.0.0", Port: 7809},
			MongoDB:   MongoConfig{Host: "localhost", Database: "webqueue", Timeout: 2},
		})
}
