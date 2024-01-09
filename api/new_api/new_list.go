package new_api

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/requests"
	"io"
	"time"
)

// form 用于query
type paramsQuery struct {
	ID   string `json:"id" form:"id" structs:"id"`
	Size int    `json:"size" form:"size" structs:"size" `
}

type params struct {
	ID   string `json:"id" form:"id" structs:"id"`
	Size int    `json:"size" form:"size" structs:"size"`
}

type header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

//type NewData struct {
//	Index    int    `json:"index"`
//	Title    string `json:"title"`
//	HotValue string `json:"hotValue"`
//	Link     string `json:"link"`
//}

type NewResponse struct {
	Code int                 `json:"code"`
	Data []redis_ser.NewData `json:"data"`
	Msg  string              `json:"msg"`
}

const newAPI = "https://api.codelife.cc/api/top/list"

// const newAPI = "https://api.codelife.cc/api/top/list?id=KqndgxeLl9"
const timeout = 2 * time.Second

// NewListView post请求拿到数据，然后显示在itab
func (NewApi) NewListView(c *gin.Context) {
	var cr params
	var headers header
	var pqr paramsQuery
	err := c.ShouldBindQuery(&pqr)
	err = c.ShouldBindJSON(&cr)
	err = c.ShouldBindHeader(&headers)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//fmt.Println(structs.Map(pqr))
	//fmt.Println(structs.Map(headers))
	if cr.Size == 0 {
		cr.Size = 1
	}
	//var temp map[string]interface{}
	//temp["id"] = cr.ID
	//temp["size"] = cr.Size

	key := fmt.Sprintf("%s-%d", cr.ID, cr.Size)
	newData, _ := redis_ser.GetNews(key) //从redis读新闻数据
	if len(newData) != 0 {               //如果redis有，就读redis并且返回，否则去读itab
		res.OkWithData(newData, c)
		return
	}

	httpResponse, err := requests.Post(newAPI, cr, structs.Map(cr), structs.Map(headers), timeout)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	//fmt.Println(httpResponse)

	var response NewResponse
	byteData, err := io.ReadAll(httpResponse.Body)
	//fmt.Println(string(byteData))
	err = json.Unmarshal(byteData, &response)
	//fmt.Println(response)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if response.Code != 200 {
		res.FailWithMessage(response.Msg, c)
		return
	}
	res.OkWithData(response.Data, c)
	fmt.Println(response.Data)
	redis_ser.SetNews(key, response.Data)
	return
}
