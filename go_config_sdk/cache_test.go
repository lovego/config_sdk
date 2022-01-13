package go_config_sdk

import (
	"fmt"
	"testing"

	"github.com/lovego/strmap"
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

func Test_getSecret(t *testing.T) {
	type args struct {
		conf strmap.StrMap
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "lch",
			args: args{conf: map[string]interface{}{
				"lch": "lch",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := getSecret(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("getSecret() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
