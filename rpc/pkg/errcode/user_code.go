package errcode

/**
* @Author: super
* @Date: 2020-11-24 18:45
* @Description:
**/

var (
	ErrorUsernameUnRegister    = NewError(20060001, "用户未注册")
	ErrorUserIncorrectPassword = NewError(20060002, "用户密码错误")
	ErrorUserRegisterFail      = NewError(20060003, "用户注册失败")
	ErrorUserListFail          = NewError(20060004, "获取用户列表失败")
	ErrorDuplicateUsername     = NewError(20060005, "用户名已经注册")
	ErrorDuplicateMobile       = NewError(20060006, "手机号已经被占用")
)
