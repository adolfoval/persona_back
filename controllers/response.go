package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Res(context echo.Context, data interface{}, message string) error {

	return context.JSON(http.StatusOK, ResData{
		Message: message,
		Data:    data,
	})
}

func ErrorRes(context echo.Context, status int, err error) error {

	return context.JSON(status, err.Error())
}
