package dao

import (
	"jxbdproject/common"
	"jxbdproject/model"
	"strconv"
	"strings"
)
//创建评论
func CreateCommentRedis(comment model.Comment){
	commentkey := "comment:"+strconv.Itoa(comment.CommentId)
 	common.Conn.HMSet(commentkey, map[string]interface{}{
		"PosterId":comment.PosterId,
		"CommentConnet":comment.Connet,
		"TotalLikeCount":comment.TotalCount,
		"PracticeName":comment.PracticeName,
		"CreateTime":comment.CreateTime,
		"UpdateTime":comment.UpdateTime,
	})
}
//点赞
func AddLikeCommentId(Like model.Like)(bool,string){
	//comment:commentID:posterId 存储用户集合的key
	commentID := "comment:"+strconv.Itoa(Like.CommentId)+":"+strconv.Itoa(Like.PosterId)
	arr := strings.Split(commentID,":")
	//将点赞用户添加到集合中
	if common.Conn.SAdd(commentID,Like.UserId).Val() == 0{
		//说明此用户已经点过赞了
		TotalLikeCount := common.Conn.HGet(arr[0]+":"+arr[1],"TotalLikeCount").Val()
		return false,TotalLikeCount
	}
	//将评论信息hash表中点赞数+1
	//拿到hash的key
	common.Conn.HIncrBy(arr[0]+":"+arr[1],"TotalLikeCount",1)
	//将用户点赞的信息Id加入集合中
	//用户表的key user:user_id
	userId := "user:"+strconv.Itoa(Like.UserId)
	common.Conn.SAdd(userId,Like.CommentId)
	TotalLikeCount := common.Conn.HGet(arr[0]+":"+arr[1],"TotalLikeCount").Val()
	return true,TotalLikeCount
}
//取消点赞
func CancelLikeComment(CancelLike model.Like)(string){
	commentID := "comment:"+strconv.Itoa(CancelLike.CommentId)+":"+strconv.Itoa(CancelLike.PosterId)
	userId := "user:"+strconv.Itoa(CancelLike.UserId)
	arr := strings.Split(commentID,":")
	common.Conn.SRem(commentID,CancelLike.UserId)//移除评论集合的用户
	common.Conn.SRem(userId,CancelLike.CommentId)//移除用户集合的点赞评论
	common.Conn.HIncrBy(arr[0]+":"+arr[1],"TotalLikeCount",-1)////将哈希表中评论信息hash表中点赞数-1
	TotalLikeCount := common.Conn.HGet(arr[0]+":"+arr[1],"TotalLikeCount").Val()
	return  TotalLikeCount
}
//同步点赞数并且返回同步失败的
func SyncChroCount(commentsId []int){
	for _,commentId := range commentsId{
		key := "comment:"+strconv.Itoa(commentId)
		totalcount := common.Conn.HGet(key,"TotalLikeCount").Val()
		tmp,_ := strconv.ParseInt(totalcount,10,64)
		if err := UpdateLikeCount(commentId,tmp);err != nil{
			//如果出错直接continue因为下一次进行同步时还会将此Id读取出来
			continue
		}
	}

}