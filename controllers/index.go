package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(context echo.Context) error {
	return context.String(http.StatusOK, "Home.")
}
