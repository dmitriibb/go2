package client

import (
	"client/internal/client/ws"
	"client/internal/utils"
	"fmt"
	"github.com/dmitriibb/go-common/logging"
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"github.com/dmitriibb/go-common/restaurant-common/model/clientmodel"
	"time"
)

const (
	orderItemsLimit int = 3
)

type Client interface {
	Start()
	EnterRestaurant()
	GoToTheTable()
	AskForMenu()
	MakeOrder()
	WaitForOrder()
	Eat()
	Pay()
	LeaveRestaurant()
}

type client struct {
	Name         string
	logger       logging.Logger
	Id           string
	TableNumber  int
	OrderedItems []string
	wsClient     *ws.Client
}

func New(clientName string) Client {
	return &client{
		Name:   clientName,
		logger: logging.NewLogger(fmt.Sprintf("client '%s'", clientName)),
	}
}

func (c *client) Start() {
	//c.EnterRestaurant()
	//c.GoToTheTable()
	//c.AskForMenu()
	//c.MakeOrder()
	c.WaitForOrder()
}

func (c *client) EnterRestaurant() {
	response, err := enterRestaurantViaRest(c.Name, c.Id)
	if err != nil {
		c.logger.Error("can't enter restaurant because '%v'", err.Error())
		return
	}
	if response.Status != clientmodel.EnterRestaurantResponseStatusWelcome {
		c.logger.Error("can't enter restaurant because '%v'", response.Message)
		return
	}

	c.TableNumber = response.TableNumber
	c.logger.Info("enter restaurant. Assigned id '%v', assigned table ", response.ClientId, response.TableNumber)
}

func (c *client) GoToTheTable() {
	if c.TableNumber < 1 {
		c.logger.Error("don't have table to seat")
		panic(fmt.Sprintf("%v - don't have table to seat", c.Name))
	}
	c.logger.Info("seat at the table %v", c.TableNumber)
}

func (c *client) AskForMenu() {
	c.logger.Info("asking for the menu")
	menu, err := askForMenuViaRest()
	if err != nil {
		panic(err.Error())
	}
	c.logger.Info("received menu with %v items", len(menu.Items))

	clientMaxOrder := utils.GetRandomInt(orderItemsLimit) + 1
	for i := 0; i < clientMaxOrder; i++ {
		indexToOrder := utils.GetRandomInt(len(menu.Items))
		itemToOrder := menu.Items[indexToOrder]
		c.OrderedItems = append(c.OrderedItems, itemToOrder.Name)
	}
}

func (c *client) MakeOrder() {
	resp := makeAnOrderViaRest(c.Id, c.OrderedItems)
	c.logger.Info("ordered %v. result - %s", c.OrderedItems, resp)
}

func (c *client) WaitForOrder() {
	c.logger.Info("connection to ws %s", managerServiceWsUrl)
	c.wsClient = ws.NewWsClient(c.Name, c)
	c.wsClient.ConnectToManager()
	time.Sleep(time.Duration(3) * time.Second)
	c.logger.Info("send hello to manager")
	c.wsClient.SendMessage("manager", "Hello manager!")
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		c.logger.Info("waiting for the order for %v sec", i)
	}
}

func (c *client) Eat() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		c.logger.Info("eating...")
	}
}

func (c *client) Pay() {
	c.logger.Info("paying for the order")
}

func (c *client) LeaveRestaurant() {
	c.logger.Info("leaving")
}

func (c *client) OnNewWaiterUrl(waiterServiceUrl string) {
	c.wsClient.ConnectToWaiter(waiterServiceUrl)
	time.Sleep(time.Duration(3) * time.Second)
	c.wsClient.SendMessage("waiter", "hello waiter!")
}

func (c *client) OnReadyOrderItem(item *model.ReadyOrderItem) {
	c.logger.Info("received ready order item %+v", item)
}

func (c *client) OnError(err error) {
	c.logger.Error("on ws error - %v", err.Error())
}
