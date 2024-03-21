package utils

import (
	"fmt"
	"github.com/dmitriibb/go2/common/logging"
	"github.com/joho/godotenv"
)

var logger = logging.NewLogger("CommonUtils")

func GetEnvProperty(propertyName string, defaultVals ...string) string {
	envMap, err := godotenv.Read()
	if err != nil {
		logger.Error("can't load .env file. %v", err.Error())
		panic(fmt.Sprintf("can't load .env file. %v", err.Error()))
	}

	res, ok := envMap[propertyName]
	if ok == false {
		if len(defaultVals) > 0 {
			return defaultVals[0]
		} else {
			panic(fmt.Sprintf("can't find %v in the .env file", propertyName))
		}
	}
	return res
}

func PanicOnError(err error, message string) {
	if err == nil {
		return
	}
	logger.Error("Panic '%v' - %v", message, err.Error())
	panic(err.Error())
}
