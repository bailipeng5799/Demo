package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)
func IndexHandler(w http.ResponseWriter,r *http.Request){
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo==""{
		pageNo="1"
	}
	//调用bookdao中获取分页的函数
	page,_ := dao.GetPageBooks(pageNo)
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w,page)
}

//获取所有图书GetPageBooks
func GetPageBooks(w http.ResponseWriter,r *http.Request){
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo==""{
		pageNo="1"
	}
	//调用bookdao中获取分页的函数
	page,_ := dao.GetPageBooks(pageNo)
	//解析模板 文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}
//func Getbooks(w http.ResponseWriter,r *http.Request){
//	//调用bookdao中的函数
//	books, _:=dao.GetBooks()
//	//解析模板 文件
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	t.Execute(w, books)
//}
////添加图书addbook
//func Addbook(w http.ResponseWriter,r *http.Request){
//	title := r.PostFormValue("title")
//	author := r.PostFormValue("author")
//	price := r.PostFormValue("price")
//	sales := r.PostFormValue("sales")
//	stock := r.PostFormValue("stock")
//	//将价格销量库存进行转换
//	fprice,_  := strconv.ParseFloat(price,64)
//	isales,_ := strconv.ParseInt(sales,10,0)
//	istock,_ := strconv.ParseInt(stock,10,0)
//	book:= &model.Book{
//		Title:title,
//		Author:author,
//		Price:fprice,
//		Sales:int(isales),
//		Stock:int(istock),
//		ImgPath:"static/img/default.jpg",
//	}
//	//调用bookdao中添加图书的函数
//	dao.AddBook(book)
//	//调用GetBooks处理器函数再次查询一次数据库
//	Getbooks(w,r)
//
//}

//Deletebook
func Deletebook(w http.ResponseWriter,r *http.Request){
	bookID :=r.FormValue("bookId")
	//调用bookdao中删除图书的函数
	dao.DeleteBook(bookID)
	GetPageBooks(w,r)

}
//更新或者添加图书的一个页面
func ToUpdateBookpage(w http.ResponseWriter,r *http.Request){
	//获取要更新的图书id
	bookID :=r.FormValue("bookId")
	//调用bookdao中获取图书的函数
	book,_ := dao.GetBookByID(bookID)
	if book.Id >0{
		//在更新图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w,book)
	}else{
		//在添加图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w,"")
	}

}

//UpdateBook 更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter,r *http.Request){
	bookID := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	//将价格销量库存进行转换
	fprice,_  := strconv.ParseFloat(price,64)
	isales,_ := strconv.ParseInt(sales,10,0)
	istock,_ := strconv.ParseInt(stock,10,0)
	ibookID,_ := strconv.ParseInt(bookID,10,0)
	book:= &model.Book{
		Id:	int(ibookID),
		Title:title,
		Author:author,
		Price:fprice,
		Sales:int(isales),
		Stock:int(istock),
		ImgPath:"static/img/default.jpg",
	}
	if book.Id>0{
		//调用bookdao中更新图书的函数
		dao.UpdateBook(book)
	}else{
		dao.AddBook(book)
	}
	//调用GetBooks处理器函数再次查询一次数据库
	GetPageBooks(w,r)
}
//获取带分页和价格的图书GetPageBooksByPrice
func GetPageBooksByPrice(w http.ResponseWriter,r *http.Request){
	//获取页码
	pageNo := r.FormValue("pageNo")
	//获取价格范围
	minPrice :=r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo==""{
		pageNo="1"
	}
	var page *model.Page
	if minPrice == ""&&maxPrice == ""{
		page,_ = dao.GetPageBooks(pageNo)
	}else {
		//调用bookdao中获取分页和价格范围的图书的的函数
		page,_ = dao.GetPageBooksByPrice(pageNo,minPrice,maxPrice)
		page.MaxPrice=maxPrice
		page.MinPrice=minPrice
	}
	//获取cookie
	cookie,_ := r.Cookie("user")
	if cookie!=nil{
		//获取cookie的value
		cookieValue := cookie.Value
		session := dao.GetsessionId(cookieValue)
		if session.UserID > 0{
				page.Status=true
				page.Username=session.Username
		}
	}
	//解析模板 文件
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}