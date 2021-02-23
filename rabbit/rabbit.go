package rabbit

import (
	"fmt"

	"github.com/streadway/amqp"
	"vnfco.ir/rabbit/fail"
)

//ConnectToAMPQServerAndCreateChannel to conncet to amq server and return connection
func ConnectToAMPQServerAndCreateChannel(host string, port string, username string, password string) (*amqp.Connection, *amqp.Channel) {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port))
	fail.FailOnError(err, "Failed to connect to RabbitMQ")
	channel, err := connection.Channel()
	fail.FailOnError(err, "Failed to Create Channel to RabbitMQ")
	return connection, channel
}

// CreateOrJoinSimpleQueue A queue on rabbit mq
func CreateOrJoinSimpleQueue(Channel *amqp.Channel, QueueName string) *amqp.Queue {
	queue, err := Channel.QueueDeclare(
		QueueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	fail.FailOnError(err, "Failed to Create Queue Or Join")
	return &queue
}

//Listen on A queue in rabbit mq
func Listen(Channel *amqp.Channel, Queue *amqp.Queue) <-chan amqp.Delivery {
	messages, err := Channel.Consume(
		Queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	fail.ShowError(err, "Failed to Consume a Queue")
	return messages
}

//Publish : send message to the queue
func Publish(Channel *amqp.Channel, Queue *amqp.Queue, Body string, ContentType string) error {
	err := Channel.Publish(
		"",         // exchange
		Queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: ContentType,
			Body:        []byte(Body),
		})
	fail.ShowError(err, "Failed to Send Message to Queue")
	return err
}
