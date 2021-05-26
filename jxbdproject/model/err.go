package model

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}
type ErrResponse struct{
	HttpSC int
	Error Err
}
var (
	//请求体不符合
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC:400,Error:Err{Error:"Request body is not corret",ErrorCode:"001"}}
	//用户验证失败
	ErrorNotAuthUser = ErrResponse{HttpSC: 401,Error:Err{Error:"User authentication failed.",ErrorCode:"002"}}
	//数据库失败
	ErrorDBError = ErrResponse{HttpSC:500,Error:Err{Error: "Db failed",ErrorCode:"003"}}
	//用户注册失败
	ErrorRegistFailed = ErrResponse{HttpSC:500,Error:Err{Error: "User registration failed",ErrorCode:"004"}}
	//服务器问题
	ErrorInternalFaults = ErrResponse{HttpSC:500,Error: Err{Error:"Internal service error",ErrorCode:"005"}}
)