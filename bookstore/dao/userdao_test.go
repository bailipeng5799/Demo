package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)
func TestMain(m *testing.M){
	fmt.Println("测试book中的方法")
	m.Run()
}



func TestUser(t *testing.T){
	//fmt.Println("测试userdao中的函数")
	//t.Run("验证用户名或密码：",testLogin)
	//t.Run("验证用户名：",testRegist)
	//t.Run("保存用户：",testSave)
	t.Run("测试获取分页的图书",testGetPageBooksByPrice)
}
func testLogin(t *testing.T){
	user, _ :=CheckUserNameAndPassword("admin","123456")
	fmt.Println("获取的用户信息是 ",user)
}
func testRegist(t *testing.T){
	user, _ :=CheckUserName("admin")
	fmt.Println("获取的用户信息是 ",user)
}
func testSave(t *testing.T){
	SaveUser("admin9","123456","admin@atguigu.com")
}
func TestBook(t *testing.T){
	//fmt.Println("测试bookdao中的相关函数")
	//t.Run("测试获取所有图书",testGetBooks)
	//t.Run("测试添加图书",testAddBook)
	//t.Run("测试删除图书",testDeletebook)
	//t.Run("测试查找图书",testGetbook)
	//t.Run("测试更新图书",TestUpdateBook)
	t.Run("测试更新图书",TestUpdateBook)
}
func testGetBooks(t *testing.T){
	books,_ := GetBooks()
	for k,v :=range books{
		fmt.Printf("第%v本图书的信息是：%v\n",k+1,v)
	}
}
func testAddBook(t *testing.T){
	book := &model.Book{
		Title:	"三国演义",
		Author:	"罗贯中",
		Price:	88.88,
		Sales:	100,
		Stock:	100,
		ImgPath:"/static/img/default.jpg",
	}
	//调用添加图书的函数
	AddBook(book)
}
func testDeletebook(t *testing.T) {
	//调用删除图书的函数
	DeleteBook("39")
}
func testGetbook(t *testing.T) {
	//调用删除图书的函数
	book,_ := GetBookByID("31")
	fmt.Println("获取的图书信息是：",book)
}
func TestUpdateBook(t *testing.T) {
	book := &model.Book{
		Id:  31,
		Title:	"三义",
		Author:	"罗中",
		Price:	888,
		Sales:	10,
		Stock:	10,
		ImgPath:"/static/img/default.jpg",
	}
	UpdateBook(book)
}
func Testsession(t *testing.T){
	fmt.Println("测试session相关函数")
	//t.Run("测试添加Session",TestAddsession)
	t.Run("测试删除session",TestDeletesession)
}
func TestAddsession(t *testing.T) {
	sess:=	&model.Session{
		SessionID:"0202020202",
		Username:"张某某",
		UserID:2,
		}
		Addsession(sess)
}
func TestDeletesession(t *testing.T) {
	Deletesession("0202020202")
}
func testGetPageBooks(t *testing.T){
	page,_ :=  GetPageBooks("1")
	fmt.Println("当前页是",page.PageNo)
	fmt.Println("总页数是",page.TotalPageNo)
	fmt.Println("总记录是",page.TotalRecord)
	fmt.Println("当前页中的图书有:")
	for _,v := range page.Books{
		fmt.Println("图书的信息是:",v)
	}

}
func testGetPageBooksByPrice(t *testing.T){
	page,_ :=  GetPageBooksByPrice("1","10","30")
	fmt.Println("当前页是",page.PageNo)
	fmt.Println("总页数是",page.TotalPageNo)
	fmt.Println("总记录是",page.TotalRecord)
	fmt.Println("当前页中的图书有:")
	for _,v := range page.Books{
		fmt.Println("图书的信息是:",v)
	}

}