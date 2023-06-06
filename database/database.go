package database

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connection *mongo.Client
var database string

func GetConnection() *mongo.Client {
	if database == "" {
		database = viper.GetString("database.mongodb.database")
	}

	if connection == nil {
		clientOptions := options.Client().ApplyURI(viper.GetString("database.mongodb.uri"))

		connection, _ = mongo.Connect(context.TODO(), clientOptions)
	}

	return connection
}

func Collection(collection string) *mongo.Collection {
	return GetConnection().Database(database).Collection(collection)
}
