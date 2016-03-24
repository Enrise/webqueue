package webqueue

import (
	"fmt"
	"github.com/streadway/amqp"
)

func StartLine(rabbitConf RabbitMQConfig, lineConf LineConfig) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%d/", rabbitConf.Host, rabbitConf.Port))
	panicOnError(err, "Could not connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	panicOnError(err, "Could not create channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("webqueue", false, false, false, false, nil)
	panicOnError(err, "Could not create queue")

	ch.Qos(lineConf.MaxConcurrent, 0, false)

	consumer, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	panicOnError(err, "Could not create consumer")

	forever := make(chan bool)

	processor := Processor{}
	processor.Init(lineConf)

	go func() {
		for d := range consumer {
			Log.Info("Received message: %s", d.Body)
			processor.HandleMessage(string(d.Body))
			// Log.Info("Received message: %s", d.Body)
			// time.Sleep(10000000 * time.Second)
			// respBody, err := processMessage(lineConf, string(d.Body))
			// if err != nil {
			// Log.Warning("Message handling failed: %s", err)
			d.Reject(true)
			// d.Nack(false, true)
			// continue
			// }
			// Log.Info("Message handling successful, target response: %s", string(respBody))
			// d.Ack(false)
		}
	}()

	Log.Notice("Waiting for messages. Press CTRL+C to exit...")
	<-forever
}
