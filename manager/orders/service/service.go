package service

import (
	"dmbb.com/go2/common/model"
	"dmbb.com/go2/manager/orders/repository"
)

const price = 10.0

func NewOrder(order *model.ClientOrder) {
	repository.SaveInDb(order, float32(order.Quantity)*price)
}
