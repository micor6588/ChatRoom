// process2 处理用户相关业务逻辑，比如登录，注册等...
package process2

import (
	"ChatRoom/common/message"
	"ChatRoom/server/model"
	"ChatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

// UserProcess 将方法关联到结构体当中
type UserProcess struct {
	//思考：需要哪些字段。主要考虑到自己关联的方法需要哪些字段
	Conn   net.Conn
	UserID int //表示该Conn是哪个用户
}

// ServerProcessRegist 编写一个ServerProcessRegist
// 功能：专门处理注册相关逻辑
func (pro *UserProcess) ServerProcessRegist(mes *message.Message) (err error) {
	//1.先从mes中取出mes.MessageData,并发序列化成RegisterMesssages
	var registMessage message.RegisterMesssages
	err = json.Unmarshal([]byte(mes.MessageData), &registMessage)
	if err != nil {
		fmt.Println("login Unmarshal faild err=", err)
		return
	}
	//1.先声明一个resMes
	var resMes message.Message
	resMes.MessageType = message.RegisterResMessageType
	//2.再声明一个registResMes,并完成赋值
	var registResMes message.RegisterResMessage
	//到数据库验证
	//使用model.MyUserDao到redis验证
	err = model.MyUserDao.RegistVerify(&registMessage.User)
	if err != nil {
		if err == model.ERROR_USER_EXITS {
			//不合法
			registResMes.Code = 505
			registResMes.Error = model.ERROR_USER_EXITS.Error()
			//我们在这里先测试成功，然后再返回具体的错误信息
		} else {
			registResMes.Code = 506
			registResMes.Error = "注册发生未知错误"
		}
	} else {
		registResMes.Code = 200
	}

	//3.将loginResMes进行序列化
	data, err := json.Marshal(registResMes)
	if err != nil {
		fmt.Println("loginResMes json Marshal faild err=", err)
		return
	}
	//4.将data赋值给resMes
	resMes.MessageData = string(data)
	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("resMes json Marshal faild err=", err)
		return
	}
	//6.将其封装到函数WritePackage当中
	//因为采用了分层的设计模式（MVC)，我们先创建一个Transfer实例然后读取
	tf := utils.Transfer{
		Conn: pro.Conn,
	}
	err = tf.WritePackage(data)
	return
}

// ServerProcessLogin 编写一个ServerProcessLogin
// 功能：专门处理登录相关逻辑
func (pro *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes中取出mes.MessageData,并发序列化成LoginMeassage
	var loginMessage message.LoginMessage
	err = json.Unmarshal([]byte(mes.MessageData), &loginMessage)
	if err != nil {
		fmt.Println("login Unmarshal faild err=", err)
		return
	}
	//1.先声明一个resMes
	var resMes message.Message
	resMes.MessageType = message.LoginResponceMessageType
	//2.再声明一个LoginResMes,并完成赋值
	var loginResMes message.LoginResponMessage
	//先将登录验证写死，之后改写为数据库验证
	//使用model.MyUserDao到redis验证
	user, err := model.MyUserDao.LoginVerify(loginMessage.UserID, loginMessage.UserPwd)
	if err != nil {
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在请注册之后再使用"
		//我们在这里先测试成功，然后再返回具体的错误信息
	} else {
		loginResMes.Code = 200
		//由于用户已经登录成功，于是将该登录成功的用户放到UserManger中
		//将登录成功的UserID赋值给pro
		pro.UserID = loginMessage.UserID
		userManager.AddOnlineUsers(pro)
		//将当前用户的UserID放到loginResponceMessage.UserID
		//遍历UserManager.OnlineUsers
		for id, _ := range userManager.OnlineUsers {
			loginResMes.UsersID = append(loginResMes.UsersID, id)
		}
		fmt.Println(user, "登录成功了")
	}

	/*
		//如股用户id=100,密码等于abcdef，就认定为合法
		if loginMessage.UserID == 100 && loginMessage.UserPwd == "abcdef" {
			//合法
			loginResMes.Code = 200
		} else {
			//不合法
			loginResMes.Code = 500
			loginResMes.Error = "该用户不存在请注册之后再使用"
		}
	*/
	//3.将loginResMes进行序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("loginResMes json Marshal faild err=", err)
		return
	}
	//4.将data赋值给resMes
	resMes.MessageData = string(data)
	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("resMes json Marshal faild err=", err)
		return
	}
	//6.将其封装到函数WritePackage当中
	//因为采用了分层的设计模式（MVC)，我们先创建一个Transfer实例然后读取
	tf := utils.Transfer{
		Conn: pro.Conn,
	}
	err = tf.WritePackage(data)
	return
}
