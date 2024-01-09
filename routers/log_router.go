package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) LogRouter() {
	LogApi := api.ApiGroupApp.LogAPi
	router.GET("logs", LogApi.LogListView)
	router.DELETE("logs", middleware.JwtAdmin(), LogApi.LogRemoveListView)

}
