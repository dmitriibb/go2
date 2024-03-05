package utils

import (
	"dmbb.com/go2/common/logging"
	"fmt"
	"github.com/joho/godotenv"
)

var logger = logging.NewLogger("CommonUtils")

func GetEnvProperty(propertyName string) string {
	envMap, err := godotenv.Read()
	if err != nil {
		logger.Error("can't load .env file. %v", err.Error())
		panic(fmt.Sprintf("can't load .env file. %v", err.Error()))
	}

	res, ok := envMap[propertyName]
	if ok == false {
		panic(fmt.Sprintf("can't find %v in the .env file", propertyName))
	}
	return res
}