package clientorders

import (
	"context"
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
	SaveOrderInDb = func(txWrapper pg.TxWrapperer, order *ClientOrder) (*ClientOrder, error) {
		savedClientNames = append(savedClientNames, order.ClientId)
		order.Id = expectedOrderId
		return order, nil
	}

	itemId := 0
	savedItems := make([]*ClientOrderItem, 0)
	SaveOrderItemInDb = func(txWrapper pg.TxWrapperer, orderItem *ClientOrderItem) (*ClientOrderItem, error) {
		savedItems = append(savedItems, orderItem)
		orderItem.ItemId = itemId
		itemId++
		return orderItem, nil
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

type txWrapperMock struct {
	pg.TxWrapper
}

func (receiver *txWrapperMock) Commit() {
	transactionMockCommit = true
}

func (receiver *txWrapperMock) Rollback() {
	transactionMockRollback = true
}
