package routers

import "gvb_server/api"

func (router RouterGroup) AdvertRouter() {
	AdvertApi := api.ApiGroupApp.AdvertApi
	router.POST("adverts", AdvertApi.AdvertCreateView)
	router.GET("adverts", AdvertApi.AdvertListView)
	router.PUT("adverts/:id", AdvertApi.AdvertUpdateView)
	router.DELETE("adverts", AdvertApi.AdvertRemoveView)

}
