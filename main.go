package main

import (
	"fmt"

	"github.com/backend-timedoor/go-transaction-module/database"
	"github.com/backend-timedoor/go-transaction-module/routes/api"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	configInit()
	databaseInit()
}

func main() {
	e := echo.New()

	routesConfig(e)

	e.Logger.Print("Starting ", viper.GetString("appName"))
	e.Logger.Fatal(e.Start(":" + viper.GetString("server.port")))
}

func configInit() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func databaseInit() {
	database.GetConnection()
}

func routesConfig(e *echo.Echo) {
	api.RegisterApiV1Routes(e)
}
