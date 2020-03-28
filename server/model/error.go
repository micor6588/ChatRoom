package model

import "errors"

//根据业务逻辑自己定义一些错误
var (
	ERROR_USER_NOTEXITS = errors.New("用户不存在")
	ERROR_USER_EXITS    = errors.New("用户已经存在")
	ERROR_USER_PWD      = errors.New("用户密码不正确")
)
