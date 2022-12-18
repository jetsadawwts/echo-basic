package main

import (
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"	
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}


type Message struct {
	Text  string `json:"text"`
}

var messages = []Message{
	{Text: "Hello World"},
}

type Err struct {
	Message string `json:"message"`
}



func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, messages)
}

func getUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func createUsersHandler(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	users = append(users, u)

	return c.JSON(http.StatusCreated, u)

}

var users = []User{
	{ID: 1, Name: "jane", Age: 26},
	{ID: 2, Name: "nan", Age: 25},
}


func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	e.GET("/", helloHandler)
	e.GET("/users", getUsersHandler)
	e.POST("/users", createUsersHandler)

	log.Fatal(e.Start(":3000"))
}