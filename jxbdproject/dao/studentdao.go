package dao

import (
	"fmt"
	"jxbdproject/common"
	"jxbdproject/model"
	"log"
	"time"
)
//学员查找驾校通过地址
func StudentCheckSchoolByAddress(address string)([]*model.DrvingSchool,error){
	sqlStr := "select * from drving_school where address like ?"
	rows,err := common.Db.Query(sqlStr,"%"+address+"%")
	if err != nil{
		log.Printf("StudentCheckSchoolByAddress Db err : %v\n",err)
		return nil,err
	}
	defer rows.Close()
	var allschool []*model.DrvingSchool
	for rows.Next(){
		school := &model.DrvingSchool{}
		rows.Scan(&school.Id,&school.Name,&school.Address,&school.Phone)
		allschool=append(allschool,school)
	}
	return allschool,nil
}
//学员通过名字查找驾校
func StudentCheckSchoolByName(name string)([]*model.DrvingSchool,error){
	sqlStr := "select * from drving_school where name like ? "
	rows,err := common.Db.Query(sqlStr,"%"+name+"%")
	if err!=nil{
		log.Printf("StudentCheckSchoolByName Db err : %v\n",err)
		return nil,err
	}
	defer rows.Close()
	var allschool []*model.DrvingSchool
	for rows.Next(){
		school := &model.DrvingSchool{}
		rows.Scan(&school.Id,&school.Name,&school.Address,&school.Phone)
		fmt.Println(school)
		allschool=append(allschool,school)
	}
	fmt.Println(allschool)
	return allschool,nil
}
//顺序练习通过参数为科目几
func StudentOrderByObject(objectname string)([]*model.Topic,error){
	sqlStr := "select * from question_bank where subject = ?"
	rows,err := common.Db.Query(sqlStr,objectname)
	if err!=nil{
		log.Printf("StudentOrderByObject Db err : %v\n",err)
		return nil,err
	}
	defer rows.Close()
	var topics []*model.Topic
	for rows.Next() {
		topic := &model.Topic{}
		rows.Scan(&topic.Id,&topic.Question,&topic.A,&topic.B,&topic.C,&topic.D,&topic.Photo,&topic.Answer,&topic.Variety,&topic.Kind,&topic.Subject)
		topics =append(topics,topic)
	}
	return topics,nil

}
//专项练习通过科目几的
func StudentSpecialByObject(objectname string,variety string)([]*model.Topic,error){
	sqlStr := "select * from question_bank where variety like ? and subject = ?"
	rows,err := common.Db.Query(sqlStr,"%"+variety+"%",objectname)
	if err != nil{
		log.Printf("StudentSpecialByObject Db err : %v\n",err)
		return nil,err
	}
	defer rows.Close()
	var topics []*model.Topic
	for rows.Next(){
		topic := &model.Topic{}
		rows.Scan(&topic.Id,&topic.Question,&topic.A,&topic.B,&topic.C,&topic.D,&topic.Photo,&topic.Answer,&topic.Variety,&topic.Kind,&topic.Subject)
		topics =append(topics,topic)
	}
	return topics,nil
}
//随即查询一道判断
func StudentJudgeTest()(model.Topic){
	sqlStr := "SELECT * FROM question_bank where kind = ? ORDER BY RAND() LIMIT 1"
	row := common.Db.QueryRow(sqlStr,1)
	var topic model.Topic
	row.Scan(&topic.Id,&topic.Question,&topic.A,&topic.B,&topic.C,&topic.D,&topic.Photo,&topic.Answer,&topic.Variety,&topic.Kind,&topic.Subject)
	return topic
}
//随机查询一道选择题
func StudentSelectTest()(model.Topic){
	sqlStr := "select * from question_bank where kind = ? order by rand() limit 1 "
	row := common.Db.QueryRow(sqlStr,0)
	var topic model.Topic
	row.Scan(&topic.Id,&topic.Question,&topic.A,&topic.B,&topic.C,&topic.D,&topic.Photo,&topic.Answer,&topic.Variety,&topic.Kind,&topic.Subject)
	return topic

}
//随机查询40个判断
func StudentJudgeTest40(subject string)([]*model.Topic,error){
	sqlStr := "select * from question_bank where kind = ? and subject = ? order by rand() limit 40"
	rows,err := common.Db.Query(sqlStr,1,subject)
	defer rows.Close()
	if err != nil{
		log.Printf("StudentJudgeTest40 Db err : %v\n",err)
		return nil,err
	}
	var topics []*model.Topic
	for rows.Next(){
		topic := &model.Topic{}
		rows.Scan(&topic.Id,&topic.Question,&topic.A,&topic.B,&topic.C,&topic.D,&topic.Photo,&topic.Answer,&topic.Variety,&topic.Kind,&topic.Subject)
		topics=append(topics,topic)
	}
	return topics,nil
}
//20 个判断
func StudentJudgeTest20(subject string)([]*model.Topic,error){
	sqlStr := "select * from question_bank where kind = ? and subject = ? order by rand() limit 20"
	rows,err := common.Db.Query(sqlStr,1,subject)
	defer rows.Close()
	if err != nil{
		log.Printf("StudentJudgeTest20 Db err : %v\n",err)
		return nil,err
	}
	var topics []*model.Topic
	for rows.Next(){
		topic := &model.Topic{}
		rows.Scan(&topic.Id,&topic.Question,&topic.A,&topic.B,&topic.C,&topic.D,&topic.Photo,&topic.Answer,&topic.Variety,&topic.Kind,&topic.Subject)
		topics=append(topics,topic)
	}
	return topics,nil
}
//60个选择
func StudentSelectTest60(subject string)([]*model.Topic,error){
	sqlStr := "select * from question_bank where kind = ? and subject = ? order by rand() limit 60"
	rows,err := common.Db.Query(sqlStr,0,subject)
	if err != nil{
		log.Printf("StudentSelectTest60 Db err : %v\n",err)
		return nil,err
	}
	defer rows.Close()
	var topics []*model.Topic
	for rows.Next(){
		topic := &model.Topic{}
		rows.Scan(&topic.Id,&topic.Question,&topic.A,&topic.B,&topic.C,&topic.D,&topic.Photo,&topic.Answer,&topic.Variety,&topic.Kind,&topic.Subject)
		topics=append(topics,topic)
	}
	return topics,nil
}
//30 个选择
func StudentSelectTest30(subject string)([]*model.Topic,error){
	sqlStr := "select * from question_bank where kind = ? and subject = ?order by rand() limit 30"
	rows,err := common.Db.Query(sqlStr,0,subject)
	defer rows.Close()
	if err != nil{
		log.Printf("StudentSelectTest30 Db err : %v\n",err)
		return nil,err
	}
	var topics []*model.Topic
	for rows.Next(){
		topic := &model.Topic{}
		rows.Scan(&topic.Id,&topic.Question,&topic.A,&topic.B,&topic.C,&topic.D,&topic.Photo,&topic.Answer,&topic.Variety,&topic.Kind,&topic.Subject)
		topics=append(topics,topic)
	}
	return topics,nil
}
func AddFavorite(userid int,topicid int)error{
	//UPDATE `question_bank` SET question = CONCAT(question,'#') WHERE id =39;
	sqlStr := "insert into user_favquestion(userid,favorite) value (?,?)"
	_,err := common.Db.Exec(sqlStr,userid,topicid)
	if err != nil {
		log.Printf("AddFavorite Db err : %v\n",err)
		return err
	}
	return nil

}
//根据id查找题目
func CheckTopicByid(id int)*model.Topic{
	sqlStr := "select *from topic where id = ?"
	row := common.Db.QueryRow(sqlStr,id)
	var topic model.Topic
	row.Scan(&topic.Id,&topic.Question,&topic.A,&topic.B,&topic.C,&topic.D,&topic.Photo,&topic.Answer,&topic.Variety,&topic.Kind,&topic.Subject)
	return &topic

}
//我的收藏
func MyFavorite(id int)([]*model.Topic,error){
	sqlStr := "select favorite from user_favquestion where userid = ?"
	rows,err := common.Db.Query(sqlStr,id)
	defer rows.Close()
	if err != nil{
		log.Printf("MyFavorite Db err : %v\n",err)
		return nil,err
	}
	var topicsid []int
	for rows.Next(){
		var topicid int
		rows.Scan(&topicid)
		topicsid = append(topicsid,topicid)
	}
	var topics []*model.Topic
	for _,value :=range topicsid {
		temp := &model.Topic{}
		temp = CheckTopicByid(value)
		topics=append(topics,temp)
	}
	return topics,nil

}
//移除收藏
func DeleteMyFavorite(userid int,topicid int)error{
	sqlStr := "delete from user_favquestion where userid = ? and favorite = ?"
	_,err := common.Db.Exec(sqlStr,userid,topicid)
	if err!=nil{
		log.Printf("DeleteMyFavorite Db err : %v\n",err)
		return err
	}
	return nil
}
//增加错题
func AddMistakes(userid int,topicid int)error{
	sqlStr := "insert into user_mistakes(userid,mistake) value (?,?)"
	_,err := common.Db.Exec(sqlStr,userid,topicid)
	if err != nil{
		log.Printf("AddMistakes Db err : %v\n",err)
		return err
	}
	return nil
}
//移除错题
func DeleteMistake(userid int,topicid int)error{
	sqlStr := "delete from user_mistakes where userid = ? and mistake = ?"
	_,err := common.Db.Exec(sqlStr,userid,topicid)
	if err!=nil{
		log.Printf("DeleteMistake Db err : %v\n",err)
		return err
	}
	return nil
}
//我的错题显示全部错题
func MyMistakes(id int)([]*model.Topic,error){
	sqlStr := "select mistake from user_mistakes where userid = ?"
	rows,err := common.Db.Query(sqlStr,id)
	defer rows.Close()
	if err != nil{
		log.Printf("MyMistakes Db err : %v\n",err)
		return nil,err
	}
	var topicsid []int
	for rows.Next(){
		var topicid int
		rows.Scan(&topicid)
		topicsid = append(topicsid,topicid)
	}
	var topics []*model.Topic
	for _,value :=range topicsid {
		temp := &model.Topic{}
		temp = CheckTopicByid(value)
		topics=append(topics,temp)
	}
	return topics,nil

}
//提交卷子
func SubmitTest(mytest model.Mytest)error{
	sqlStr := "insert into test(userid,subject,score,testtime) value(?,?,?,?)"
	_,err := common.Db.Exec(sqlStr,mytest.Userid,mytest.Subject,mytest.GetTestScore(),mytest.Testtime)
	if err!=nil{
		log.Printf("SubmitTest Db err : %v\n",err)
		return err
	}
	return nil
}
//获取所有考试
func AllTest(userid int)([]*model.Mytest,error){
	sqlStr := "select * from test where userid = ?"
	rows,err := common.Db.Query(sqlStr,userid)
	defer rows.Close()
	if err!= nil{
		log.Printf("AllTest Db err : %v\n",err)
		return nil,err
	}
	var alltest []*model.Mytest
	for rows.Next(){
		mytest := &model.Mytest{}
		rows.Scan(&mytest.Testid,&mytest.Userid,&mytest.Subject,&mytest.Score,&mytest.Testtime)
		fmt.Println(mytest.Testtime)
		alltest=append(alltest,mytest)
	}
	return alltest,nil

}
//根据联系名称获取视频相关信息
func CheckVideoName(practice string)(string,error){
	sqlStr := "select videoname from videos where practicename = ?"
	row := common.Db.QueryRow(sqlStr,practice)
	var videoName string
	if err := row.Scan(&videoName);err!=nil{
		log.Printf("CheckVideoName Db err : %v\n",err)
		return videoName,err
	}
	return videoName,nil
}
//添加评论
func AddComment(comment model.Comment)(bool){
	sqlStr := "insert into comments(poster_id,comment_connet,practice_name,create_time)value(?,?,?,?)"
	create_name := time.Now().Format("2006-01-02 15:04:05")
	_, err := common.Db.Exec(sqlStr,comment.PosterId,comment.Connet,comment.PracticeName,create_name)
	if err != nil{
		log.Printf("AddComment Db err :%v\n",err)
		return false
	}
	return true
}
func CheckComment(posterid int,practicename string,commentconnet string)(model.Comment){
	sqlStr := "select *from comments where poster_id = ? and practice_name = ? and comment_connet = ?"
	row := common.Db.QueryRow(sqlStr,posterid,practicename,commentconnet)
	var comment model.Comment
	err := row.Scan(&comment.CommentId,&comment.PosterId,&comment.Connet,&comment.TotalCount,&comment.PracticeName,&comment.CreateTime,&comment.UpdateTime)
	if err != nil{
		log.Printf("CheckComment Db err :%v\n",err)
		return comment
	}
	return comment


}