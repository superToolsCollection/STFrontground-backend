package errcode

/**
* @Author: super
* @Date: 2020-09-18 07:53
* @Description: 统一错误代码
**/

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(10000000, "服务内部错误")
	InvalidParams = NewError(10000001, "入参错误")
	NotFound      = NewError(10000002, "找不到")

	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")

	JsonUnMarshalError  = NewError(10000008, "json序列化失败")
	NewHttpRequestError = NewError(10000009, "新建http请求失败")
	HttpBadRequestError = NewError(10000011, "http异常请求")
	ReadRequestError    = NewError(10000012, "读取http请求返回数据异常")

	UpdateTokenError = NewError(10000013, "更新Token失败")
)
