package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pouyam79i/simple_quera/server/internal/app/server/config"
)

func SendCodeX(data config.ClientCode) {
	// TODO: the you must match the receiving result to token - save it and email it to user
	go uploadCodex(data)
}

// TODO: complete code uploader!
func uploadCodex(data config.ClientCode) {

	// TODO: hold the user token for completing respond!
	// userToken := data.Token

	queryParams := url.Values{
		"code":     {data.CodeX.Code},
		"language": {data.CodeX.Language},
		"input":    {data.CodeX.Input},
	}

	payload := bytes.NewBuffer([]byte(queryParams.Encode()))
	client := http.Client{}

	req, err := http.NewRequest(http.MethodPost, config.API_CODEX_URL, payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		fmt.Println("Failed to build req. reason:\n", err.Error())
		return
	}

	res, err1 := client.Do(req)

	if err1 != nil {
		fmt.Println("Failed to read respond from Codex. Reason:\n", err1.Error())
	}

	byteString, _ := io.ReadAll(res.Body)

	var r config.ResCodeX
	err3 := json.Unmarshal(byteString, &r)

	if err3 == nil {
		fmt.Println("Respond from codeX:\n", r)
	}
}
