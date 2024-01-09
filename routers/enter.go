package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvb_server/api/user_api"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) //设置模式
	router := gin.Default()

	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	router.GET("login", user_api.UserApi{}.QQLoginView)
	//router.GET("", func(c *gin.Context) {
	//	c.String(200, "xxx")
	//})

	apiRouterGroup := router.Group("api")
	apiRouterGroupApp := RouterGroup{apiRouterGroup}
	apiRouterGroupApp.SettingsRouter()
	apiRouterGroupApp.ImagesRouter()
	apiRouterGroupApp.AdvertRouter()
	apiRouterGroupApp.MenuRouter()
	apiRouterGroupApp.UserRouter()
	apiRouterGroupApp.TagRouter()
	apiRouterGroupApp.MessageRouter()
	apiRouterGroupApp.ArticleRouter()
	apiRouterGroupApp.CommentRouter()
	apiRouterGroupApp.NewsRouter()
	apiRouterGroupApp.ChatRouter()
	apiRouterGroupApp.LogRouter()
	apiRouterGroupApp.DataRouter()
	return router
}
