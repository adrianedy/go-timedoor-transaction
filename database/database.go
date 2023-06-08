package database

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database string
var connection *mongo.Client

func GetConnection() *mongo.Client {
	if database == "" {
		database = viper.GetString("database.mongodb.database")
	}

	if connection == nil || !checkConnectionStatus() {
		clientOptions := options.Client().ApplyURI(viper.GetString("database.mongodb.uri"))

		connection, _ = mongo.Connect(context.TODO(), clientOptions)
	}

	return connection
}

func checkConnectionStatus() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := connection.Ping(ctx, nil)
	if err != nil {
		return false
	}

	return true
}

func Collection(collection string) *mongo.Collection {
	return GetConnection().Database(database).Collection(collection)
}
