package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pouyam79i/simple_quera/authX/api"
	"github.com/pouyam79i/simple_quera/authX/config"
)

func SingIn(c echo.Context) error {

	receivedKey := c.QueryParam("key")

	configs, err := config.GetConfigs()

	if err != nil {
		res := config.API_SIGNIN_RESULT{
			Result: false,
			Token:  "",
			Info:   "Cannot Load Key!. reason: " + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	mainKey := configs.ApiConnections.MainKey

	if strings.Compare(mainKey, receivedKey) != 0 {
		res := config.API_SIGNIN_RESULT{
			Result: false,
			Token:  "",
			Info:   "false api key",
		}
		return c.JSON(http.StatusUnauthorized, res)
	}

	reqBody := config.API_SIGNIN_REQUEST{}
	err = c.Bind(&reqBody)
	if err != nil {
		res := config.API_SIGNIN_RESULT{
			Result: false,
			Token:  "",
			Info:   "Decoding JSON failed!. reason: " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	token, err := api.SingInUser(reqBody)
	res := config.API_SIGNIN_RESULT{}
	code := http.StatusOK
	if err != nil {
		res.Result = false
		res.Token = ""
		res.Info = err.Error()
		code = http.StatusForbidden
	} else {
		res.Result = true
		res.Token = token
		res.Info = "successful"
	}
	return c.JSON(code, res)
}

func ValidateToken(c echo.Context) error {

	receivedKey := c.QueryParam("key")

	configs, err := config.GetConfigs()

	if err != nil {
		res := config.API_IDENTIFY_RESULT{
			Result: false,
			Info:   "Cannot Load Key!. reason: " + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	mainKey := configs.ApiConnections.MainKey

	if strings.Compare(mainKey, receivedKey) != 0 {
		res := config.API_IDENTIFY_RESULT{
			Result: false,
			Info:   "false api key",
		}
		return c.JSON(http.StatusUnauthorized, res)
	}

	reqBody := config.API_IDENTIFY_REQUEST{}
	err = c.Bind(&reqBody)
	if err != nil {
		res := config.API_IDENTIFY_RESULT{
			Result: false,
			Info:   "Decoding JSON failed!. reason: " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	isValid, err := api.CheckToken(reqBody.Token)
	res := config.API_IDENTIFY_RESULT{}
	code := http.StatusOK
	if err != nil {
		res.Result = false
		res.Info = err.Error()
		code = http.StatusForbidden
	} else {
		if isValid {
			res.Result = true
			res.Info = "successful"
		} else {
			res.Info = "invalid token"
			code = http.StatusForbidden
		}
	}
	return c.JSON(code, res)

}
