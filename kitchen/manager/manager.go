package manager

import (
	"dmbb.com/go2/common/logging"
	"kitchen/buffers"
	"kitchen/model"
	"kitchen/orders/handler"
	"kitchen/workers"
)

var logger = logging.NewLogger("Manager")
var allWorkerList = []string{"dima", "john", "mark", "kate", "alex"}
var activeWorkers = make(map[string]workers.Worker)
var initialized = false

func Init(newOrders chan *handler.PutNewOrderRequest, closeChan chan string) {
	if initialized {
		logger.Warn("Already initialized")
		return
	}

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
	initialized = true
	logger.Debug("initialized")
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

func processReadyOrderItem(readyDish *model.OrderItem) {
	if readyDish.Status != model.OrderItemReady {
		logger.Warn("Received order item '%v' is not ready. Return it to workers")
		buffers.NewOrderItems <- readyDish
		return
	}

	// TODO
	logger.Info("Dish item %v is ready. Now need to do something", readyDish)
}

func startWorkers() {
	for _, workerName := range allWorkerList {
		worker := workers.New(workerName)
		activeWorkers[workerName] = worker
		worker.Start()
	}
}
