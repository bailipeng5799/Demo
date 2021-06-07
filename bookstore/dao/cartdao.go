package dao

import (
	"bookstore/model"
	"bookstore/views/utils"
)

//AdddCart 向购物车表中插入购物车
func AddCart(cart *model.Cart)error {
	sqlStr := "insert into carts  (id,total_count,total_amount,user_id) values (?,?,?,?)"
	_,err := utils.Db.Exec(sqlStr,cart.CartId,cart.GetTatalCount(),cart.GetTotalAmount(),cart.UserID)
	if err!=nil{
		return nil
	}
	//获取购物车中的所有购物项
	cartitems := cart.CartItems
	for _,cartitem := range cartitems{
		//将每一个购物项插入到购物项所对应的表中
		AddCartItem(cartitem)
	}
	return nil
}
//根据userid拿到购物车的消息
func GetCartByUserId(userId int)(*model.Cart,error){
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id = ? "
	//执行sql
	row := utils.Db.QueryRow(sqlStr,userId)
	//创建一个购物车
	cart := &model.Cart{}
	err :=row.Scan(&cart.CartId,&cart.TotalCount,&cart.TotalAmount,&cart.UserID)
	if err!= nil {
		return nil,err
	}
	//根据当前购物车id获取当前购物车中的所有购物项
	cartitems, _ :=GetCartItemsByCartId(cart.CartId)
	cart.CartItems=cartitems
	return cart,nil
}
