package database

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connection *mongo.Client

var database string

func GetConnection() *mongo.Client {
	fmt.Println(viper.GetString("database.mongodb.collection"))
	if connection == nil {
		clientOptions := options.Client().ApplyURI(viper.GetString("database.mongodb.uri"))

		connection, _ = mongo.Connect(context.TODO(), clientOptions)
	}

	if database == "" {
		database = viper.GetString("database.mongodb.collection")
	}

	return connection
}

func Collection(input string) *mongo.Collection {
	return GetConnection().Database(database).Collection(input)
}
