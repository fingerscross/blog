package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	AricleApi := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), AricleApi.ArticleCreateView)
	router.GET("articles", AricleApi.ArticleListView)
	router.GET("articles/detail", AricleApi.ArticleDetailByTitleView)
	router.GET("articles/calendar", middleware.JwtAuth(), AricleApi.ArticleCalendarView)
	router.GET("articles/tags", AricleApi.ArticleTagListView)
	router.PUT("articles", AricleApi.ArticleUpdateView)
	router.DELETE("articles", AricleApi.ArticleRemoveView)
	router.POST("articles/coll", middleware.JwtAuth(), AricleApi.ArticleCollCreateView)
	router.GET("articles/coll", middleware.JwtAuth(), AricleApi.ArticleCollListView)
	router.DELETE("articles/coll", middleware.JwtAuth(), AricleApi.ArticleCollBatchRemoveView)
	router.GET("articles/text", AricleApi.FullTextSearchView)
	router.GET("articles/:id", AricleApi.ArticleDetailView)
	router.POST("articles/digg", AricleApi.DiggArticleView)
}
