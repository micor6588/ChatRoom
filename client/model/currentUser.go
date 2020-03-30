package model

import (
	"ChatRoom/common/message"
	"net"
)

// CurrentUser 当前用户
type CurrentUser struct {
	Conn         net.Conn
	message.User //匿名结构体
}
