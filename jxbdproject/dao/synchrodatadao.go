package dao

import (
	"jxbdproject/common"
	"log"
	"time"
)
//获取到所有的评论Id来当作key获取Redis 哈希表中的点赞数值
func CheckAllCommentId()[]int{
	sqlStr := "select comment_id from comments where update_time >= from_unixtime(?)"
	nowtime :=time.Now().Unix()
	//如果此字段上一次修改的时间小于现在的时间减去一周那么不进行修改
	rows,err := common.Db.Query(sqlStr,nowtime-7 * 24 * 3600)
	defer rows.Close()
	if err != nil{
		log.Printf("CheckAllCommentId Db err:%v\n",err)
		return nil
	}
	var (
		commentsId []int
		commentId int
		)

	for rows.Next(){
		rows.Scan(&commentId)
		commentsId = append(commentsId,commentId)
	}
	return commentsId
}
//修改点赞数
func UpdateLikeCount(commentid int,totalcount int64)error{
	sqlStr := "update comments set total_like_count = ? where comment_id = ? "
	_,err := common.Db.Exec(sqlStr,totalcount,commentid)
	if err != nil{
		log.Printf("UpdateLikeCount Db err:%v\n",err)
		return err
	}
	return nil
}