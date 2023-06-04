package cmd

import (
	"fmt"

	"github.com/pouyam79i/simple_quera/server/internal/app/server/config"
	"github.com/pouyam79i/simple_quera/server/internal/app/server/model"
)

// TODO: complete server builder

// build a server
func CreateServer(sInfo config.ServerInfo) {
	if sInfo.IP == "" {
		sInfo.IP = "localhost"
	}
	fmt.Println("Creating server on ip: ", sInfo.IP, " port: ", sInfo.Port)
	var address string = sInfo.IP + ":" + sInfo.Port
	myServer := model.BuildServer(sInfo.Debug)
	myServer.Logger.Fatal(myServer.Start(address))
}
