package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pouyam79i/simple_quera/job-exec/config"
)

func CodeX(data config.ApiCall) {
	// TODO: send code to codex
	// TODO: get result and email it
	userEmail := data.Email
	executable := data.Code
	fmt.Println("Receiving data from email: ", userEmail)
	res, err := uploadCodex(executable)
	if err == nil {
		Mail(res, data.Email)
	}
}

// TODO: complete code uploader!
func uploadCodex(data config.CodexAPI) (config.ResCodeX, error) {

	// TODO: hold the user token for completing respond!
	// userToken := data.Token

	fmt.Println("Uploading code to codex...")

	queryParams := url.Values{
		"code":     {data.Code},
		"language": {data.Language},
		"input":    {data.Input},
	}

	payload := bytes.NewBuffer([]byte(queryParams.Encode()))
	client := http.Client{}

	req, err := http.NewRequest(http.MethodPost, config.API_CODEX_URL, payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		fmt.Println("Failed to build req. reason:\n", err.Error())
		return config.ResCodeX{}, errors.New("bad request")
	}

	res, err1 := client.Do(req)

	if err1 != nil {
		fmt.Println("Failed to read respond from Codex. Reason:\n", err1.Error())
		return config.ResCodeX{}, errors.New("bad response from codex")
	}

	byteString, _ := io.ReadAll(res.Body)

	var r config.ResCodeX
	err3 := json.Unmarshal(byteString, &r)

	if err3 != nil {
		fmt.Println("Failed to read respond from Codex. Reason:\n", err1.Error())
		return config.ResCodeX{}, errors.New("bad response from codex")
	}

	return r, nil

}
