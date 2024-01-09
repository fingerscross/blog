package routers

import "gvb_server/api"

func (router RouterGroup) TagRouter() {
	TagApi := api.ApiGroupApp.TagApi
	router.POST("tags", TagApi.TagCreateView)
	router.GET("tags", TagApi.TagListView)
	router.PUT("tags/:id", TagApi.TagUpdateView)
	router.DELETE("tags", TagApi.TagRemoveView)

}
