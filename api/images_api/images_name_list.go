package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImagesResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"` //图片路径
	Name string `json:"name"` //图片名称
}

// ImageNameListView 图片名称列表
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Router /api/images_name [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]ImagesResponse}
func (ImagesApi) ImageNameListView(c *gin.Context) {

	var ImageList []ImagesResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&ImageList)
	res.OkWithData(ImageList, c)
	return

}
