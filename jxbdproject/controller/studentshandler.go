package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jxbdproject/common"
	"jxbdproject/dao"
	"jxbdproject/model"
	"log"
	"net/http"
)

//学员根据地址查询驾校
func StudentsCheckSchoolByAddress(c *gin.Context){
	var address model.Common
	if err := c.ShouldBind(&address);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	fmt.Println(address)
	allschool,err:=dao.StudentCheckSchoolByAddress(address.Strings)
	if err!=nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"AllSchool":allschool})
}
//学员根据驾校名字查询驾校
func StudentCheckSchoolByName(c *gin.Context){
	var schoolname model.Common
	if err := c.ShouldBind(&schoolname);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	fmt.Println(schoolname)
	allschool,err := dao.StudentCheckSchoolByName(schoolname.Strings)
	if err!=nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"AllSchool":allschool})
}
//顺序练习
func OrderToPractice(c *gin.Context){
	var object model.Common
	if err := c.ShouldBind(&object);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	fmt.Println(object)
	topics,err := dao.StudentOrderByObject(object.Strings)
	if err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Topics":topics})
 }
 //专项练习
func SpecialPractice(c *gin.Context){
 	var temp model.Special
 	if err := c.ShouldBind(&temp);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
 		return
	}
 	fmt.Println(temp)
 	topics,err := dao.StudentSpecialByObject(temp.Subject,temp.Variety)
 	if err!=nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Topics":topics})

}
//判断选择一起科目一
func SimulationTestOne(c *gin.Context){
	var subject model.Common
	if err := c.ShouldBind(&subject);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	topics,err := dao.StudentJudgeTest40(subject.Strings)
	temp,err2 := dao.StudentSelectTest60(subject.Strings)
	topics =append(topics,temp...)
	if err != nil || err2 != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Topics":topics})
}
//模拟考试科目四
func SimulationTestFour(c *gin.Context){
	var subject model.Common
	if err := c.ShouldBind(&subject);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}//科目几
	topics,err := dao.StudentJudgeTest20(subject.Strings)
	temp,err2 := dao.StudentSelectTest30(subject.Strings)
	topics =append(topics,temp...)
	if err != nil || err2 != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Topics":topics})
}
//添加收藏
func MyfavoriteAdd(c *gin.Context){

	var id model.Common
	if err := c.ShouldBind(&id);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	err := dao.AddFavorite(id.Userid,id.Topicid)
	if err!=nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Message":"收藏成功"})
}
//我的收藏得到userid
func MyFavorite(c *gin.Context){
	var Id model.Common
	if err := c.ShouldBind(&Id);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	topics,err :=dao.MyFavorite(Id.Id)
	if err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Topics":topics})
}
func MyFavoriteDelete(c *gin.Context){
	var id model.Common
	if err := c.ShouldBind(&id);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	if err := dao.DeleteMyFavorite(id.Userid,id.Topicid);err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Message":"删除成功"})

}
func MyMistakesAdd(c *gin.Context){
	var Id model.Common
	if err := c.ShouldBind(&Id);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	if err := dao.AddMistakes(Id.Userid,Id.Topicid);err !=nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Message":"错误增加成功"})
}
func MyMistakeDelete(c *gin.Context){
	//先拿到两个id
	var id model.Common
	if err := c.ShouldBind(&id);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	if err:=dao.DeleteMistake(id.Userid,id.Topicid);err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Message":"删除成功"})

}
//用户中心返回用户的所有信息
func UserCenter(c *gin.Context){
	var userid model.Common
	if err := c.ShouldBind(&userid);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	fmt.Println(userid.Id)
	user := dao.CheckUserById(userid.Id)
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"user":user})
}
//我的错题
func MyMistakes(c *gin.Context){
	var userid model.Common
	if err := c.ShouldBind(&userid);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	topics,err:=dao.MyMistakes(userid.Id)
	if err!=nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Topics":topics})
}
//考试记录提交
func MyTestSubmit(c *gin.Context){
	var mytest model.Mytest
	if err := c.ShouldBind(&mytest);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	fmt.Println(mytest)
	if err := dao.SubmitTest(mytest);err != nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
	}

	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Message":"交卷成功"})

}
//获取考试记录
func MyTest(c *gin.Context){
	var id model.Common
	if err :=c.ShouldBind(&id);err != nil{
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Err":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	alltest,err:=dao.AllTest(id.Id)
	if err!=nil{
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Err" : model.ErrorDBError.Error})
		return
	}
	common.SendNormalResPonse(c,http.StatusOK,gin.H{"Alltest":alltest})
}
//播放视频的请求
func PracticeVideo(c *gin.Context){
	var video model.Video
	//先根据练习名称获取视频相关信息
	if err := c.ShouldBind(&video.PracticeName);err != nil{
		log.Printf("PracticeName for failure err:%v",err)
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Error":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	video.VideoName,_ = dao.CheckVideoName(video.PracticeName)
	c.Header("Content-Type","video/mp4")
	filepath := fmt.Sprintf("G:/goproject/src/jxbdproject/photo/%s",video.VideoName)
	c.File(filepath)
}
//对这个视频创建评论
func CreateComment(c *gin.Context){
	var comment model.Comment
	if err := c.ShouldBind(&comment);err != nil || comment.PracticeName == ""{
		log.Printf("Comment for failure err : %v",err)
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Error":model.ErrorRequestBodyParseFailed.Error})
		return
	}
	if !dao.AddComment(comment){
		common.SendErrorResponse(c,model.ErrorDBError.HttpSC,gin.H{"Error":model.ErrorDBError})
		return
	}
	//拿到刚刚在mysql中的记录
	comment = dao.CheckComment(comment.PosterId,comment.PracticeName,comment.Connet)
	//创建hash表存储comment的记录
	dao.CreateCommentRedis(comment)
	common.SendNormalResPonse(c,http.StatusCreated,gin.H{"message":"Comment create successfully!"})
}
//对这个评论进行点赞返回点赞数量
func AddLikeComment(c *gin.Context){
	var Like model.Like
	if err := c.ShouldBind(&Like);err != nil || Like.CommentId == 0{
		log.Printf("LikeComment for failure err:%v\n",err)
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Error":model.ErrorRequestBodyParseFailed.Error})
		return
	}

	status,totalCount := dao.AddLikeCommentId(Like)
	//点赞失败
	if !status{
		log.Printf("AddLikeComment failed")//已经点过赞了
		common.SendNormalResPonse(c,http.StatusOK,gin.H{"message":"It's already been liked!","TotalLikeCount":totalCount})
		return
	}
	common.SendNormalResPonse(c,http.StatusCreated,gin.H{"message":"AddLike successfully!","TotalLikeCount":totalCount})
}
	//取消点赞
func CancelLikeComment(c *gin.Context){
	var CancelLike model.Like
	if err := c.ShouldBind(&CancelLike); err != nil || CancelLike.CommentId == 0{
		log.Printf("CancelLikeComment for failure err:%v\n",err)
		common.SendErrorResponse(c,model.ErrorRequestBodyParseFailed.HttpSC,gin.H{"Error":model.ErrorRequestBodyParseFailed.Error})
	}
	totalCount := dao.CancelLikeComment(CancelLike)
	common.SendNormalResPonse(c,http.StatusCreated,gin.H{"message":"CancelLike successfully!","TotalLikeCount":totalCount})
}