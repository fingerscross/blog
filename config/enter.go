package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"qq"`
	Jwt      Jwt      `yaml:"jwt"`
	Email    Email    `yaml:"email"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Upload   Upload   `yaml:"upload"`
	Redis    Redis    `yaml:"redis"`
	Es       ES       `yaml:"es"`
}

//从yaml注册过来
