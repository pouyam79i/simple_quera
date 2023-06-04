package api

import (
	"encoding/json"
	"fmt"

	"github.com/pouyam79i/simple_quera/job-exec/config"
)

func Mail(result config.ResCodeX, email string) {
	// TODO: Mail result
	out, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Failed to stringify results")
	} else {
		fmt.Printf("Emailing Res To %s:\n %s\n", email, string(out))
	}
}
