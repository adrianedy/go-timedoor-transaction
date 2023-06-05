package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
    "github.com/spf13/viper"
)

func main() {
	e := echo.New()

	viper.SetConfigType("json")
    viper.AddConfigPath(".")
    viper.SetConfigName("config")

    err := viper.ReadInConfig()
    if err != nil {
        e.Logger.Fatal(err)
    }
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	e.Logger.Print("Starting", viper.GetString("appName"))
	e.Logger.Fatal(e.Start(":" + viper.GetString("server.port")))
}