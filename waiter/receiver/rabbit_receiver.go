package receiver

import (
	"context"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/model"
	"dmbb.com/go2/common/queue/rabbit"
	"dmbb.com/go2/common/utils"
	commonInitializer "dmbb.com/go2/common/utils/initializer"
	"dmbb.com/go2/waiter/buffers"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"strings"
	"time"
)

var logger = logging.NewLogger("RabbitReceiver")
var initializer = commonInitializer.New(logger)
var rabbitUri = utils.GetEnvProperty(rabbit.RabbitMqUriEnv)
var myCtx context.Context
var myCtxCancel context.CancelFunc
var readyOrdersQueueName = utils.GetEnvProperty("READY_ORDERS_QUEUE_NAME")
var readyOrderItemsQueueConfig rabbit.RabbitQueueConfig

func Init(ctx context.Context) {
	initFunc := func() error {
		myCtx, myCtxCancel = context.WithCancel(ctx)

		qConfig, err := rabbit.GetQueueConfig(readyOrdersQueueName)
		if err != nil {
			return err
		}
		readyOrderItemsQueueConfig = qConfig

		listenToReadyOrdersFromKitchen(myCtx)

		//listenToQueueHello()
		//listenToTest(myCtx, "listener A")
		//listenToTest(ctx, "listener B")
		//listenToTest(ctx, "listener C")
		return nil
	}
	initializer.Init(initFunc)
}

func listenToQueueHello() {
	conn, err := amqp.Dial(rabbitUri)
	utils.PanicOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.PanicOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.PanicOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.PanicOnError(err, "Failed to register a consumer")

	go func() {
		for d := range messages {
			logger.Info("Received a message: %s", d.Body)
		}
	}()
}

func listenToTest(ctx context.Context, listenerName string) {
	qConf, _ := rabbit.GetQueueConfig("test2")
	rabbit.SubscribeToQueue(ctx, qConf, func(delivery amqp.Delivery) {
		body := string(delivery.Body)
		dots := strings.Count(body, ".")
		logger.Info("'%v' received: %v (%v)", listenerName, body, dots)
		for i := dots; i > 0; i-- {
			logger.Info("'%v' working: %v", listenerName, i)
			time.Sleep(time.Second)
		}
		delivery.Ack(false)
	})

	go func() {
		time.Sleep(20 * time.Second)
		logger.Info("Cancel my context and unsubscribe ")
		myCtxCancel()
	}()
}

func listenToReadyOrdersFromKitchen(ctx context.Context) {
	rabbit.SubscribeToQueue(ctx, readyOrderItemsQueueConfig, func(delivery amqp.Delivery) {
		var payload model.ReadyOrderItem
		err := json.Unmarshal(delivery.Body, &payload)
		if err != nil {
			logger.Error("can't unmarshal message from %v queue to model.ReadyOrderItem. message: %s",
				readyOrderItemsQueueConfig.Name, delivery.Body)
		}
		delivery.Ack(false)
		buffers.ReadyOrderItems <- &payload

	})
}
