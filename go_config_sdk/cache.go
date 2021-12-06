package go_config_sdk

import "time"

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

func GetConfig(addr, secret string, arg Arg) (*Config, error) {

	if cache.Hash == arg.Hash && cache.Hash != "" {
		return &cache, nil
	}

	if !pull {
		pull = true
		req := func(isBreack bool) {
			for true {
				data, err := Pull(addr, secret, arg)
				if err != nil {
					t.WaitTime()
					continue
				}
				t.Clear()

				if data.Data == nil {
					continue
				}
				cache = data.Data.Config
				if isBreack {
					break
				}
			}
		}
		req(true)
		go req(false)
	}

	return &cache, nil
}
