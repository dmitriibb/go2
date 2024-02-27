package storage

import (
	"dmbb.com/go2/common/logging"
	"time"
)

var logger = logging.NewLogger("StorageRepository")
var ingredientsRequests = make(chan *IngredientRequest, 100)
var started = false

func Start(closeChan chan string) {
	if started {
		logger.Warn("Already started")
		return
	}
	started = true
	logger.Debug("Start")
	go func() {
		for {
			select {
			case request := <-ingredientsRequests:
				processRequest(request)
			case closeMessage := <-closeChan:
				logger.Debug("Stop because %v", closeMessage)
				return
			}
		}
	}()
}

func RequireIngredients(ingredients []string, responseChan chan *IngredientsResponse) {
	ingredientsRequests <- &IngredientRequest{ingredients, responseChan}
}

func processRequest(request *IngredientRequest) {
	// TODO use db
	time.Sleep(200 * time.Millisecond)
	request.ResponseChan <- &IngredientsResponse{Success: true}
}

type IngredientRequest struct {
	Ingredients  []string
	ResponseChan chan *IngredientsResponse
}

type IngredientsResponse struct {
	Success bool
	Comment string
}
