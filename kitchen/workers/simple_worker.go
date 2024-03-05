package workers

import (
	"dmbb.com/go2/common/logging"
	"fmt"
	"kitchen/buffers"
	"kitchen/model"
	"kitchen/recipes"
	"kitchen/storage"
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
		logger:   logging.NewLogger(fmt.Sprintf("worker-%v", name)),
	}
}

type Worker interface {
	Start()
	Stop()
}

func (worker *simpleWorker) Start() {
	worker.logger.Debug("Init working")
	go func() {
		for {
			select {
			case newOrderItem := <-buffers.NewOrderItems:
				worker.logger.Info("take new order item %v", newOrderItem)
				worker.processOrderItem(newOrderItem)
			case orderItemWrapper := <-conveyorItems:
				if orderItemWrapper.RecipeStage != nil {
					worker.logger.Info("take not item from conveyor %v", orderItemWrapper)
					worker.cook(orderItemWrapper)
				} else {
					worker.logger.Info("take item without recipe from conveyor %v", orderItemWrapper)
					worker.processOrderItem(orderItemWrapper.OrderItem)
				}
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
	itemWrapper := &OrderItemWrapper{
		OrderItem: item,
	}
	recipe, err := recipes.GetRecipe(item.Name)
	if err != nil {
		comment := fmt.Sprintf("can't get recipe for '%v' because - '%v', Retry again after 30 sec", item.Name, err.Error())
		worker.logger.Error(comment)
		item.Status = model.OrderItemError
		item.Comment = comment
		startConveyorTimer(itemWrapper, 30)
		return
	}

	itemWrapper.RecipeStage = &recipe

	worker.cook(itemWrapper)
}

func (worker *simpleWorker) cook(itemWrapper *OrderItemWrapper) {
	ready := worker.cookRecipeStage(itemWrapper, itemWrapper.RecipeStage)

	if ready && itemWrapper.RecipeStage.Status == recipes.RecipeStageStatusFinished {
		worker.logger.Info("finished to cook %v", itemWrapper.OrderItem)
		saveOrderItemWrapper(itemWrapper)
		itemWrapper.OrderItem.Status = model.OrderItemReady
		buffers.ReadyOrderItems <- itemWrapper.OrderItem
	}
}

// return true if can proceed to the next stage
func (worker *simpleWorker) cookRecipeStage(itemWrapper *OrderItemWrapper, currentStage *recipes.RecipeStage) bool {
	if currentStage.Status == recipes.RecipeStageStatusFinished {
		return true
	}

	currentStage.Status = recipes.RecipeStageStatusInProgress

	// TODO maybe need to check all ingredients at the root stage
	responseChan := make(chan *storage.IngredientsResponse)
	storage.RequireIngredients(currentStage.Ingredients, responseChan)

	response := <-responseChan
	if !response.Success {
		currentStage.Status = recipes.RecipeStageStatusInProgress
		currentStage.Comment = fmt.Sprintf("can't get ingredients because %v. Will try again after 30 sec", response.Comment)
		saveOrderItemWrapper(itemWrapper)
		startConveyorTimer(itemWrapper, 30)
		return false
	}

	worker.logger.Info("cooking - %v", currentStage.Name)
	if currentStage.TimeToWaitSec > 5 {
		worker.logger.Info("keep stage '%v' cooking for %v sec and continue it later", currentStage.Name, currentStage.TimeToWaitSec)
		saveOrderItemWrapper(itemWrapper)
		startConveyorTimer(itemWrapper, currentStage.TimeToWaitSec)
		return false
	}

	timeStart := time.Now().Unix()
	if currentStage.SubStages != nil {
		for _, subStage := range currentStage.SubStages {
			if !worker.cookRecipeStage(itemWrapper, subStage) {
				return false
			}
		}
	}

	time.Sleep(time.Duration(currentStage.TimeToWaitSec) * time.Second)
	worker.logger.Info("%v ready after %v seconds", currentStage.Name, time.Now().Unix()-timeStart)
	currentStage.Status = recipes.RecipeStageStatusFinished
	return true
}
