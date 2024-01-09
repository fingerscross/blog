package utils

import "gvb_server/global"

func PrintSystem() {

	ip := global.Config.System.Host
	port := global.Config.System.Port
	if ip == "0.0.0.0" { //获取本机的可用ip
		ipList := GetIPList()
		for _, i := range ipList {
			global.Log.Infof("程序运行在: http://%s:%d/api", i, port)
			global.Log.Infof("api文档运行在: http://%s:%d/swagger/index.html", i, port)
		}
	} else { //自己写的ip
		global.Log.Infof("程序运行在: http://%s:%d/api", ip, port)
		global.Log.Infof("api文档运行在: http://%s:%d/swagger/index.html", ip, port)
	}

}
