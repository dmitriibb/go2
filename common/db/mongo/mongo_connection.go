package mongo

import (
	"context"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger = logging.NewLogger("MongoConnection")
var initialized = false
var mongoUri = ""

func Init() {
	if initialized {
		logger.Warn("already initialized")
		return
	}
	mongoUri = utils.GetEnvProperty(MongoUriEnv)

	initialized = true
	logger.Debug("initialized")
}

func TestConnection() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	logger.Debug("test connection to mongo")
}

// Deprecated: you must manually close the client
func GetClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}
	return client
}
