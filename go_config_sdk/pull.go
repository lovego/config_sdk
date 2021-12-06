package go_config_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lovego/strmap"
)

type (
	Conf struct {
		strmap.StrMap
	}

	ConfigData struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    *struct {
			Config Config `json:"config" c:""`
		} `json:"data"`
	}

	Config struct {
		Hash string                 `json:"hash"`
		Conf map[string]interface{} `json:"conf" c:""`
	}

	Arg struct {
		Project      string `json:"project" c:"项目"`
		Env          string `json:"env" c:"环境"`
		EndPointType string `json:"endPointType" c:"终端类型"`
		Version      string `json:"version" c:"版本"`
		Hash         string `json:"hash" c:"hash"`
	}
)

func Pull(api, secret string, arg Arg) (*ConfigData, error) {

	addr := arg.Url(api, secret)

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, addr, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
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

func (a Arg) Url(host, secret string) string {
	return fmt.Sprintf(`%s?project=%s&env=%s&version=%s&endPointType=%s&secret=%s&hash=%s`,
		host, a.Project, a.Env, a.Version, a.EndPointType, secret, a.Hash)
}
