package config

import "fmt"

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

//地址用拼接方式

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
