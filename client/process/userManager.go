package process

import (
	"ChatRoom/client/model"
	"ChatRoom/common/message"
	"fmt"
)

// OnlineUsers 客户端需要维护的Map
var OnlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurrentUser model.CurrentUser //我们在用户登录成功后，对CurrentUser进行初始化

// ShowOnlineUsers 在客户端显示当前在线用户
func ShowOnlineUsers() {
	fmt.Println("当前在线用户有:")
	for id, _ := range OnlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

// UpdataUserStatus 编写一个方法，处理服务器返回的用户状态
func UpdataUserStatus(notifyUserStatusMessage *message.NotifyUserStatusMessage) {
	//适当优化,先观察该用户ID是否处于在线或则忙碌状态
	user, ok := OnlineUsers[notifyUserStatusMessage.UserID]
	if !ok { //原来没有
		user = &message.User{
			UserID: notifyUserStatusMessage.UserID,
		}
	}
	user.UserStatus = notifyUserStatusMessage.UserStatus

	OnlineUsers[notifyUserStatusMessage.UserID] = user
	ShowOnlineUsers()
}
