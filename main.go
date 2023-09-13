package main

import (
	"net/http"
	"persona/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// datbase, err := db.Database.DB()
	// if err != nil {
	// 	log.Fatal("Error en la bd", err.Error())
	// }
	// models.MigratePerson()
	echoServer := echo.New()
	echoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
	routes.RouterInit(echoServer)
	echoServer.Logger.Fatal(echoServer.Start(":8080"))
}
