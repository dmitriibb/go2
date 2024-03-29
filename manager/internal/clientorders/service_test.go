package clientorders

import (
	"context"
	"errors"
	"github.com/dmitriibb/go-common/db/pg"
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"testing"
)

var transactionMockCommit = false
var transactionMockRollback = false

func TestNewOrderSuccess(t *testing.T) {
	// given
	transactionMockCommit = false
	transactionMockRollback = false

	StartTransaction = func(ctx context.Context) pg.TxWrapperer {
		return &txWrapperMock{}
	}

	savedClientNames := make([]string, 0)
	expectedOrderId := 123
	SaveOrderInDb = func(txWrapper pg.TxWrapperer, order ClientOrder) (*ClientOrder, error) {
		savedClientNames = append(savedClientNames, order.ClientId)
		order.Id = expectedOrderId
		return &order, nil
	}

	itemId := 0
	savedItems := make([]*ClientOrderItem, 0)
	SaveOrderItemInDb = func(txWrapper pg.TxWrapperer, orderItem ClientOrderItem) (*ClientOrderItem, error) {
		orderItem.ItemId = itemId
		itemId++
		savedItem := &orderItem
		savedItems = append(savedItems, savedItem)
		return savedItem, nil
	}

	sentToGrpc := false
	actualOrderId := 0
	SendNewOrderEvent = func(ctx context.Context, ctxCancel context.CancelFunc, order *ClientOrder) {
		sentToGrpc = true
		actualOrderId = order.Id
	}

	//when
	order := &model.ClientOrderDTO{
		ClientId: "test_client",
		Items: []model.OrderItemDTO{
			{
				DishName: "cola",
				Quantity: 1,
			},
			{
				DishName: "coffee",
				Quantity: 2,
			},
		},
	}
	actualResponse := NewOrder(order)

	//then
	expectedResponse := &model.ClientOrderResponseDTO{
		Comment: "Success",
	}
	if actualResponse.Comment != expectedResponse.Comment {
		t.Errorf("actual response %+v, expected response %+v", actualResponse, expectedResponse)
	}
	if actualOrderId != expectedOrderId {
		t.Errorf("actual order id %v, expected order id %v", actualOrderId, expectedOrderId)
	}
	if len(savedItems) != 3 {
		t.Errorf("saved %v order items, expected %v", len(savedItems), 3)
	}
	if !sentToGrpc {
		t.Errorf("Send to kitchen service grpc has not been called")
	}
	if !transactionMockCommit {
		t.Errorf("Transaction not committed")
	}
}

func TestNewOrder_failOnSavingTheSecondItem(t *testing.T) {
	// given
	transactionMockCommit = false
	transactionMockRollback = false

	StartTransaction = func(ctx context.Context) pg.TxWrapperer {
		return &txWrapperMock{}
	}

	savedClientNames := make([]string, 0)
	expectedOrderId := 123
	SaveOrderInDb = func(txWrapper pg.TxWrapperer, order ClientOrder) (*ClientOrder, error) {
		savedClientNames = append(savedClientNames, order.ClientId)
		order.Id = expectedOrderId
		return &order, nil
	}

	itemId := 0
	savedItems := make([]*ClientOrderItem, 0)
	SaveOrderItemInDb = func(txWrapper pg.TxWrapperer, orderItem ClientOrderItem) (*ClientOrderItem, error) {
		if orderItem.DishName == "coffee" {
			return &orderItem, errors.New("fake error - can't save coffe")
		}
		orderItem.ItemId = itemId
		itemId++
		savedItem := &orderItem
		savedItems = append(savedItems, savedItem)
		return savedItem, nil
	}
	SendNewOrderEvent = func(ctx context.Context, ctxCancel context.CancelFunc, order *ClientOrder) {
		t.Errorf("SendNewOrderEvent() must not be called")
	}

	//when
	order := &model.ClientOrderDTO{
		ClientId: "test_client",
		Items: []model.OrderItemDTO{
			{
				DishName: "cola",
				Quantity: 1,
			},
			{
				DishName: "coffee",
				Quantity: 2,
			},
		},
	}
	actualResponse := NewOrder(order)

	//then
	expectedResponse := &model.ClientOrderResponseDTO{
		Comment: "Fail",
	}
	if actualResponse.Comment != expectedResponse.Comment {
		t.Errorf("actual response %+v, expected response %+v", actualResponse, expectedResponse)
	}
	if len(savedItems) != 1 {
		t.Errorf("saved %v order items, expected %v", len(savedItems), 1)
	}
	if transactionMockCommit {
		t.Errorf("Transaction must not be committed")
	}
	if !transactionMockRollback {
		t.Errorf("Transaction not rolled back")
	}
}

type txWrapperMock struct {
	pg.TxWrapper
}

func (receiver *txWrapperMock) Commit() {
	transactionMockCommit = true
}

func (receiver *txWrapperMock) Rollback() {
	transactionMockRollback = true
}
