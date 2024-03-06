package rabbit

import (
	"context"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

var logger = logging.NewLogger("CommonRabbit")
var uri = utils.GetEnvProperty(RabbitMqUriEnv)

func failOnError(err error, msg string) {
	if err != nil {
		logger.Error("%s: %s", msg, err)
	}
}

func TestConnection() {
	conn, err := amqp.Dial(uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	logger.Debug("test connection to rabbit")
	ch, err := conn.Channel()
	failOnError(err, "failed to get channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	logger.Info(" [x] Sent %s\n", body)

}
