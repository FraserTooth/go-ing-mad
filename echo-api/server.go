package main

import (
	"echo-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func main() {
	db, dbErr := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if dbErr != nil {
    panic("failed to connect database")
  }
	
	// Migrate the schema
  db.AutoMigrate(&models.Note{})
	
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})


	e.Logger.Fatal(e.Start(":1323"))
}