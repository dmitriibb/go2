package service

import (
	"dmbb.com/go2/common/model"
	"dmbb.com/go2/manager/orders/repository"
)

const price = 10.0

func NewOrder(orderApi *model.ClientOrderApi) {
	order := &model.ClientOrder{ClientId: orderApi.ClientId}
	order = repository.SaveOrderInDb(order)
	items := make([]model.OrderItem, 0)
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
}
