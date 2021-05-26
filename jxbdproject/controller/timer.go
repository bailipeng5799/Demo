package controller

import (
	"fmt"
	"jxbdproject/dao"
	"time"
)
func SynchroData(){
	commentsId := dao.CheckAllCommentId()
	dao.SyncChroCount(commentsId)
	fmt.Println("成功将Redis数据同步到Mysql中")
}
//设置一个定时任务
//定时将redis数据存入mysql中
//一天刷新一次
func TickerHandler() {
	//每隔一天向tk.C发送一个当前时间
	tk := time.NewTicker(24 * 3600 * time.Second)
	defer tk.Stop()
	for{
		select {
		case <- tk.C:
			go SynchroData()
		}
	}
}