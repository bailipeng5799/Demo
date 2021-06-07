package model
//Page结构
type Page struct {
	Books []*Book//每页查询的图书存放的切片
	PageNo int64//当前第几页
	PageSize int64//每页显示的条数
	TotalPageNo int64//总页数，通过计算得到
	TotalRecord int64//总记录数，书本总数，通过查询数据库得到
	MaxPrice string // 查询 价格时候的最大值
	MinPrice string
	Status   bool	//判断登录的状态
	Username string//登录用户的用户名

}

//ISHasPrev 判断是够有上一页
func (p *Page) IsHanPrev() bool{
	return p.PageNo>1
}
//IsHasNext判断是否有下一页
func (p *Page) IsHanNext() bool{
	return p.PageNo < p.TotalPageNo
}
//GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int64{
	if p.IsHanPrev(){
		return p.PageNo-1
	}else{
		return p.PageNo
	}
}
//GetNextPageNo 获取上一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHanNext(){
		return p.PageNo+1
	}else{
		return p.PageNo
	}
}