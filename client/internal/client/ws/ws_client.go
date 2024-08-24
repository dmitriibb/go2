package ws

import (
	"client/internal/constants"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"github.com/dmitriibb/go-common/utils"
	"github.com/dmitriibb/go-common/ws"
)

const (
	msgSourceManager = "manager"
	msgSourceWaiter  = "waiter"
)

type messageWrapper struct {
	source  string
	message *ws.Message
}

var logger = logging.NewLogger("wsClient")

// TODO move somewhere
var managerServiceWsUrl = fmt.Sprintf("ws://%v:%v%v",
	utils.GetEnvProperty(constants.ManagerServiceUrl),
	utils.GetEnvProperty(constants.ManagerServiceHttpPort),
	"/ws")

// connect to manager
// manager tells url of waiter
// connect to waiter
// waiter taker ready order - goes through connected clients - if found needed - gives the order - else puts order back
type ClientEventListener interface {
	OnNewWaiterUrl(waiterServiceUrl string)
	OnReadyOrderItem(item *model.ReadyOrderItem)
	OnError(err error)
}

type Client struct {
	clientId            string
	connectionToManager *ws.ClientConnectionWrapper
	connectionToWaiter  *ws.ClientConnectionWrapper
	messagesBuffer      chan *messageWrapper
	eventListener       ClientEventListener
}

func NewWsClient(clientId string, eventListener ClientEventListener) *Client {
	client := &Client{
		clientId:       clientId,
		eventListener:  eventListener,
		messagesBuffer: make(chan *messageWrapper, 10),
	}

	go func() {
		for msg := range client.messagesBuffer {
			switch msg.source {
			case msgSourceManager:
				client.handleMessagesFromManager(msg.message)
			case msgSourceWaiter:
				client.handleMessagesFromWaiter(msg.message)
			default:
				logger.Warn("unexpected message source '%v' for client %v", msg.source, client.clientId)
			}
		}
	}()

	return client
}

func (client *Client) ConnectToManager() {
	wrapper, err := ws.NewClientConnectionWrapper(client.clientId, managerServiceWsUrl, func(wsClient *ws.ClientConnectionWrapper, message ws.Message) {
		client.messagesBuffer <- &messageWrapper{source: msgSourceManager, message: &message}
	})
	if err != nil {
		client.eventListener.OnError(err)
	} else {
		client.connectionToManager = wrapper
	}
}

func (client *Client) ConnectToWaiter(url string) {
	wrapper, err := ws.NewClientConnectionWrapper(client.clientId, url, func(_ *ws.ClientConnectionWrapper, message ws.Message) {
		client.messagesBuffer <- &messageWrapper{source: msgSourceWaiter, message: &message}
	})
	if err != nil {
		client.eventListener.OnError(err)
	} else {
		if client.connectionToWaiter != nil {
			client.connectionToWaiter.Close()
		}
		client.connectionToWaiter = wrapper
	}
}

func (client *Client) handleMessagesFromManager(message *ws.Message) {
	switch message.Type {
	case ws.MessageTypeString:
		client.eventListener.OnNewWaiterUrl(message.Payload)
	default:
		logger.Warn("%v received unexpected message type '%v' with payload '%v'",
			client.clientId,
			message.Type,
			message.Payload,
		)
	}
}

func (client *Client) handleMessagesFromWaiter(message *ws.Message) {
	switch message.Type {
	case model.WsMessageTypeReadyOrderItem:
		readyOrderItem := &model.ReadyOrderItem{}
		err := json.Unmarshal([]byte(message.Payload), readyOrderItem)
		if err != nil {
			logger.Error("%v can't parse message to ReadyOrderItem. payload = '%v'", client.clientId, message.Payload)
		}
		client.eventListener.OnReadyOrderItem(readyOrderItem)
	default:
		logger.Warn("%v received unexpected message type '%v' with payload '%v'",
			client.clientId,
			message.Type,
			message.Payload,
		)
	}
}

func (client *Client) SendMessage(destination string, message string) {
	msg := ws.Message{
		Type:    ws.MessageTypeString,
		Payload: message,
	}
	switch destination {
	case "manager":
		client.connectionToManager.SendMessage(msg)
	case "waiter":
		client.connectionToWaiter.SendMessage(msg)
	default:
		client.eventListener.OnError(errors.New(fmt.Sprintf("can't send message to unknows destination '%v'", destination)))
	}
}
