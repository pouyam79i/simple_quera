package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pouyam79i/simple_quera/job-exec/handler"
)

func main() {
	fmt.Println("Booting Job Executer Server ...")

	e := echo.New()

	// attaching api handlers
	e.POST("/run", handler.RunCode)

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	var address string = "localhost" + ":" + "8095"

	// fmt.Println("Main Connection Key: ", configs.ApiConnections.MainKey)

	e.Logger.Fatal(e.Start(address))
}
