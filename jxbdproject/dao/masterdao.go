package dao

import (
	"jxbdproject/common"
	"jxbdproject/model"
	"log"
)
//所有题目
func MasterTopics()([]*model.Topic,error){
		sqlStr := "select *from question_bank"
		rows,err := common.Db.Query(sqlStr)
		if err != nil{
			log.Printf("MasterTopics Db err : %v\n",err)
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

func MasterDetailTopic(id int) model.Topic{
	sqlStr := "select * from question_bank where id = ?"
	row := common.Db.QueryRow(sqlStr,id)
	var topic model.Topic
	row.Scan(&topic.Id,&topic.Question,&topic.A,&topic.B,&topic.C,&topic.D,&topic.Photo,&topic.Answer,&topic.Variety,&topic.Kind,&topic.Subject)
	return topic
}



//管理员添加习题
func MasterAddTopic(topic model.Topic) error{

	sqlStr := "insert into question_bank(id,question,a,b,c,d,photo,answer,variety,kind,subject) value (?,?,?,?,?,?,?,?,?,?,?)"
	_,err := common.Db.Exec(sqlStr,topic.Id,topic.Question,topic.A,topic.B,topic.C,topic.D,topic.Photo,topic.Answer,topic.Variety,topic.Kind,topic.Subject)
	if err != nil{
		log.Printf("MasterAddTopic Db err : %v\n",err)
		return err
	}
	return nil
}
//管理员删除文件
func MasterDeleteTopic(id int) error{
	sqlStr := "delete from question_bank where id = ? "
	_,err :=common.Db.Exec(sqlStr,id)
	if err!=nil{
		return err
	}
	return nil
}
//管理员查找文件
func MasterCheckTopic(question string) ([]*model.Topic,error){
	sqlStr := "select * from question_bank where question like ? "
	rows,err:= common.Db.Query(sqlStr,"%"+question+"%")
	if err != nil {
		log.Printf("MasterCheckTopic Db err : %v\n",err)
		return nil, err
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
//管理员按照分类查找题目
func MasterCheckByKindTopic(kind model.Topic)([]*model.Topic,error){
	sqlStr := "select * from question_bank where variety = ? and kind = ? and subject = ?"
	rows,err := common.Db.Query(sqlStr,kind.Variety,kind.Kind,kind.Subject)
	if err != nil{
		log.Printf("MasterCheckByKindTopic Db err : %v\n",err)
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

//管理员修改题目
func MasterUpdateTopic(topic model.Topic)(error){
	sqlStr := "update question_bank set question = ?,a = ?,b = ?,c = ?,d = ?,photo = ?,answer = ?,variety = ?,kind = ?,subject = ? where id = ?"
	_ , err:= common.Db.Exec(sqlStr,topic.Question,topic.A,topic.B,topic.C,topic.D,topic.Photo,topic.Answer,topic.Variety,topic.Kind,topic.Subject,topic.Id)
	if err != nil{
		log.Printf("MasterUpdateTopic Db err : %v\n",err)
		return err
	}
	return nil
}


//查询所有驾校
func MasterAllSchool()([]*model.DrvingSchool,error){
	sqlStr := "select *from drving_school"
	rows,err := common.Db.Query(sqlStr)
	if err!=nil{
		log.Printf("MasterAllSchool Db err : %v\n",err)
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

//管理员添加驾校
func MasterAddDrvingSchool(school model.DrvingSchool) error{
	sqlStr := "insert into drving_school(name,address,phone) value (?,?,?) "
	_,err := common.Db.Exec(sqlStr,school.Name,school.Address,school.Phone)
	if err != nil{
		log.Printf("MasterAddDrvingSchool Db err : %v\n",err)
		return err
	}
	return nil

}
//管理员删除驾校
func MasterDeleteDrvingSchool(id int)error{
	sqlStr := "delete from drving_school where id =?"
	_,err := common.Db.Exec(sqlStr,id)
	if err!=nil{
		log.Printf("MasterDeleteDrvingSchool Db err : %v\n",err)
		return err
	}
	return nil
}
//管理员修改驾校
func MasterUpdateDrvingSchool(school model.DrvingSchool)error{
	sqlStr := "update drving_school set name = ?,address = ?,phone = ? where id =?"
	_,err := common.Db.Exec(sqlStr,school.Name,school.Address,school.Phone,school.Id)
	if err != nil{
		log.Printf("MasterUpdateDrvingSchool Db err : %v\n",err)
		return err
	}
	return nil
}
//管理员根据id查询题目
func MasterCheckDrvingSchool(id int)model.DrvingSchool{
	sqlStr := "select * from drving_school where id = ?"
	row := common.Db.QueryRow(sqlStr,id)
	var school model.DrvingSchool
	row.Scan(&school.Id,&school.Name,&school.Address,&school.Phone)
	return school

}
func MasterAddVideo(video model.Video)error{
	sqlStr := "insert into videos(practicename,videoname) value (?,?)"
	_,err := common.Db.Exec(sqlStr,video.PracticeName,video.PracticeName)
	if err != nil{
		log.Printf("MasterAddVideo Db err : %v\n",err)
		return err
	}
	return nil
}
func CheckVideoNameByPracticeName(ptn string)string{
	sqlStr := "select videoname from videos where practicename = ?"
	row := common.Db.QueryRow(sqlStr,ptn)
	var videoName string
	row.Scan(&videoName)
	return videoName
}
func DeleteVideoByPracticeName(ptn string)bool{
	sqlStr := "delete from videos where practicename = ?"
	_,err := common.Db.Exec(sqlStr,ptn)
	if err != nil{
		log.Printf("Delete video failed err:%v",err)
		return false
	}
	return true
}

