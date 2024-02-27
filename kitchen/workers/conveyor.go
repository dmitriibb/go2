package workers

import (
	"kitchen/model"
	"kitchen/recipes"
	"time"
)

var conveyorItems = make(chan *OrderItemWrapper, 100)

type OrderItemWrapper struct {
	OrderItem   *model.OrderItem
	RecipeStage *recipes.RecipeStage
	Comment     string
}

// TODO create a list or db with timers and shut them down properly when application stops
func startConveyorTimer(orderItemWrapper *OrderItemWrapper, timeDelaySec int64) {
	go func() {
		time.Sleep(time.Duration(timeDelaySec) * time.Second)
		conveyorItems <- orderItemWrapper
	}()
}
