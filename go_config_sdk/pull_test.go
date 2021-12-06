package go_config_sdk

import (
	"fmt"
	"testing"
)

func TestPull(t *testing.T) {
	api := "http://127.0.0.1:3000/config/pull"

	data, err := Pull(api, "123", Arg{
		Project:      "erp",
		Env:          "dev",
		EndPointType: "server",
		Version:      "1.0",
		Hash:         "ae317ce311be12a4ae315a240ec0d304",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
