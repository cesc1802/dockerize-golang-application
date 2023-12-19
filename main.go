package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (Todo) TableName() string {
	return "todos"
}

func uri(dbHost, dbPort, dbName, dbUsername, dbPassword string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUsername, dbPassword, dbName, dbPort)
}

func main() {
	connStr := uri(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"))

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	gormDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	if err := gormDB.Ping(); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&Todo{}); err != nil {
		panic(err)
	}

	router := gin.Default()
	TodoHandlers(router, db)
	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))

}
