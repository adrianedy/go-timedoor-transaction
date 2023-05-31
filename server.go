package main

import (
	"fmt"
	"net/http"

	routes "github.com/adrianedy/go-timedoor-transaction/routes/api"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	routes.RegisterApiRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
