package manager

import (
	"dmbb.com/go2/common/logging"
	"kitchen/model"
	"kitchen/orders/handler"
	"kitchen/workers"
)

var logger = logging.NewLogger("Kitchen.Manager")
var DishQueue = make(chan *model.DishItem, 100)
var ReadyDishQueue = make(chan *model.DishItem, 100)
var allWorkerList = []string{"dima", "john", "mark", "kate", "alex"}
var activeWorkers = make(map[string]workers.Worker)

type manager struct{}

var Manager = &manager{}

func (manager *manager) Start(newOrders chan *handler.PutNewOrderRequest, closeChan chan string) {
	logger.Info("Start manager")
	startWorkers()
	go func() {
		for {
			select {
			case newOrder := <-newOrders:
				processNewOrders(newOrder)
			case readyDish := <-ReadyDishQueue:
				processReadyDish(readyDish)
			case closeMessage := <-closeChan:
				logger.Info("Stop manager because %v", closeMessage)
				return
			}
		}
	}()
}

func processNewOrders(newOrder *handler.PutNewOrderRequest) {
	logger.Info("Received new order %v", newOrder)
	for _, orderDishItem := range newOrder.Items {
		logger.Info("Received new dish %v : %v", orderDishItem.DishName, orderDishItem.Quantity)
		for i := 0; i < int(orderDishItem.Quantity); i++ {
			dishItem := &model.DishItem{
				OrderId: int(newOrder.OrderId),
				Name:    orderDishItem.DishName,
				Status:  model.DishItemNew,
			}
			DishQueue <- dishItem
		}
	}
}

func processReadyDish(readyDish *model.DishItem) {
	if readyDish.Status != model.DishItemReady {
		logger.Warn("Received dish item '%v' is not ready. Return it to workers")
		DishQueue <- readyDish
		return
	}

	// TODO
	logger.Info("Dish item %v si ready. Now need to do something", readyDish)
}

func startWorkers() {
	for _, workerName := range allWorkerList {
		worker := workers.New(workerName, DishQueue, ReadyDishQueue)
		activeWorkers[workerName] = worker
		worker.Start()
	}
}
