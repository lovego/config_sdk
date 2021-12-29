package go_config_sdk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/lovego/strmap"
)

type (
	Conf struct {
		strmap.StrMap
	}

	ConfigData struct {
		Code    string  `json:"code"`
		Message string  `json:"message"`
		Data    *Config `json:"data"`
	}

	Config struct {
		Hash string                 `json:"hash"`
		Conf map[string]interface{} `json:"conf" c:""`
	}

	ConfigTag struct {
		Project      string `json:"project" c:"项目"`
		Env          string `json:"env" c:"环境"`
		EndPointType string `json:"endPointType" c:"终端类型"`
		Version      string `json:"version" c:"版本"`
		Hash         string `json:"hash" c:"hash"`
	}
)

func Pull(api, secret string, arg ConfigTag) (*ConfigData, error) {

	addr := arg.Url(api, secret)

	method := "GET"
	client := &http.Client{
		Timeout: time.Second * 10, // Set 10ms timeout.
	}
	req, err := http.NewRequest(method, addr, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil,errors.New("网络错误")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data ConfigData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.Code != "ok" {
		return nil, errors.New(data.Message)
	}

	return &data, nil
}

func (a ConfigTag) Url(host, secret string) string {

	u, err := url.Parse(host)
	if err != nil {
		panic(err)
	}
	value := u.Query()

	if a.Project != "" {
		value.Set("project", a.Project)
	}
	if a.Env != "" {
		value.Set("env", a.Env)
	}
	if a.Version != "" {
		value.Set("version", a.Version)
	}
	if a.EndPointType != "" {
		value.Set("endPointType", a.EndPointType)
	}
	if secret != "" {
		value.Set("secret", secret)
	}
	if a.Hash != "" {
		value.Set("hash", a.Hash)
	}

	u.RawQuery = value.Encode()

	return u.String()
}
