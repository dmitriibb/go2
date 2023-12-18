package repository

import (
	"database/sql"
	"dmbb.com/go2/common/db"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/model"
	"fmt"
	"strings"
)

var logger = logging.NewLogger("ordersRepository")

func SaveOrderInDb(order *model.ClientOrder) *model.ClientOrder {
	query := "INSERT INTO orders (client_id) VALUES ($1) RETURNING id"
	f := func(db *sql.DB) any {
		if strings.HasSuffix(order.ClientId, "error") {
			logger.Warn("fake panic")
			panic("Fake panic because clientId has error")
		}
		id := -1
		row := db.QueryRow(query, order.ClientId)
		err := row.Scan(&id)
		if err != nil {
			logger.Error("Can't scan inserted id because %v", err)
		} else {
			logger.Debug("Order saved with id %v", id)
		}
		order.Id = id
		return 1
	}
	db.UseConnection(f)
	return order
}

func SaveOrderItemInDb(order *model.OrderItem) {
	query := "INSERT INTO order_items (order_id, client_id, dish_name, quantity, price) VALUES ($1, $2, $3, $4, $5)"
	f := func(db *sql.DB) any {
		res, err := db.Exec(query, order.OrderId, order.ClientId, order.DishName, order.Quantity, order.Price)
		if err != nil {
			logger.Error(fmt.Sprintf("error = %v", err))
		} else {
			logger.Debug(fmt.Sprintf("Order item saved - %v", order))
		}
		return res
	}
	db.UseConnection(f)
}
