package clientorders

import (
	"database/sql"
	"fmt"
	"github.com/dmitriibb/go-common/db/pg"
	"github.com/dmitriibb/go-common/logging"
	"strings"
)

var loggerRepo = logging.NewLogger("ordersRepository")

// for unit tests
var SaveOrderInDb = saveOrderInDb
var SaveOrderItemInDb = saveOrderItemInDb

func saveOrderInDb(txWrapper pg.TxWrapperer, order ClientOrder) (*ClientOrder, error) {
	// will fail with mock tx wrapper
	txWrapperCasted := txWrapper.(*pg.TxWrapper)
	query := "INSERT INTO client_orders (client_id) VALUES ($1) RETURNING id"
	var err error
	if strings.HasSuffix(order.ClientId, "error") {
		loggerRepo.Warn("fake panic")
		panic("Fake panic because clientId has error")
	}
	id := -1
	row := txWrapperCasted.Tx.QueryRow(query, order.ClientId)
	err = row.Scan(&id)
	if err != nil {
		loggerRepo.Error("Can't scan inserted id because %v", err)
	} else {
		loggerRepo.Debug("Order saved with id %v", id)
	}
	order.Id = id
	return &order, err
}

func saveOrderItemInDb(txWrapper pg.TxWrapperer, orderItem ClientOrderItem) (*ClientOrderItem, error) {
	// will fail with mock tx wrapper
	txWrapperCasted := txWrapper.(*pg.TxWrapper)
	query := `INSERT INTO client_order_items 
    	(client_order_id, client_id, dish_name, comment, price)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var err error
	id := -1
	row := txWrapperCasted.Tx.QueryRow(query, orderItem.OrderId, orderItem.ClientId, orderItem.DishName, orderItem.Comment, orderItem.Price)
	err = row.Scan(&id)
	if err != nil {
		loggerRepo.Error(fmt.Sprintf("error = %v", err))
	} else {
		loggerRepo.Debug(fmt.Sprintf("Order item saved with id %v", id))
	}
	orderItem.ItemId = id
	return &orderItem, err
}

func getClientIdByOrderIdFromDb(orderId int) string {
	query := `select client_id from client_order_items where client_order_id = $1 limit 1`
	var res string
	pg.UseConnection(func(db *sql.DB) any {
		row := db.QueryRow(query, orderId)
		err := row.Scan(&res)
		if err != nil {
			loggerRepo.Error("can't getClientIdByOrderIdFromDb because %v", err.Error())
			res = ""
		}
		return nil
	})
	return res
}
