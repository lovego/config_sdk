package go_config_sdk

import (
	"fmt"
	"testing"
)

func TestPull(t *testing.T) {
	api := "https://cc-qa.threesoft.cn/config/pull?endPointType=server&env=dev&hash=&project=erp&secret=123&version=1.0"
	api = "https://cc-qa.threesoft.cn/config/pull?endPointType=server&env=qa2&hash=&project=erp&secret=123&version=1.0"

	data, err := Pull(api, "123", ConfigTag{
		Project:      "erp",
		Env:          "qa2",
		EndPointType: "server",
		Version:      "1.0",
		Hash:         "",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}

func TestConfigTag_Url(t *testing.T) {
	type fields struct {
		Project      string
		Env          string
		EndPointType string
		Version      string
		Hash         string
	}
	type args struct {
		host   string
		secret string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "lch",
			args: args{
				host:   "https://cc-qa.threesoft.cn/config/pull?project=s&env=dev&version=1.0&endPointType=server&secret=123&hash=",
				secret: "123",
			},
			fields: fields{
				Project:      "erp",
				Env:          "dev",
				EndPointType: "server",
				Version:      "1.0",
				Hash:         "",
			},
			want: "https://cc-qa.threesoft.cn/config/pull?endPointType=server&env=dev&hash=&project=erp&secret=123&version=1.0",
		},
		{
			name: "?",
			args: args{
				host:   "https://cc-qa.threesoft.cn/config/pull",
				secret: "123",
			},
			fields: fields{
				Project:      "erp",
				Env:          "dev",
				EndPointType: "server",
				Version:      "1.0",
				Hash:         "",
			},
			want: "https://cc-qa.threesoft.cn/config/pull?endPointType=server&env=dev&hash=&project=erp&secret=123&version=1.0",
		},
		{
			name: "?",
			args: args{
				host:   "https://cc-qa.threesoft.cn/config/pull",
				secret: "123",
			},
			fields: fields{
				Project:      "erp",
				Env:          "dev",
				EndPointType: "server",
				Version:      "1.0",
				Hash:         "1234",
			},
			want: "https://cc-qa.threesoft.cn/config/pull?endPointType=server&env=dev&hash=1234&project=erp&secret=123&version=1.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ConfigTag{
				Project:      tt.fields.Project,
				Env:          tt.fields.Env,
				EndPointType: tt.fields.EndPointType,
				Version:      tt.fields.Version,
				Hash:         tt.fields.Hash,
			}
			if got := a.Url(tt.args.host, tt.args.secret); got != tt.want {
				t.Errorf("Url() = %v, want %v", got, tt.want)
			}
		})
	}
}
