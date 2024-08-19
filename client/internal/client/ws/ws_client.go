package ws

import (
	"encoding/json"
	"github.com/dmitriibb/go-common/logging"
	"github.com/gorilla/websocket"
	"net/http"
)

type WsMessage struct {
	Message  string
	SomeData int
}

var logger = logging.NewLogger("wsClient")
var clientToWsConnection = make(map[string]*websocket.Conn)

func ConnectToWs(clientId string, url string) {
	header := http.Header{}
	header.Set("ClientId", clientId)
	conn, response, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		logger.Error("%s can't establish ws connection because %v", clientId, err.Error())
		return
	}

	logger.Debug("%s ws connection response status = %s", clientId, response.Status)
	clientToWsConnection[clientId] = conn
	go handleWsConnection(clientId, conn)
}

// TODO create proper wsClient with channel, CloseMessage etc
func SendMessage(clientId string, messageString string) {
	conn := clientToWsConnection[clientId]
	message := WsMessage{Message: messageString}
	messageBytes, _ := json.Marshal(message)
	err := conn.WriteMessage(websocket.TextMessage, messageBytes)
	if err != nil {
		logger.Error("%s can't send ws message because %s", clientId, err.Error())
	}
}

func handleWsConnection(clientId string, conn *websocket.Conn) {
	defer func() {
		logger.Debug("close ws for %s", clientId)
		conn.Close()
	}()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			logger.Error("%s can't read message from ws because %s", clientId, err.Error())
			break
		}
		logger.Debug("%s received ws message - type %v", clientId, messageType)
		messageAsStruct := WsMessage{}
		err = json.Unmarshal(p, &messageAsStruct)
		if err != nil {
			logger.Error("can't unmarshal ws message %s", string(p))
			continue
		}
		logger.Info("%s received message - %+v", clientId, messageAsStruct)
	}
}
