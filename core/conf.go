package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"
)

// 读yaml文件配置
const ConfigFile = "settings.yaml"

func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
		panic(fmt.Errorf("get yaml error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init Unmarshal:%v", err)
	}
	log.Println("Config load init success")
	fmt.Println(c)
	global.Config = c //全局变量
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config) //转化成byte Unmarshal 转化成结构体
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
