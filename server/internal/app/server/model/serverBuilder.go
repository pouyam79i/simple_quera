package model

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// This function build and config server for us
func BuildServer(debug bool) *echo.Echo {
	fmt.Println("Building Server ...")
	e := echo.New()
	if debug {
		AttachALL(e)
	} else {
		AttachMain(e)
	}
	fmt.Println("Done Building!")
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	return e
}
