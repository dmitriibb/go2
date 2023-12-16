package db

import (
	"database/sql"
	"dmbb.com/go2/common/logging"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "qwerty1"
	dbname   = "go2"
)

var logger = logging.NewLogger("dbConnections")

func TestConnectPostgres() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	err = connection.Ping()
	if err != nil {
		panic(err)
	}
	logger.Info(fmt.Sprintf("Successfully connected to '%v'!", dbname))
}
func UseConnection(f func(db *sql.DB) any) any {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	logger.Debug("db connection open")
	defer func() {
		connection.Close()
		//fmt.Println("db connection close")
		logger.Debug("db connection close")
	}()
	err = connection.Ping()
	if err != nil {
		panic(err)
	}
	logger.Debug("execute sql")
	return f(connection)
}
