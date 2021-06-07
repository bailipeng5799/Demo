package dao

import (
	"bookstore/model"
	"bookstore/views/utils"
)
//验证用户名和密码根据用户名和密码从数据库中查询一条记录
func CheckUserNameAndPassword(username string,password string) (*model.User,error){
	sqlStr := "select id,username,password,Email from users where username = ? and password =? "
	//执行
	row := utils.Db.QueryRow(sqlStr,username,password)
	user := &model.User{}
	row.Scan(&user.Id,&user.Username,&user.Password, &user.Email)
	return user,nil
}
//根据用户名和密码从数据库中查询一条记录
func CheckUserName(username string) (*model.User,error){
	sqlStr := "select id,username,password,email from users where username = ? "
	//执行
	row := utils.Db.QueryRow(sqlStr,username)
	user := &model.User{}
	row.Scan(&user.Id,&user.Username,&user.Password, &user.Email)
	return user,nil
}
//向数据库中加入信息
func SaveUser(username string,password string,email string) error{
	//写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_, err :=utils.Db.Exec(sqlStr,username,password,email)
	if err != nil{
		return err
	}
	return nil
}