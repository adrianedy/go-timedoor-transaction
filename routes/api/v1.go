package routes

import (
	TransactionHandler "github.com/adrianedy/go-timedoor-transaction/handlers/api/transaction/v1"
	UserHandler "github.com/adrianedy/go-timedoor-transaction/handlers/api/user/v1"
	"github.com/labstack/echo/v4"
)

func RegisterApiRoutes(e *echo.Echo) {
	e.GET("/transactions", TransactionHandler.Index)
	e.POST("/transactions", TransactionHandler.Store)
	e.GET("/transactions/:id", TransactionHandler.Show)
	e.PATCH("/transactions/:id", TransactionHandler.Update)
	e.DELETE("/transactions/:id", TransactionHandler.Delete)

	e.POST("/users", UserHandler.Store)
}
