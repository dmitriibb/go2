package workers

import (
	"dmbb.com/go2/common/logging"
	"fmt"
	"kitchen/buffers"
	"kitchen/model"
	"time"
)

type simpleWorker struct {
	id       string
	stopChan chan string
	logger   logging.Logger
}

func New(name string) Worker {
	return &simpleWorker{
		id:       name,
		stopChan: make(chan string),
		logger:   logging.NewLogger(fmt.Sprintf("kitchen.worker-%v", name)),
	}
}

type Worker interface {
	Start()
	Stop()
}

func (worker *simpleWorker) Start() {
	worker.logger.Debug("Start working")
	go func() {
		for {
			select {
			case newOrderItem := <-buffers.NewOrderItems:
				worker.processOrderItem(newOrderItem)
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

func (worker *simpleWorker) processOrderItem(item *model.OrderItem) {
	item.Status = model.OrderItemInProgress
	for i := 0; i < 5; i++ {
		worker.logger.Info("Cooking %v for %v sec", item.Name, i)
		time.Sleep(1 * time.Second)
	}
	item.Status = model.OrderItemReady
	worker.logger.Info("Finished cooking %v", item.Name)
	buffers.ReadyOrderItems <- item
}
