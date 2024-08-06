package main

import (
	"client/internal/clientmanager"
	"github.com/dmitriibb/go-common/logging"
)

var logger = logging.NewLogger("Clients-main")

func main() {
	logger.Debug("Starting...")

	go clientmanager.StartRandomClients()

	forever := make(chan int)
	<-forever
}
