package ws

import (
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/ws"
	"net/http"
)

var logger = logging.NewLogger("wsHandler")

var clientIdToConnection = make(map[string]*ws.ClientHandler)

func HandleMapping(apiPrefix string) {
	http.HandleFunc(fmt.Sprintf("%s", apiPrefix), handleWsConnection)
}

func handleWsConnection(w http.ResponseWriter, r *http.Request) {
	clientHandler, err := ws.NewClientHandler(w, r, handleWsMessagesFromClient)
	if err != nil {
		logger.Error("can't handle new ws connection because '%v'", err.Error())
	} else {
		clientIdToConnection[clientHandler.ClientId] = clientHandler
	}
}

func handleWsMessagesFromClient(client *ws.ClientHandler, message ws.Message) {
	if message.Type == ws.MessageTypeString {
		client.SendMessage(ws.Message{ws.MessageTypeString, "ws://localhost:9030/ws"})
	} else {
		logger.Warn("unexpected ws message type '%v' from %v", message.Type, client.ClientId)
	}
}
