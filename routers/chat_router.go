package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) ChatRouter() {
	ChatApi := api.ApiGroupApp.ChatApi
	router.GET("chat_group", ChatApi.ChatGroupView)
	router.GET("chat_group_record", ChatApi.ChatListView)

}
