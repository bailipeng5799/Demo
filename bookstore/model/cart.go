package model
//Cart 购物车结构体
type Cart struct {
	CartId string //购物车Id
	CartItems []*CartItem//购物项
	TotalCount int64//购物车中图书的总数量
	TotalAmount float64 // 购物车中图书的总金额，通过计算得到
	UserID int//购物车属于谁x
}

//GetTatalCount获取购物车中的总数量
func(cart *Cart) GetTatalCount () int64{
	var totalCount int64

	//遍历购物车中的购物项切片
	for _,v := range cart.CartItems{
		totalCount=totalCount+v.Count
	}
	return totalCount
}
//GetTotalAmount 获取购物车中的总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _,v := range cart.CartItems{
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}