package main

import (
	"goUsers/users" //TODO: Seems like people use github URLs here. But this is not on github yet so is this import path ok?
	"net/http"

	"github.com/labstack/echo"
)

type format struct {
	Users []users.UsersStruct `json:"users"`
}

//TODO: Is the structure of the app ok?
func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World!")
	})

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, format{users.All()})
	})

	//TODO: Add to .ENV
	e.Logger.Fatal(e.Start(":1323"))

}
