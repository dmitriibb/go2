package ws

import (
	"encoding/json"
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

type WsMessage struct {
	Message  string
	SomeData int
}

var logger = logging.NewLogger("wsHandler")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var allConnection = make([]*websocket.Conn, 0)

func HandleMapping(apiPrefix string) {
	http.HandleFunc(fmt.Sprintf("%s", apiPrefix), handleWsConnection)
}

func handleWsConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	clientId := r.Header.Get("ClientId")
	if len(clientId) == 0 {
		logger.Error("clientId is empty. Reject ws connection")
		return
	}

	allConnection = append(allConnection, conn)
	logger.Debug("established new ws connection for %s", clientId)
	go handleClientWsConnection(clientId, conn)
}

// TODO - create dedicated reader with channel
func handleClientWsConnection(clientId string, conn *websocket.Conn) {
	defer func() {
		logger.Debug("close ws for %s", clientId)
		conn.Close()
	}()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			logger.Error("can't read message from ws of %s because %s", clientId, err.Error())
			break
		}
		logger.Debug("received ws message from %s - type %v", clientId, messageType)
		messageAsStruct := WsMessage{}
		err = json.Unmarshal(p, &messageAsStruct)
		if err != nil {
			logger.Error("can't unmarshal ws message %s", string(p))
			continue
		}
		logger.Info("received from %v message - %+v", clientId, messageAsStruct)
		msgResponse := WsMessage{Message: fmt.Sprintf("Response for %s - %s", clientId, strings.ToUpper(messageAsStruct.Message))}
		msgRespBytes, _ := json.Marshal(msgResponse)
		conn.WriteMessage(websocket.TextMessage, msgRespBytes)
	}
}
