package route

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/adrianedy/go-timedoor-transaction/api/v1"
)

func Init() *echo.Echo {
	e := echo.New()

	av1 := e.Group("api/v1/")
	{
		av1.GET("transaction", apiV1.PostTransaction())
	}
	
	return e
}