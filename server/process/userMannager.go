//功能：用于显示个用户在线列表
package process2

import "fmt"

// UserManager 用户管理的结构体
//由于在很多的地方都要使用UserManager这个结构体，所以将其定义为全局的变量
type UserManager struct {
	OnlineUsers map[int]*UserProcess
}

var (
	userManager *UserManager
)

//完成对UserManger的初始化工作
func init() {
	userManager = &UserManager{
		OnlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// AddOnlineUsers 完成对OnlineUsers的添加
func (man *UserManager) AddOnlineUsers(up *UserProcess) {
	man.OnlineUsers[up.UserID] = up
}

// DelectOnlineUsers 完成对OnlineUsers的删除
func (man *UserManager) DelectOnlineUsers(userID int) {
	delete(man.OnlineUsers, userID)
}

// GetAllOnlineUser 返回当前所有的在线用户
func (man *UserManager) GetAllOnlineUser() map[int]*UserProcess {
	return man.OnlineUsers
}

// GetAllOnlineUserByID 根据UerID返回对应的值
func (man *UserManager) GetAllOnlineUserByID(userID int) (up *UserProcess, err error) {
	//如何从map当中取出一个值,带检测方式
	up, ok := man.OnlineUsers[userID]
	if !ok { //说明你查找的用户当前不在线
		err = fmt.Errorf("用户:%d,不在线", userID)
		return

	}
	return
}
