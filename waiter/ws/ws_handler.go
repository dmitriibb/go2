package ws

import (
	"encoding/json"
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/restaurant-common/model"
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
		logger.Debug("'%v' connected to ws", clientHandler.ClientId)
		clientIdToConnection[clientHandler.ClientId] = clientHandler
	}
}

func handleWsMessagesFromClient(client *ws.ClientHandler, message ws.Message) {
	if message.Type == ws.MessageTypeString {
		client.SendMessage(ws.Message{ws.MessageTypeString, "Hello fromm waiter service"})
	} else {
		logger.Warn("unexpected ws message type '%v' from %v", message.Type, client.ClientId)
	}
}

func SendReadyOrderItemToClient(clientId string, item *model.ReadyOrderItem) {
	client, ok := clientIdToConnection[clientId]
	if !ok {
		logger.Error("no ws connection for client %v", clientId)
		// TODO retry for failed order item delivery
		return
	}
	msgJson, _ := json.Marshal(item)
	msgString := string(msgJson)
	client.SendMessage(ws.Message{Type: model.WsMessageTypeReadyOrderItem, Payload: msgString})
}
