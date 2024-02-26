package workers

import (
	"dmbb.com/go2/common/logging"
	"fmt"
	"kitchen/model"
	"time"
)

type simpleWorker struct {
	id             string
	dishQueue      chan *model.DishItem
	readyDishQueue chan *model.DishItem
	stopChan       chan string
	logger         logging.Logger
}

func New(name string, dishQueue chan *model.DishItem, readyDishQueue chan *model.DishItem) Worker {
	return &simpleWorker{
		id:             name,
		dishQueue:      dishQueue,
		readyDishQueue: readyDishQueue,
		stopChan:       make(chan string),
		logger:         logging.NewLogger(fmt.Sprintf("kitchen.worker-%v", name)),
	}
}

type Worker interface {
	Start()
	Stop()
	processDishItem(item *model.DishItem)
}

func (worker *simpleWorker) Start() {
	worker.logger.Debug("Start working")
	go func() {
		for {
			select {
			case dishItem := <-worker.dishQueue:
				worker.processDishItem(dishItem)
			case stop := <-worker.stopChan:
				worker.logger.Debug("Stop because &v", stop)
				return
			}
		}
	}()
}

func (worker *simpleWorker) Stop() {
	worker.logger.Debug("Finish")
	worker.stopChan <- "Stop() called"
}

func (worker *simpleWorker) processDishItem(item *model.DishItem) {
	item.Status = model.DishItemInProgress
	for i := 0; i < 5; i++ {
		worker.logger.Info("Cooking %v for %v sec", item.Name, i)
		time.Sleep(1 * time.Second)
	}
	item.Status = model.DishItemReady
	worker.logger.Info("Finished cooking %v", item.Name)
	worker.readyDishQueue <- item
}
