package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/pouyam79i/simple_quera/server/internal/app/server/config"
)

// Sing a user
func Authx_SingIn(userInfo config.SingInInfo) (string, error) {

	ui_str, err := json.Marshal(userInfo)
	if err != nil {
		return "", errors.New("cannot marshal given json struct")
	}

	payload := bytes.NewBuffer(ui_str)
	client := http.Client{}

	req, err := http.NewRequest(http.MethodPost, config.API_AUTHX_SIGNIN+config.AUTHX_KEY, payload)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Failed to build req. reason:\n", err.Error())
		return "", errors.New("req building for authentication failed")
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Failed to read respond from Authx. Reason:\n", err.Error())
		return "", errors.New("authentication failed")
	}

	byteString, _ := io.ReadAll(res.Body)

	var r config.ClientMSG
	err = json.Unmarshal(byteString, &r)

	if err != nil {
		return "", errors.New("cannot Unmarshal given json string")
	}

	if r.Result {
		return r.Token, nil
	} else {
		return "", errors.New("invalid token or server error")
	}

}

// Validate user token
func Authx_Validate(token string) (bool, error) {

	if token == "" {
		return false, errors.New("No Token Given")
	}

	userInfo := config.JustToken{
		Token: token,
	}
	ui_str, err := json.Marshal(userInfo)
	if err != nil {
		return false, errors.New("cannot marshal given json struct")
	}

	payload := bytes.NewBuffer(ui_str)
	client := http.Client{}

	req, err := http.NewRequest(http.MethodPost, config.API_AUTHX_VALIDATE+config.AUTHX_KEY, payload)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Failed to build req. reason:\n", err.Error())
		return false, errors.New("req building for authentication failed")
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Failed to read respond from Authx. Reason:\n", err.Error())
		return false, errors.New("authentication failed")
	}

	byteString, _ := io.ReadAll(res.Body)

	var r config.ValidatorResult
	err = json.Unmarshal(byteString, &r)

	if err != nil {
		return false, errors.New("cannot Unmarshal given json string")
	}

	if r.Result {
		return true, nil
	} else {
		return false, errors.New("invalid token or server error")
	}

}
