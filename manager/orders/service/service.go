package service

import (
	"dmbb.com/go2/common/model"
	"dmbb.com/go2/manager/orders/repository"
	"time"
)

const price = 10.0

// NewOrder TODO add informative response
func NewOrder(orderApi *model.ClientOrderApi) {
	order := &model.ClientOrder{ClientId: orderApi.ClientId}
	order = repository.SaveOrderInDb(order)
	items := make([]model.OrderItem, 0)
	// Items with prices
	for _, itemApi := range orderApi.Items {
		item := model.OrderItem{
			OrderId:  order.Id,
			ClientId: order.ClientId,
			DishName: itemApi.DishName,
			Quantity: itemApi.Quantity,
			Price:    float32(itemApi.Quantity) * price,
		}
		repository.SaveOrderItemInDb(&item)
		items = append(items, item)
	}

	// Items for kitchen
	for _, itemApi := range orderApi.Items {
		for i := 0; i < itemApi.Quantity; i++ {
			dishItem := model.OrderDishItem{
				ClientId:    order.ClientId,
				OrderId:     order.Id,
				DishName:    itemApi.DishName,
				TimeCreated: time.Now(),
				Status:      model.Created,
			}
			id := repository.SaveOrderDishItemInDb(&dishItem)
			repository.SaveOrderDishItemStatus(id, dishItem.TimeCreated, model.Created)
		}
	}
}
