package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"persona/readenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Database = func() (db *gorm.DB) {

	configConex()
	var host = os.Getenv("HOST")
	var database = os.Getenv("DATABASE")
	var user = os.Getenv("USER")
	var pass = os.Getenv("PASS")
	var port = os.Getenv("PORT")
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable "+
		"TimeZone=Asia/Shanghai",
		host, user, pass, database, port)
	if database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "backend.",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		fmt.Println("Error en la conexion", err.Error())
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
		return database
	}
}()

// var db *sql.DB

func configConex() {
	err := readenv.RegisterEnvFile(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
}

// func Close() {
// 	db.Close()
// }

//verificar conexion

func Ping(dbIn *sql.DB) {
	if err := dbIn.Ping(); err != nil {
		panic(err.Error())
	}
}
