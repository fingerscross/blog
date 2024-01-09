package cron_ser

import (
	"github.com/robfig/cron/v3"
	"time"
)

func CronInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	Cron.AddFunc("0 0 0 * * *", SyncArticleData) //0分0秒0点更新
	Cron.AddFunc("0 0 0 * * *", SyncCommentData)
	Cron.Start()
}

//func CronInit() {
//	timezone, err := time.LoadLocation("Asia/Shanghai")
//	if err != nil {
//		logrus.Error(err.Error())
//		return
//	}
//	cron := gocron.NewScheduler(timezone)
//	cron.Cron("0 0 0 * *").Do(SyncArticleData)
//	cron.Cron("0 0 0 * *").Do(SyncCommentData)
//	cron.StartBlocking()
//}
