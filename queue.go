package webqueue

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

func StartLine(rabbitConf RabbitMQConfig, lineConf LineConfig) {
	consumerConn := createRabbitMqConnection(rabbitConf)
	publisherConn := createRabbitMqConnection(rabbitConf)

	defer consumerConn.Close()
	defer publisherConn.Close()

	consumerChannel, err := consumerConn.Channel()
	panicOnError(err, "Could not create consumer channel")
	defer consumerChannel.Close()

	producerChannel, err := publisherConn.Channel()
	panicOnError(err, "Could not create producer channel")
	defer consumerChannel.Close()
	defer producerChannel.Close()

	// Create the exchange we (re)publish to
	createRabbitMqExchange(producerChannel, lineConf)

	q, err := consumerChannel.QueueDeclare(lineConf.Queue, false, false, false, false, nil)
	panicOnError(err, "Could not create queue")
	err = consumerChannel.QueueBind(lineConf.Queue, "#", lineConf.Queue, false, nil)
	panicOnError(err, "Could not bind exchange to queue")

	consumerChannel.Qos(lineConf.MaxConcurrent, 0, false)

	consumer, err := consumerChannel.Consume(q.Name, "", false, false, false, false, nil)
	panicOnError(err, "Could not create consumer")

	forever := make(chan bool)

	processor := Processor{}
	processor.Init(lineConf)

	go func() {
		for d := range consumer {
			Log.Info("Received message: %s", d.Body)
			success := processor.HandleMessage(string(d.Body))

			if success {
				d.Ack(false)
				continue
			}

			d.Ack(false)
			producerChannel.Publish(lineConf.Queue, "", false, false, amqp.Publishing{
				DeliveryMode: d.DeliveryMode,
				Timestamp:    time.Now(),
				Body:         d.Body,
			})
		}
	}()

	Log.Notice("Waiting for messages. Press CTRL+C to exit...")
	<-forever
}

func createRabbitMqConnection(rabbitConf RabbitMQConfig) *amqp.Connection {
	publisherConn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%d/", rabbitConf.Host, rabbitConf.Port))
	panicOnError(err, "Could not connect to RabbitMQ")

	return publisherConn
}

func createRabbitMqExchange(channel *amqp.Channel, lineConf LineConfig) {
	err := channel.ExchangeDeclare(lineConf.Queue, "topic", true, false, false, false, nil)
	panicOnError(err, "Could not create exchange")
}

func PublishToLine(payload string, lineConf LineConfig, rabbitConf RabbitMQConfig) error {
	conn := createRabbitMqConnection(rabbitConf)

	producerChannel, err := conn.Channel()
	panicOnError(err, "Could not create new RabbitMQ channel")

	createRabbitMqExchange(producerChannel, lineConf)

	return producerChannel.Publish(lineConf.Queue, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Transient,
		Timestamp:    time.Now(),
		Body:         []byte(payload),
	})
}
