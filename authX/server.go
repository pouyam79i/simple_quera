package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pouyam79i/simple_quera/authX/config"
	"github.com/pouyam79i/simple_quera/authX/db"
	"github.com/pouyam79i/simple_quera/authX/handler"
)

func main() {

	fmt.Println("Booting AuthX Server ...")

	configs, err := config.LoadAll()
	if err != nil {
		panic(err.Error())
	}

	_ = db.Connect(configs.DBI)

	e := echo.New()

	// attaching api handlers
	e.POST("/signin", handler.SingIn)
	e.POST("/validate", handler.ValidateToken)

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	var address string = configs.Server.IP + ":" + configs.Server.Port

	// fmt.Println("Main Connection Key: ", configs.ApiConnections.MainKey)

	e.Logger.Fatal(e.Start(address))
}
