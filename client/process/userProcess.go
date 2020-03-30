package process

import (
	"ChatRoom/client/utils"
	"ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// UserProcess 给关联一个用户登录的结构体
type UserProcess struct {
	//暂时不需要任何字段
}

// Regist 写一个函数，完成登录功能
func (pro *UserProcess) Regist(userID int, userPwd string, userName string) (err error) {
	//下一个就要开始定协议
	fmt.Printf(" useID=%d  userPwd=%s userName=%s\n", userID, userPwd, userName)
	//1.链接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8687")
	if err != nil {
		fmt.Println("net.Dtal faild err=", err)
		return
	}
	//延时关闭
	defer conn.Close()
	//2.准备通过conn发消息给服务器
	var mes message.Message
	mes.MessageType = message.RegisterMesssageType
	//3.创建一个LoginMeaage结构体
	var registMessage message.RegisterMesssages
	registMessage.User.UserID = userID
	registMessage.User.UserPwd = userPwd
	registMessage.User.UserName = userName

	//4.将registMessage序列化
	data, err := json.Marshal(registMessage)
	if err != nil {
		fmt.Println("loginmessage json Marshar faild err=", err)
		return
	}
	//5.把data赋值给mes.MessageData字段
	mes.MessageData = string(data)
	//6.将mes进行序列化操作
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json Marshar faild err=", err)
		return
	}

	//这里还需要处理服务器返回的消息
	//创建一个TranSfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器
	err = tf.WritePackage(data)
	if err != nil {
		fmt.Println("注册信息发送失败，Package (conn) faild err=", err)
		return
	}
	mes, err = tf.ReadPackage()
	if err != nil {
		fmt.Println("ReadPackage (conn) faild err=", err)
		return
	}

	//进行解包
	//将mes的Data部分进行反序列化
	var registResMes message.RegisterResMessage
	err = json.Unmarshal([]byte(mes.MessageData), &registResMes)
	if registResMes.Code == 200 {
		fmt.Println("恭喜您，注册成功了")
		os.Exit(0)

	} else {
		fmt.Println("注册失败了", registResMes.Error)
		os.Exit(0)
	}
	return
}

// Login 写一个函数，完成登录功能
func (pro *UserProcess) Login(userID int, userPwd string) (err error) {
	//下一个就要开始定协议
	fmt.Printf(" useID=%d  userPwd=%s", userID, userPwd)
	//1.链接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8687")
	if err != nil {
		fmt.Println("net.Dtal faild err=", err)
		return
	}
	//延时关闭
	defer conn.Close()
	//2.准备通过conn发消息给服务器
	var mes message.Message
	mes.MessageType = message.LoginMessageType
	//3.创建一个LoginMeaage结构体
	var loginMessage message.LoginMessage
	loginMessage.UserID = userID
	loginMessage.UserPwd = userPwd
	//4.将loginmessage序列化
	data, err := json.Marshal(loginMessage)
	if err != nil {
		fmt.Println("loginmessage json Marshar faild err=", err)
		return
	}
	//5.把data赋值给mes.MessageData字段
	mes.MessageData = string(data)
	//6.将mes进行序列化操作
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json Marshar faild err=", err)
		return
	}
	//7.到这个时候，data就是我们将要发送的数据，为了防止丢包，需要验证数据长度
	//7.1 先把data的数据长度，发送给服务器
	//先获取到data的长度—>转成一个表示长度的byte的切片
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//发送长度
	num, err := conn.Write(buf[0:4])
	if err != nil {
		fmt.Println("length send err=", err)
		return
	}
	//验证data长度是否发送成功
	if num != 4 || err != nil {
		fmt.Println("conn write(bytes) err=", err)
		return
	}
	fmt.Printf("客户端发送的消息长度为:%d,消息内容为:%s", len(data), string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn write(data) err=", err)
		return
	}
	//休眠20秒
	// time.Sleep(20 * time.Second)
	// fmt.Println()
	// fmt.Println("休眠了20秒")
	//这里还需要处理服务器返回的消息
	//创建一个TranSfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPackage()
	if err != nil {
		fmt.Println("ReadPackage (conn) faild err=", err)
		return
	}
	//进行解包
	//将mes的Data部分进行反序列化
	var loginResMes message.LoginResponMessage
	err = json.Unmarshal([]byte(mes.MessageData), &loginResMes)
	if loginResMes.Code == 200 {
		// fmt.Println("登录成功了")
		//初始化CurrentUser
		CurrentUser.Conn = conn
		CurrentUser.User.UserID = userID
		CurrentUser.User.UserStatus = message.UserOnline
		//可以显示当前用户的在线列表,遍历loginResponceMessage.UserID
		fmt.Println("当前用户的列表如下")
		for _, value := range loginResMes.UsersID {
			//如果不显示自己在线可以执行以下代码
			if value == userID {
				continue
			}
			fmt.Println("用户ID:\t", value)
			//完成客户端的OnlineUsers的初始化
			user := &message.User{
				UserID:     value,
				UserStatus: message.UserOnline,
			}
			OnlineUsers[value] = user
		}
		fmt.Print("\n\n")
		//这里我们还需要启动一个携程
		//该协程保持与服务器的连接，如果服务器有数据，就推送给客户端
		//则接收并显示在客户端的终端上面。
		go ServerProcessMes(conn)
		//1.显示我们登录成功的菜单，循环显示
		for {
			ShowMenu()
		}

	} else if loginResMes.Code == 500 {
		fmt.Println("登录失败", loginResMes.Error)
	}
	return
}
