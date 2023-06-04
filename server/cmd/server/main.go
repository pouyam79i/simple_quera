package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pouyam79i/simple_quera/server/internal/app/server/cmd"
	"github.com/pouyam79i/simple_quera/server/internal/app/server/config"
)

func main() {
	file, err := ioutil.ReadFile("./config/servers-info-default.json")
	if err != nil {
		fmt.Println("Error while reading default config file:\n ", err.Error())
		return
	}
	serverInfo := config.ServerInfo{}
	err1 := json.Unmarshal([]byte(file), &serverInfo)
	if err1 != nil {
		fmt.Println("Error while decoding default config file data:\n ", err.Error())
		return
	}
	cmd.CreateServer(serverInfo)
}
