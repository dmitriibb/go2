package workers

import (
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go2/waiter/buffers"
)

type Worker struct {
	Name   string
	logger logging.Logger
}

func NewWorker(name string) *Worker {
	return &Worker{
		Name:   name,
		logger: logging.NewLogger(fmt.Sprintf("worker.%s", name)),
	}
}

func (w *Worker) Start() {
	w.logger.Debug("start working")
	go func() {
		for {
			readyItem := <-buffers.ReadyOrderItems
			w.logger.Info("get ready order item %+v and take it to the client", readyItem)
		}
	}()
}
