package clientorders

import (
	"database/sql"
	"dmbb.com/go2/common/db"
	"dmbb.com/go2/common/logging"
	"errors"
	"fmt"
	"strings"
	"time"
)

var loggerRepo = logging.NewLogger("ordersRepository")

func SaveOrderInDb(order *ClientOrder) (*ClientOrder, error) {
	query := "INSERT INTO client_orders (client_id) VALUES ($1) RETURNING id"
	f := func(db *sql.DB) any {
		if strings.HasSuffix(order.ClientId, "error") {
			loggerRepo.Warn("fake panic")
			panic("Fake panic because clientId has error")
		}
		id := -1
		row := db.QueryRow(query, order.ClientId)
		err := row.Scan(&id)
		if err != nil {
			loggerRepo.Error("Can't scan inserted id because %v", err)
		} else {
			loggerRepo.Debug("Order saved with id %v", id)
		}
		order.Id = id
		return id
	}
	id := db.UseConnection(f)
	var err error = nil
	if id == -1 {
		err = errors.New("can't save order in DB")
	}
	return order, err
}

func SaveOrderItemInDb(order *ClientOrderItem) {
	query := "INSERT INTO client_order_items (client_order_id, client_id, dish_name, quantity, price) VALUES ($1, $2, $3, $4, $5)"
	f := func(db *sql.DB) any {
		res, err := db.Exec(query, order.OrderId, order.ClientId, order.DishName, order.Quantity, order.Price)
		if err != nil {
			loggerRepo.Error(fmt.Sprintf("error = %v", err))
		} else {
			loggerRepo.Debug(fmt.Sprintf("Order item saved - %v", order))
		}
		return res
	}
	db.UseConnection(f)
}

func SaveOrderDishItemInDb(dishItem *ClientOrderDishItem) int {
	query := `insert into client_order_dish_items (client_order_id, client_id, dish_name, time_created, status) 
		values ($1, $2, $3, $4, $5) 
		returning id`
	var id int
	db.UseConnection(func(db *sql.DB) any {
		row := db.QueryRow(query, dishItem.OrderId, dishItem.ClientId, dishItem.DishName, dishItem.TimeCreated, dishItem.Status)
		err := row.Scan(&id)
		if err != nil {
			loggerRepo.Error(fmt.Sprintf("Cant save order dish item - %v", err))
			panic(fmt.Sprintf("Cant save order dish item - %v", err))
			return -1
		}
		loggerRepo.Debug(fmt.Sprintf("Order dish item saved - %v", dishItem))
		return 1
	})
	return id
}

func SaveOrderDishItemStatus(itemId int, timestamp time.Time, status ClientOrderDishItemStatus) {
	query := "insert into client_order_dish_item_statuses(client_order_dish_item_id, timestamp, status) values ($1, $2, $3)"
	db.UseConnection(func(db *sql.DB) any {
		_, err := db.Exec(query, itemId, timestamp, status)
		if err != nil {
			loggerRepo.Error("Can't save order dish item status because %v", err)
			panic(fmt.Sprintf("Can't save order dish item status because %v", err))
		}
		return 1
	})
}
