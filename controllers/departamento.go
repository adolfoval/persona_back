package controllers

import (
	"net/http"
	"persona/db"
	"persona/models"

	"github.com/labstack/echo/v4"
)

func GetAllDepartments(context echo.Context) error {
	departments := []models.Departamento{}
	if err := db.Database.Select("dep_id", "dep_nombre").Order("dep_nombre ASC").Find(&departments).Error; err != nil {
		return ErrorRes(context, http.StatusUnprocessableEntity, err)
	}

	return Res(context, departments, "")
}
