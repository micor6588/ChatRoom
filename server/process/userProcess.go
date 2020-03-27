package process2

import (
	"ChatRoom/common/message"
	"ChatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

// UserProcess 将方法关联到结构体当中
type UserProcess struct {
	//思考：需要哪些字段。主要考虑到自己关联的方法需要哪些字段
	Conn net.Conn
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
	//如股用户id=100,密码等于abcdef，就认定为合法
	if loginMessage.UserID == 100 && loginMessage.UserPwd == "abcdef" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在请注册之后再使用"
	}
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
