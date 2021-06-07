package dao

import (
	"bookstore/model"
	"bookstore/views/utils"
)

func AddCartItem(cartItem *model.CartItem)error{
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values (?,?,?,?)"
	_,err := utils.Db.Exec(sqlStr,cartItem.Count,cartItem.GetAmount(),cartItem.Book.Id,cartItem.CartId)
	if err != nil{
		return err
	}
	return nil
}
//GetCartItemByBookId根据图书的id获取对应的购物项，因为重复添加时需要判断是否存在
// 如果存在只需要将数量进行操作
func GetCartItemByBookId(bookid int)(*model.CartItem,error){
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id= ?"
	row := utils.Db.QueryRow(sqlStr,bookid)
	//创建cartItem
	cartItem := &model.CartItem{}
	err:= row.Scan(&cartItem.CartItemId,&cartItem.Count,&cartItem.Amount,&cartItem.CartId)
	if err!=nil{
		return nil,err
	}
	return cartItem,nil
}
//GetCartItemsByCartId根据购物车的id来获取所有此购物车对应的所有的购物项
func GetCartItemsByCartId(cartid string)([]*model.CartItem,error){
	sqlStr := "select id,count,amount,cart_id from cart_items where cart_id= ?"
	rows,err :=utils.Db.Query(sqlStr,cartid)
	if err!=nil{
		return nil,err
	}
	var cartitems []*model.CartItem

	for rows.Next(){
		cartItem := &model.CartItem{}
		err2:= rows.Scan(&cartItem.CartItemId,&cartItem.Count,&cartItem.Amount,&cartItem.CartId)
		if err2!=nil{
			return nil,err2
		}
		cartitems=append(cartitems,cartItem)
	}
	return cartitems,nil
}