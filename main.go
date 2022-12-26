package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	fmt.Println("Hello world!", e)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, HelloWorld{Message: "hello world"})
	})

	e.Logger.Fatal(e.Start(":3030"))
}
