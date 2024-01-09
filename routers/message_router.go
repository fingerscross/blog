package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) MessageRouter() {
	MessageApi := api.ApiGroupApp.MessageApi
	router.POST("message", middleware.JwtAuth(), MessageApi.MessageCreateView)
	router.GET("message_all", MessageApi.MessageListAllView)
	router.GET("message", middleware.JwtAuth(), MessageApi.MessageListView)
	router.GET("message_record", middleware.JwtAuth(), MessageApi.MessageRecordView)

}
