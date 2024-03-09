package pg

import (
	"database/sql"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/utils"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

var host = utils.GetEnvProperty(DbHostEnv)
var portString = utils.GetEnvProperty(DbPortEnv)
var user = utils.GetEnvProperty(DbUserEnv)
var password = utils.GetEnvProperty(DbPasswordEnv)
var dbname = utils.GetEnvProperty(DbNameEnv)

var logger = logging.NewLogger("dbConnections")

func TestConnectPostgres() {
	port, _ := strconv.Atoi(portString)
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
	port, _ := strconv.Atoi(portString)
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
		logger.Debug("db connection close")
	}()
	err = connection.Ping()
	if err != nil {
		panic(err)
	}
	logger.Debug("execute sql")
	return f(connection)
}
func GetConnection() *sql.DB {
	port, _ := strconv.Atoi(portString)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	logger.Debug("db connection open")
	return connection
}
