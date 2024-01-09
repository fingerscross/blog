package routers

import "gvb_server/api"

func (router RouterGroup) MenuRouter() {
	MenuApi := api.ApiGroupApp.MenuApi
	router.POST("menus", MenuApi.MenuCreateView)
	router.GET("menus", MenuApi.MenuListView)
	router.GET("menus_names", MenuApi.MenuNameList)
	router.PUT("menus", MenuApi.MenuUpdateView)
	router.DELETE("menus", MenuApi.MenuRemoveView)
	router.GET("menus/:id", MenuApi.MenuDetailView)

}
