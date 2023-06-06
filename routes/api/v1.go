package api

import (
	transactionsV1Handler "github.com/backend-timedoor/go-transaction-module/handlers/api/v1/transactions"
	"github.com/labstack/echo/v4"
)

func RegisterApiV1Routes(e *echo.Echo) {
	v1 := e.Group("api/v1/")

	v1.GET("transactions", transactionsV1Handler.GetTransactions())
}
