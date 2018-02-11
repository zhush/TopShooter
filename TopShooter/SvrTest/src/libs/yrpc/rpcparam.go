package yrpc

//远程调用请求的参数
type ReqParam struct {
	MethodName  string
	JsonContent string
}

//远程调用回应的参数
type RespParam struct {
	Result      bool   //处理是否成功
	HasResponse bool   //是否有返回值
	JsonContent string //返回的json,仅当HasResponse为true时有效
}
