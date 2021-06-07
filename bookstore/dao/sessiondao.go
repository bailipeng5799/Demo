package dao

import (
	"bookstore/model"
	"bookstore/views/utils"
	"net/http"
)

func Addsession(sess *model.Session) error{
	sqlStr := "insert into sessions values(?,?,?)"
	_,err := utils.Db.Exec(sqlStr,sess.SessionID,sess.Username,sess.UserID)
	if err!=nil{
		return err
	}
	return nil
}
func Deletesession(sessID string) error{
	sqlStr := "delete from sessions where session_id = ?"
	_,err := utils.Db.Exec(sqlStr,sessID)
	if err!=nil{
		return err
	}
	return nil
}
func GetsessionId(sessionId string) *model.Session{
		sqlStr := "select *from sessions where session_id = ?"
		row:= utils.Db.QueryRow(sqlStr,sessionId)
		 session :=&model.Session{}
		row.Scan(&session.SessionID,&session.Username,&session.UserID)
		return session
}
func Islogin(r *http.Request)(bool,string){

	//获取cookie  name 为user的cookie
	cookie,_ := r.Cookie("user")

	if cookie != nil{
		cookieValue := cookie.Value
		session:= GetsessionId(cookieValue)
		if session.UserID > 0 {
			return true,session.Username
		}
	}
	return false,""

}