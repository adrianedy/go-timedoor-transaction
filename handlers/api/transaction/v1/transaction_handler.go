package transaction_api_handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Store = func(c echo.Context) error {
	var mongoUri string
	// Check if db username is provided in env, because it affect uri pattern
	if viper.Get("DB_USERNAME") != "" && viper.Get("DB_USERNAME") != nil {
		mongoUri = fmt.Sprintf("mongodb://%s:%s@%s:%s",
			viper.Get("DB_USERNAME"),
			viper.Get("DB_PASSWORD"),
			viper.Get("DB_HOST"),
			viper.Get("DB_PORT"))
	} else {
		mongoUri = fmt.Sprintf("mongodb://%s:%s",
			viper.Get("DB_HOST"),
			viper.Get("DB_PORT"))
	}
	opts := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}

	col := client.Database("timedoor_transaction").Collection("transactions")
	result, err := col.InsertOne(c.Request().Context(), map[string]interface{}{
		"name": "Naruto",
		"products": map[string]interface{}{
			"name":  "Kunai",
			"price": 10,
			"qty":   1,
		},
	})

	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

var Index = func(c echo.Context) error {
	return c.JSON(http.StatusOK, "GET transactions")
}

var Show = func(c echo.Context) error {
	return c.JSON(http.StatusOK, "GET one transaction")
}

var Update = func(c echo.Context) error {
	return c.JSON(http.StatusOK, "PATCH one transaction")
}

var Delete = func(c echo.Context) error {
	return c.JSON(http.StatusOK, "DELETE one transaction")
}
