package model

const Video_Dir  ="G:/goproject/src/jxbdproject/videos/"
//上传视频的信息
type Video struct{
	PracticeName string `json:"practicename"`
	VideoName string	`json:"videoname"`
}
//由于gin框架中特殊的参数绑定而需要的结构体
type Common struct{
	Strings string   `json:"strings"`
	Id int `json:"id"`
	Userid int `json:"userid"`
	Topicid int `json:"topicid"`
}
//关于题目的结构体
type Topic struct{
	Id int `json:"id"`//题号
	Question string`json:"question"`//问题本身
	A string `json:"a"`//A选项
	B string `json:"b"`//B选项
	C string `json:"c"`//C选项
	D string `json:"d"`//D选项
	Photo string `json:"photo"`//所要添加的图片地址
	Answer string `json:"answer"  `//题目答案
	Variety string `json:"variety"`//题目属于哪一个种内例如手势图，等
	Kind  int 	`json:"kind"` //题目属于判断题还是选择题 0 代表选择题 1 代表判断题
	Subject string `json:"subject"`//题目属于科目几
}
//关于驾校的结构体
type  DrvingSchool struct{
	Id int `json:"id"`//驾校编号
	Name string	`json:"name"`//驾校名称
	Address string `json:"address"`//驾校地址
	Phone string `json:"phone"`//驾校电话
}
