package client

import (
	"client/internal/client/ws"
	"client/internal/utils"
	"fmt"
	"github.com/dmitriibb/go-common/logging"
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
}

func New(clientName string) Client {
	return &client{
		Name:   clientName,
		logger: logging.NewLogger(fmt.Sprintf("client '%s'", clientName)),
	}
}

func (c *client) Start() {
	c.logger.Info("connection to ws %s", managerServiceWsUrl)
	ws.ConnectToWs(c.Name, managerServiceWsUrl)
	time.Sleep(time.Duration(1) * time.Second)
	ws.SendMessage(c.Name, "Hello world!")
	//c.EnterRestaurant()
	//c.GoToTheTable()
	//c.AskForMenu()
	//c.MakeOrder()
	//c.WaitForOrder()
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
	ws.ConnectToWs(c.Id, managerServiceWsUrl)
	time.Sleep(time.Duration(1) * time.Second)
	ws.SendMessage(c.Id, "Hello world!")
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
