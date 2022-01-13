package go_config_sdk

import (
	"errors"
	"net/url"
	"time"

	"github.com/lovego/strmap"
)

var cache Config

var pull bool

var t Timer

type Timer struct {
	T     time.Duration
	Count int
}

func (a *Timer) Clear() {
	a.T = 0
	a.Count = 0
}

// WaitTime
// 1 1
// 2 2
// 3 6
// 4 24
// 5 120
// 6 720
var waitTime = []time.Duration{1, 2, 6, 24, 120}

func (a *Timer) WaitTime() {
	//a.T = a.T * time.Second * time.Duration(a.Count)
	a.Count += 1
	a.T = waitTime[a.Count%len(waitTime)] * time.Second
	time.Sleep(a.T)
}

func GetConfig(addr, secret string, arg ConfigTag) (*Config, error) {

	if cache.Hash == arg.Hash && cache.Hash != "" {
		return &cache, nil
	}

	if !pull {
		pull = true
		req := func(isBreak bool) {
			for true {

				data, err := Pull(addr, secret, arg)
				if err != nil {
					if isBreak {
						panic(err)
					}
					t.WaitTime()
					continue
				}
				t.Clear()

				if data.Data == nil {
					continue
				}
				cache = *data.Data
				if isBreak {
					break
				}
				arg.Hash = data.Data.Hash

				secret, _ = getSecret(cache.Conf)
			}
		}
		req(true)
		go req(false)
	}
	m := strmap.StrMap(cache.Conf)

	c, err := getSecret(m)
	if err != nil {
		return nil, err
	}

	if c != secret {
		return nil, errors.New("密码错误")
	}

	return &cache, nil
}

// 添加密码在线修改校验
func getSecret(conf strmap.StrMap) (s string, err error) {

	if conf == nil {
		return "", errors.New("配置不能为空")
	}

	// 捕获panic
	// recover能捕获当前的panic
	defer func() {
		if message := recover(); message != nil {
			msg, ok := message.(string)
			if !ok {
				err = errors.New("未知错误")
				return
			}
			err = errors.New(msg)
		}
	}()

	address := conf.Get("configCenter").GetString("pull")
	if address == "" {
		return "", nil
	}

	u, err := url.Parse(address)
	if err != nil {
		return "", errors.New("地址解析错误")
	}

	return u.Query().Get("secret"), nil
}
