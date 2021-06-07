package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/views/utils"
	"html/template"
	"net/http"
)

//Login 处理用户登录的函数

func Login(w http.ResponseWriter,r *http.Request){

	flag,_:=dao.Islogin(r)
	if flag{
		//已经登录
		//去首页
		GetPageBooksByPrice(w,r)
	}else {
		//获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		//调用userdao中验证用户名和密码的方法
		user, _ := dao.CheckUserNameAndPassword(username, password)
		if user.Id > 0 {
			//用户名或密码正确
			//使用uuid作为生成的sessionid
			uuid := utils.CreateUUID()
			//创建session
			sess := &model.Session{
				SessionID: uuid,
				Username:  user.Username,
				UserID:    user.Id,
			}
			//将session保存到数据库里
			dao.Addsession(sess)
			//创建cookie让他与session相关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			//用户名或密码不正确
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用户名或密码不正确！")
		}
	}
}
//用户注册函数
func Regist(w http.ResponseWriter,r *http.Request){
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email :=r.PostFormValue("email")
	//调用userdao中验证用户名和密码的方法
	user,_ := dao.CheckUserName(username)
	if user.Id > 0{
		//用户名已经存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w,"用户名已存在！")
	}else{
		//用户名可用将用户名传入数据库中
		dao.SaveUser(username,password,email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w,"")

	}
}
//CheckUserName 通过发送ajax验证用户名是否可用

func CheckUserName(w http.ResponseWriter,r *http.Request){
	//获取用户输入的用户名
	username := r.PostFormValue("username")
	user,_ := dao.CheckUserName(username)
	if user.Id > 0{
		//用户名已经存在
	w.Write([]byte("用户名已经存在"))
	}else{
		//用户名可用
	w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}

func Logout(w http.ResponseWriter,r *http.Request){
	cookie,_ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中对应的Session
		dao.Deletesession(cookieValue)
		//设置cookie失效
		cookie.MaxAge=-1
		//将修改之后的cookie发送给浏览器
		http.SetCookie(w,cookie)
	}
	//去首页
	GetPageBooksByPrice(w,r)
}
