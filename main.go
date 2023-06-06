package main

import (
	"context"

	"github.com/adrianedy/go-timedoor-transaction/database"
	"github.com/adrianedy/go-timedoor-transaction/route"
	"github.com/spf13/viper"
)

func main() {
	router := route.Init()

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		router.Logger.Fatal(err)
	}

	connection := database.GetConnection()
	defer connection.Disconnect(context.TODO())

	router.Logger.Print("Starting ", viper.GetString("appName"))
	router.Logger.Fatal(router.Start(":" + viper.GetString("server.port")))
}
