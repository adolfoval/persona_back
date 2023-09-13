package controllers

import (
	"net/http"
	"persona/db"
	"persona/models"

	"github.com/labstack/echo/v4"
)

func GetCitiesByDepartment(context echo.Context) error {
	depId := context.Param("id")
	cities := []models.CiudadResponse{}
	if err := db.Database.Select("ciu_id",
		"ciu_nombre").Where("dep_id = ?", depId).Order("ciu_nombre ASC").Find(&cities).Error; err != nil {
		return ErrorRes(context, http.StatusUnprocessableEntity, err)
	}

	return Res(context, cities, "")
}
