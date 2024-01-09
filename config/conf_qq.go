package config

import "fmt"

type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"` //登录后的回调地址
}

func (q QQ) GetPath() string {
	if q.AppID == "" || q.Key == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://%s,%s", q.AppID, q.Redirect)
}
