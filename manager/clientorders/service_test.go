package clientorders

import (
	"github.com/dmitriibb/go-common/restaurant-common/model"
	"testing"
)

func TestNewOrderSuccess(t *testing.T) {

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

	expected := &model.ClientOrderResponseDTO{
		Comment: "Success",
	}

	actual := NewOrder(order)

	if actual != expected {
		t.Errorf("actual %+v, wanted %+v", actual, expected)
	}
}
