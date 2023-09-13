package routes

import (
	"persona/controllers"

	"github.com/labstack/echo/v4"
)

func RouterInit(echoServer *echo.Echo) {

	echoServer.GET("/", controllers.Index)
	echoServer.GET("/persona/all", controllers.GetAllPersons)
	echoServer.GET("/persona/filtro/:parametro", controllers.GetAllPersonsBy)
	echoServer.GET("/persona/:id", controllers.GetPersonById)
	echoServer.GET("/departamento/all", controllers.GetAllDepartments)
	echoServer.GET("/ciudad/:id", controllers.GetCitiesByDepartment)
	echoServer.POST("/persona/create", controllers.CreatePerson)
	echoServer.PUT("/persona/:id", controllers.UpdatePerson)
	echoServer.PUT("/persona/restore/:id", controllers.RestorePerson)
	echoServer.DELETE("/persona/:id", controllers.DeletePerson)
}
