package clientmanager

import (
	"client/internal/client"
	"time"
)

var allClients = make([]client.Client, 0)

func StartRandomClients() {
	startClient("dima")
	time.Sleep(2 * time.Second)
	startClient("vova")
}

func startClient(clientName string) {
	c := client.New(clientName)
	go c.Start()
}
