package main

import (
	"bookstore/controller"
	"net/http"
)


func main(){
	//设置处理静态资源如cssjs文件
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/",http.StripPrefix("/pages/",http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/main",controller.IndexHandler)
	//去登录n)
	//去注销
	http.HandleFunc("/login",controller.Login)
	http.HandleFunc("/logout",controller.Logout)
	//去注册
	http.HandleFunc("/regist",controller.Regist)
	//通过ajex请求验证用户名是否可用
	http.HandleFunc("/checkUserName",controller.CheckUserName)
	//获取所有图书
	//http.HandleFunc("/getBooks",controller.Getbooks)
	//获取带分页的图书信息
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice",controller.GetPageBooksByPrice)
	//添加图书和下面的修改图书较为相似所以使用同一个处理器函数
	//http.HandleFunc("/addBook",controller.UpdateOrAddbook)
	//删除图书
	http.HandleFunc("/deleteBook",controller.Deletebook)
	//去更新图书的页面
	http.HandleFunc("/toUpdateBookpage",controller.ToUpdateBookpage)
	//更新或添加图书
	http.HandleFunc("/UpdateOrAddBook",controller.UpdateOrAddBook)
	http.ListenAndServe(":8080",nil)
}