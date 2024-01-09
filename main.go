package main

import (
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
	"gvb_server/service/cron_ser"
	"gvb_server/utils"
)

// @title gvb_server API文档
// @version 1.0
// @description API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("111")
	global.Log.Error("111")
	global.Log.Infof("111")
	//gorm连接数据库
	global.DB = core.InitGorm()
	core.InitAddrDB()
	defer global.AddrDB.Close()
	//fmt.Println(global.DB)
	//连接redis
	global.Redis = core.ConnectRedis()
	//连接es
	global.ESClient = core.EsConnect()

	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	go cron_ser.CronInit()

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	utils.PrintSystem()
	router.Run(addr)

}
