package main

import (
	"goUsers/users"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"

	"github.com/joho/godotenv"
)

type format struct {
	Users []users.UsersStruct `json:"users"`
}

func env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	env()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World!")
	})

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, format{users.All()})
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("API_PORT")))

}
