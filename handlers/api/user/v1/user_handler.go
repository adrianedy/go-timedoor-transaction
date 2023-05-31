package user_api_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var Store = func(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
