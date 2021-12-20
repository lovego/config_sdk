package go_config_sdk

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {

	req := func() {
		//addr := "https://cc-qa.threesoft.cn/config/pull?project=erp&env=dev&version=1.0&endPointType=server&secret=123&hash="
		addr := "https://cc-qa.threesoft.cn/config/pull"

		data, err := GetConfig(addr, "123", ConfigTag{
			Project:      "erp",
			Env:          "dev",
			EndPointType: "server",
			Version:      "1.0",
			//Hash:         "ae317ce311be12a4ae315a240ec0d304",
		})
		if err != nil {
			panic(err)
		}

		fmt.Println(*data)
	}

	for i := 0; i < 3; i++ {
		req()
	}
}
