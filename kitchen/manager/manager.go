package manager

import (
	"dmbb.com/go2/common/logging"
	commonModel "dmbb.com/go2/common/model"
	"dmbb.com/go2/common/queue/rabbit"
	"dmbb.com/go2/common/utils"
	commonInitializer "dmbb.com/go2/common/utils/initializer"
	"fmt"
	"github.com/mitchellh/hashstructure"
	"kitchen/buffers"
	"kitchen/model"
	"kitchen/orders/handler"
	"kitchen/workers"
)

var logger = logging.NewLogger("Manager")
var allWorkerList = []string{"dima", "john", "mark", "kate", "alex"}
var activeWorkers = make(map[string]workers.Worker)
var initializer = commonInitializer.New(logger)
var readyOrdersQueueName = utils.GetEnvProperty("READY_ORDERS_QUEUE_NAME")
var readyOrderItemsQueueConfig rabbit.RabbitQueueConfig

func Init(newOrders chan *handler.PutNewOrderRequest, closeChan chan string) {
	initializer.Init(func() error {
		qConfig, err := rabbit.GetQueueConfig(readyOrdersQueueName)
		if err != nil {
			return err
		}
		readyOrderItemsQueueConfig = qConfig

		startWorkers()
		go func() {
			for {
				select {
				case newOrder := <-newOrders:
					processNewOrders(newOrder)
				case readyItem := <-buffers.ReadyOrderItems:
					processReadyOrderItem(readyItem)
				case closeMessage := <-closeChan:
					logger.Info("Stop manager because %v", closeMessage)
					return
				}
			}
		}()
		return nil
	})
}

func processNewOrders(newOrder *handler.PutNewOrderRequest) {
	logger.Info("Received new order %v", newOrder)
	for _, orderItem := range newOrder.Items {
		logger.Info("Received new dish order: %v, item: %v, name: %v", newOrder.OrderId, orderItem.ItemId, orderItem.DishName)
		dishItem := &model.OrderItem{
			OrderId: int(newOrder.OrderId),
			ItemId:  int(orderItem.ItemId),
			Name:    orderItem.DishName,
			Comment: orderItem.Comment,
			Status:  model.OrderItemNew,
		}
		buffers.NewOrderItems <- dishItem

	}
}

func processReadyOrderItem(readyOrderItem *model.OrderItem) {
	if readyOrderItem.Status != model.OrderItemReady {
		logger.Warn("Received order item '%v' is not ready. Return it to workers")
		buffers.NewOrderItems <- readyOrderItem
		return
	}

	// TODO
	logger.Info("Dish item %v is ready. Send to %s queue", readyOrderItem, readyOrdersQueueName)
	msg := commonModel.ReadyOrderItem{
		OrderId:  readyOrderItem.OrderId,
		ItemId:   readyOrderItem.ItemId,
		DishName: readyOrderItem.Name,
		Payload:  readyOrderITemToPayload(readyOrderItem),
	}
	rabbit.SendToQueue(readyOrderItemsQueueConfig, msg)
}

func startWorkers() {
	for _, workerName := range allWorkerList {
		worker := workers.New(workerName)
		activeWorkers[workerName] = worker
		worker.Start()
	}
}

// hash is like an actual food
func readyOrderITemToPayload(readyOrderItem *model.OrderItem) string {
	hash, err := hashstructure.Hash(readyOrderItem, nil)
	if err != nil {
		logger.Error("Can't convert ready order item to hash. (%s)", readyOrderItem)
	}
	return fmt.Sprintf("%s", hash)
}
