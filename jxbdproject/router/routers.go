package router

import (
	"github.com/gin-gonic/gin"
	"jxbdproject/controller"
	"jxbdproject/middleware"
)

func InitRouter(r *gin.Engine) *gin.Engine{
	r.Use(middleware.CORSmw())
	usergroup :=r.Group("/api/user")
	{
		usergroup.POST("/userlogin",controller.UserLogin)
		usergroup.POST("/userregist",controller.UserRegist)
	}
	studentgroup := r.Group("/api/student",middleware.JwtAuthMiddleware())//
	{
		studentgroup.GET("/checkschoolbyaddress",controller.StudentsCheckSchoolByAddress)
		studentgroup.GET("/checkschoolbyname",controller.StudentCheckSchoolByName)
		studentgroup.GET("/ordertopractice",controller.OrderToPractice)//顺序练习
		studentgroup.GET("/specialpractice",controller.SpecialPractice)//专项练习
		studentgroup.GET("/simulationtestone",controller.SimulationTestOne)//模拟考试科一
		studentgroup.GET("/simulationtestfour",controller.SimulationTestFour)//模拟考试科四
		studentgroup.POST("/myfavoriteadd",controller.MyfavoriteAdd)//增加收藏
		studentgroup.GET("/myfavorite",controller.MyFavorite)//我的收藏
		studentgroup.POST("/mymistakesadd",controller.MyMistakesAdd)//增加错题
		studentgroup.POST("/myfavoriteDelete",controller.MyFavoriteDelete)
		studentgroup.POST("/mymistakes",controller.MyMistakes)//我的错题
		studentgroup.POST("/mymistakedelete",controller.MyMistakeDelete)//删除错题
		studentgroup.POST("/usercenter",controller.UserCenter)//用户中心信息
		studentgroup.POST("/mytestsubmit",controller.MyTestSubmit)//考试提交
		studentgroup.GET("/mytest",controller.MyTest)//我的考试记录
		//增加了限流两个
		studentgroup.GET("/videos",middleware.RateLimitMiddleware(),controller.PracticeVideo)//播放视频的操作
		//创建评论
		studentgroup.POST("/videocomment",controller.CreateComment)
		studentgroup.POST("/addlikecomment",controller.AddLikeComment)
		studentgroup.POST("/cancellikecomment",controller.CancelLikeComment)
	}
	mastergroup := r.Group("/api/master",middleware.JwtAuthMiddleware())//
	{
		mastergroup.POST("/mastertopicphoto",controller.MasterTopicPhotoAdd)//增加题目照片
		mastergroup.POST("/mastertopicadd",controller.MasterTopicAdd)//增加题目
		mastergroup.GET("/mastertopiccheck",controller.MasterTopicCheck)//查找题目
		mastergroup.GET("/mastertopiccheckbykind",controller.MasterTopicCheckByKind)
		mastergroup.GET("/mastertopicdetail",controller.MasterTopicDetail)
		mastergroup.POST("/mastertopicdelete",controller.MasterTopicDelete)
		mastergroup.POST("/mastertopicupdate",controller.MasterTopicUpdate)//修改题目
		mastergroup.GET("/masterallschool",controller.MasterAllSchool)
		mastergroup.POST("/masterschooladd",controller.MasterSchoolAdd)
		mastergroup.POST("/masterschooldelete",controller.MasterSchoolDelete)
		mastergroup.POST("/masterschoolupdate",controller.MasterSchoolUpadate)
		mastergroup.GET("/masterschoolcheck",controller. MasterSchoolCheck)
		mastergroup.POST("/masteruploadvideo",controller.MasterUploadVideo)//上传视频
		mastergroup.POST("/masteruploadvideodetail",controller.MasterUploadVideoDetail)
		mastergroup.POST("/masterdeletevideo",controller.MasterDeleteVideo)//删除视频
	}
	return r
}