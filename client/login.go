//构建用户登陆操作
package main

import (
	"ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// login 写一个函数，完成登录功能
func login(userID int, userPwd string) (err error) {
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
	//fmt.Printf("客户端发送的消息内容为:%s", string(data))
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

	mes, err = ReadPackage(conn)
	if err != nil {
		fmt.Println("ReadPackage (conn) faild err=", err)
		return
	}
	//进行解包
	//将mes的Data部分进行反序列化
	var loginResMes message.LoginResponMessage
	err = json.Unmarshal([]byte(mes.MessageData), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功了")
	} else if loginResMes.Code == 500 {
		fmt.Println("登录失败", loginResMes.Error)
	}
	return
}
