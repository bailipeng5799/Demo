package dao

import (
	"jxbdproject/model"
	"jxbdproject/common"
	"log"
)
//查找用户名和密码登陆注册时候用来验证
func CheckUserNameAndPassword(username string,password string )(*model.User,error){
	sqlStr := "select userid,username,password,email,name from users where username = ? and password = ? "
	row :=common.Db.QueryRow(sqlStr,username,password)
	user := &model.User{}
	err := row.Scan(&user.Id,&user.Username,&user.Password,&user.Email,&user.Name)
	return user,err
}
//查找账号注册时候检验是否可以使用此账号进行注册
func CheckUser(username string)(*model.User){
	sqlStr := "select userid,username,password,email,name from users where username = ?"
	row := common.Db.QueryRow(sqlStr,username)
	user := &model.User{}
	if err:=row.Scan(&user.Id,&user.Username,&user.Password,&user.Email,&user.Name);err!=nil{
		log.Printf("CheckUser Db err: %v\n",err)
		return nil
	}

	return user
}

func AddUser(user *model.User)(bool){
	sqlStr := "insert into users(username,password,email,name) values(?,?,?,?)"
	_,err := common.Db.Exec(sqlStr,user.Username,user.Password,user.Email,user.Name)
	if err != nil{
		log.Printf("AddUser Db err : %v\n",err)
		return false
	}
	return true
}
func CheckUserById(id int)(*model.User){
	sqlStr := "select userid,username,email,name from users where userid = ?"
	row := common.Db.QueryRow(sqlStr,id)
	user := &model.User{}
	if err := row.Scan(&user.Id,&user.Username,&user.Email,&user.Name);err!=nil{
		log.Printf("CheckUserById Db err : %v\n",err)
		return nil
	}
	return user
}