package controllers

import (
	"fmt"
	"net/http"
	"persona/db"
	"persona/models"

	"github.com/labstack/echo/v4"
)

func CreatePerson(context echo.Context) error {
	fmt.Println(context.Request().Body)
	persona := models.Persona{}
	err := context.Bind(&persona)
	persona.PerEstado = "A"

	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println(persona)

	if err := db.Database.Save(&persona).Error; err != nil {
		return ErrorRes(context, http.StatusUnprocessableEntity, err)
	}

	return Res(context, persona, "Registro creado con exito")
}

func GetAllPersons(context echo.Context) error {
	persons := []models.PersonaSelect{}
	if err := db.Database.Joins("JOIN backend.ciudad ON (ciudad.ciu_id = persona.ciu_id)").
		Joins("INNER JOIN backend.departamento ON (departamento.dep_id = ciudad.dep_id)").
		Select("per_identificacion",
			"per_nombre1", "per_nombre2",
			"per_apellido1", "per_apellido2", "per_correo", "per_telefono",
			"per_estado", "departamento.dep_id", "ciudad.ciu_id, ciudad.ciu_nombre").
		Order("per_nombre1 ASC, per_nombre2 ASC, per_apellido1 ASC, per_apellido2 ASC").
		Where("per_estado = ?", "A").Find(&persons).Error; err != nil {
		return ErrorRes(context, http.StatusUnprocessableEntity, err)
	}
	fmt.Println(persons)
	return Res(context, persons, "")
}

func GetAllPersonsBy(context echo.Context) error {
	busqueda := context.Param("parametro")
	persons := []models.PersonaSelect{}
	if err := db.Database.Joins("JOIN backend.ciudad ON (ciudad.ciu_id = persona.ciu_id)").
		Joins("INNER JOIN backend.departamento ON (departamento.dep_id = ciudad.dep_id)").
		Select("per_identificacion",
			"per_nombre1", "per_nombre2",
			"per_apellido1", "per_apellido2", "per_correo", "per_telefono",
			"per_estado", "departamento.dep_id", "ciudad.ciu_id, ciudad.ciu_nombre").
		Order("per_nombre1 ASC, per_nombre2 ASC, per_apellido1 ASC, per_apellido2 ASC").
		Where("per_estado = ?", "A").
		Where(db.Database.Where("per_nombre1 LIKE ?", "%"+busqueda+"%").
			Or("per_nombre2 LIKE ?", "%"+busqueda+"%").
			Or("per_correo LIKE ?", "%"+busqueda+"%").
			Or("per_identificacion LIKE ?", "%"+busqueda+"%")).
		Find(&persons).Error; err != nil {
		return ErrorRes(context, http.StatusUnprocessableEntity, err)
	}
	fmt.Println(persons)
	return Res(context, persons, "")
}

func GetPersonById(context echo.Context) error {
	id := context.Param("id")
	// name := context.QueryParam("name")
	fmt.Printf("El query param es: '%v'", id)
	// fmt.Printf("El query param es: '%v'", name)
	person := models.Persona{}
	if err := db.Database.Select("per_identificacion", "per_nombre1",
		"per_nombre2", "per_apellido1", "per_apellido2",
		"per_correo", "per_telefono", "per_estado",
		"ciu_id").Where("per_identificacion = ?", id).Find(&person).Error; err != nil {
		return ErrorRes(context, http.StatusUnprocessableEntity, err)
	}
	return Res(context, person, "")
}

func UpdatePerson(context echo.Context) error {
	person := models.Persona{}
	id := context.Param("id")
	err := context.Bind(&person)
	if err != nil {
		return ErrorRes(context, http.DefaultMaxHeaderBytes, err)
	}

	if err := db.Database.Model(&person).Where("per_identificacion = ?", id).Updates(models.Persona{
		PerNombre1:   person.PerNombre1,
		PerNombre2:   person.PerNombre2,
		PerApellido1: person.PerApellido1,
		PerApellido2: person.PerApellido2,
		PerCorreo:    person.PerCorreo,
		PerTelefono:  person.PerTelefono,
		CiuIdRef:     person.CiuIdRef,
	}).Error; err != nil {
		return ErrorRes(context, http.StatusUnprocessableEntity, err)
	}

	return Res(context, person, "Registro actualizado con exito")

}

func DeletePerson(context echo.Context) error {
	person := models.Persona{}
	id := context.Param("id")

	if err := db.Database.Model(&person).
		Where("per_identificacion = ?", id).Updates(map[string]interface{}{
		"per_estado": "I",
		"deleted_at": "NOW()",
	}).Error; err != nil {
		return ErrorRes(context, http.StatusBadRequest, err)
	}

	return Res(context, person, "Registro eliminado con exito.")
}

func RestorePerson(context echo.Context) error {
	person := models.Persona{}
	id := context.Param("id")

	if err := db.Database.Unscoped().Model(&person).
		Where("per_identificacion = ? ", id).Updates(map[string]interface{}{
		"per_estado": "A",
		"updated_at": db.Database.NowFunc(),
		"deleted_at": nil,
	}).Error; err != nil {
		return ErrorRes(context, http.StatusBadRequest, err)
	}

	return Res(context, person, "Registro restaurado con exito.")
}
