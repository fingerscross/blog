package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) CommentRouter() {
	CommentApi := api.ApiGroupApp.CommentApi
	router.POST("comment", middleware.JwtAuth(), CommentApi.CommentCreateView)
	router.GET("comment", CommentApi.CommentListView)
	router.GET("comment/:id", CommentApi.CommentDigg)
	router.DELETE("comment/:id", CommentApi.CommentRemoveView)

}
