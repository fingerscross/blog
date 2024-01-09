package redis_ser

import (
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"time"
)

const newsIndex = "news_index"

type NewData struct {
	Index    int    `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hotValue"`
	Link     string `json:"link"`
}

//// SetNews 设置某一个数据，重复执行，重复累加
//func SetNews(key string, newData []NewData) error {
//	byteData, _ := json.Marshal(newData)
//	err := global.Redis.HSet(newsIndex, key, byteData).Err()
//	return err
//}
//
//func GetNews(key string) (newData []NewData, err error) {
//	res := global.Redis.HGet(newsIndex, key).Val()
//	err = json.Unmarshal([]byte(res), &newData)
//	return
//}

// SetNews 设置某一个数据，重复执行，重复累加
func SetNews(key string, newData []NewData) error {
	byteData, _ := json.Marshal(newData)
	err := global.Redis.Set(fmt.Sprintf("%s_%s", newsIndex, key), byteData, 10*time.Second).Err() //过期后再去itab查，否则去redis查
	return err
}

func GetNews(key string) (newData []NewData, err error) {
	res := global.Redis.Get(fmt.Sprintf("%s_%s", newsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newData)
	return
}
