package model

//评论
type Comment struct {
	CommentId int `json:"comment_id"`
	PosterId int `json:"poster_id"`
	Connet  string `json:"connet"` // 内容
	TotalCount int `json:"total_count"`
	PracticeName string `json:"practice_name"`//类
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
//点赞
type Like struct{
	PosterId int	`json:"poster_id"`//发布人ID
	CommentId int	`json:"comment_id"` //评论ID
	UserId int		`json:"user_id"`	//点赞人ID
}
//专项练习时使用
type Special struct{
	Variety string `json:"variety"`
	Subject string `json:"subject"`

}
type Mytest struct{
	Testid int `json:"testid"`
	Userid int `json:"userid"`
	Subject string `json:"subject"`
	Score  int `json:"score"`
	Testtime string `json:"testtime"`
	MistakeCount int `json:"mistake_count"`
}
type User struct{
	Id int				`json:"id"`//id自增
	Username string		`json:"username"`//账号
	Password string 	`json:"password"`//密码
	Email 	 string		`json:"email"`//邮箱
	Name     string		`json:"name"`//app中所使用的昵称
}
//type UserFavQuestion struct{
//	Id int	`json:"id"`//id
//	Userid int `json:"userid"`//属于哪个用户
//	Favorite string `json:"favorite"`//收藏题
//}
//计算考试分数的方法
func(mytest *Mytest)GetTestScore() int{
	if mytest.Subject=="c1"{
		return int(100-mytest.MistakeCount)
	}
	return int(100-2*mytest.MistakeCount)
}