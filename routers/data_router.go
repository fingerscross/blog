package routers

import "gvb_server/api"

func (router RouterGroup) DataRouter() {
	DataApi := api.ApiGroupApp.DataApi
	router.GET("data_login", DataApi.SevenLogin)
	router.GET("data_sum", DataApi.DataSumView)

}
