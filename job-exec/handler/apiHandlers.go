package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pouyam79i/simple_quera/job-exec/config"
	"github.com/pouyam79i/simple_quera/job-exec/handler/api"
)

func RunCode(c echo.Context) error {
	reqBody := config.ApiCall{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := config.Response{
			Result: false,
			Info:   "Decoding JSON failed!. reason: " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	go api.CodeX(reqBody)
	res := config.Response{
		Result: true,
		Info:   "We have received your code. Your result will be emailed soon.",
	}
	return c.JSON(http.StatusOK, res)
}
