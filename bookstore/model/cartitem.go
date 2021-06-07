package model
//CartItem 购物项结构体
type CartItem struct {
	CartItemId int64  //购物项的id
	Book  *Book		//购物项中的图书信息
	Count int64		//购物项中的图书数量
	Amount float64	//购物项中图书的金额小计，通过数量*图书中的价格
	CartId string //当前购物项属于哪一个购物车的id
}
// 获取购物项中图书的金额小计，由图书的价格和图书的数量计算得到
func (cartItem *CartItem)GetAmount() float64{
	price := cartItem.Book.Price
	return float64(cartItem.Count)*price
}
