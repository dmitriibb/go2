package repository

import (
	"database/sql"
	"dmbb.com/go2/common/db"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/model"
	"fmt"
)

var logger = logging.NewLogger("ordersRepository")

func SaveInDb(order *model.ClientOrder, price float32) {
	logger.Debug(fmt.Sprintf("insert %v, price %f", order, price))
	query := "INSERT INTO orders (client_id, dish_name, quantity, price) VALUES ($1, $2, $3, $4)"
	//query := fmt.Sprintf("INSERT INTO orders (client_id, dish_name, quantity, price) VALUES (%v, '%v', %v, %v)", order.ClientId, order.DishName, order.Quantity, price)
	f := func(db *sql.DB) any {
		res, err := db.Exec(query, order.ClientId, order.DishName, order.Quantity, price)
		//res, err := db.Exec(query)
		if err != nil {
			logger.Error(fmt.Sprintf("error = %v", err))
		} else {
			logger.Debug(fmt.Sprintf("Order saved - %v", order))
		}
		return res
	}
	db.UseConnection(f)
	//logger.Debug(fmt.Sprintf("Order saved - %v", order))
}
