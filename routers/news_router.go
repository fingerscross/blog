package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) NewsRouter() {
	NewsApi := api.ApiGroupApp.NewsApi
	router.GET("news", NewsApi.NewListView)

}
