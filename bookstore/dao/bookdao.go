package dao

import (
	"bookstore/model"
	"bookstore/views/utils"
	"strconv"
)

//GetBooks 获取数据库中所有图书
func GetBooks()([]*model.Book,error) {
	//写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	//执行sqlStr
	rows, err := utils.Db.Query(sqlStr)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.Id,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	return books,nil
}

//addbooks 向数据库中添加一本图书
func AddBook(b *model.Book) error{
	//写sql语句
	sqlStr:="insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	//执行
	_ ,err := utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ImgPath)
	if err!=nil{
		return err
	}
	return nil
}
//deletebook 根据图书的id删除数据库中的图书
func DeleteBook(bookID string) error{
	//写sql语句
	sqlStr := "delete from books where id=?"
	_,err := utils.Db.Exec(sqlStr,bookID)
	if err!=nil{
		return err
	}
	return nil
}

func GetBookByID(bookID string)(*model.Book, error){
	sqlStr :="select id,title,author,price,sales,stock,img_path from books where id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr,bookID)
	//创建一个book
	book :=&model.Book{}
	//为book中的字段赋值
	row.Scan(&book.Id,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
	return book,nil
}
//更新图书根据图书的id
func UpdateBook(b *model.Book) error{
	sqlStr := "update books set title =?,author =?,price =?,sales =?,stock =? where id =?"
	//执行
	 _,err := utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.Id)
	if err!=nil{
		return err
	}
	return nil
}
//GetPageBooks获取带分页的图书信息
func GetPageBooks(pageNo string)(*model.Page,error){
	//将页码转换为int64
	iPageNo,_:= strconv.ParseInt(pageNo,10,64)
	//获取数据库中图书的总数
	sqlStr := "select count(*) from books"
	//设置一个变量接收总记录数
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//设置每一页显示4条记录
	var pageSize int64
	pageSize=4
	//一共有多少页
	var totalPageNo int64
	if totalRecord%pageSize==0{
		totalPageNo=totalRecord/pageSize
	}else{
		totalPageNo=totalRecord/pageSize+1
	}
	//获取当前页中的图书
	sqlStr2 :="select id,title,author,price,sales,stock,img_path from books limit ?,? "
	//执行
	rows,_ :=utils.Db.Query(sqlStr2,(iPageNo-1)*pageSize,pageSize)

	var books []*model.Book
	for rows.Next(){
		book := &model.Book{}
		rows.Scan(&book.Id,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		books =append(books,book)
	}
	page := &model.Page{
		Books:books,
		PageNo:iPageNo,
		PageSize:pageSize,
		TotalPageNo:totalPageNo,
		TotalRecord:totalRecord,
	}
	return page,nil
}
//GetPageBooks获取带分页的图书信息
func GetPageBooksByPrice(pageNo string,minPrice string,maxPrice string)(*model.Page,error){
	//将页码转换为int64
	iPageNo,_:= strconv.ParseInt(pageNo,10,64)
	//获取数据库中图书的总数
	sqlStr := "select count(*) from books where price between ? and ?"
	//设置一个变量接收总记录数
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr,minPrice,maxPrice)
	row.Scan(&totalRecord)
	//设置每一页显示4条记录
	var pageSize int64 = 4
	//一共有多少页
	var totalPageNo int64
	if totalRecord%pageSize==0{
		totalPageNo=totalRecord/pageSize
	}else{
		totalPageNo=totalRecord/pageSize+1
	}
	//获取当前页中的图书
	sqlStr2 :="select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,? "
	//执行
	rows,_ :=utils.Db.Query(sqlStr2,minPrice,maxPrice,(iPageNo-1)*pageSize,pageSize)

	var books []*model.Book
	for rows.Next(){
		book := &model.Book{}
		rows.Scan(&book.Id,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		books =append(books,book)
	}
	page := &model.Page{
		Books:books,
		PageNo:iPageNo,
		PageSize:pageSize,
		TotalPageNo:totalPageNo,
		TotalRecord:totalRecord,
	}
	return page,nil
}