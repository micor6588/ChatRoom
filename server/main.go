package main

import (
	"ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	_ "errors"
	"fmt"
	"io"
	"net"
)

// ReadPackage 读取客户端发过来的数据
func ReadPackage(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了connm,就不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		// fmt.Println("conn read datd faild,err=", err)
		// errors.New("read package hander faild")
		return
	}
	//根据buf[:4]转成一个uint32类型
	pkgLen := binary.BigEndian.Uint32(buf[:4])
	//一句pkgLen读取消息数据
	con, err := conn.Read(buf[:pkgLen])
	if con != int(pkgLen) || err != nil {
		fmt.Println("conn Read Package faild err=", err)
		return
	}
	//将pkgLen反序列化成->message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("PkgLen json.Unmarshal faild err=", err)
		return
	}
	return
}

// WritePackage 读取客户端发过来的数据
func WritePackage(conn net.Conn, data []byte) (err error) {
	//先发送一个长度到客户端
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//发送长度
	num, err := conn.Write(buf[0:4])
	//验证data长度是否发送成功
	if num != 4 || err != nil {
		fmt.Println("conn write(bytes) err=", err)
		return
	}
	//发送消息本身
	num, err = conn.Write(data)
	if num != int(pkgLen) || err != nil {
		fmt.Println("conn write(data) err=", err)
		return
	}
	return
}

// ServerProcessLogin 编写一个ServerProcessLogin
// 功能：专门处理登录相关逻辑
func ServerProcessLogin(conn net.Conn, mes *message.Message) (err error) {
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
	err = WritePackage(conn, data)
	return
}

// ServerProcessMessage 编写一个ServerProcessMessage
//功能：依据客户端发送的消息种类，决定调用哪个函数处理
func ServerProcessMessage(conn net.Conn, mes *message.Message) (err error) {
	switch mes.MessageType {
	case message.LoginMessageType:
		//处理登录的逻辑
		err = ServerProcessLogin(conn, mes)
	case message.RegisterMesssageType:
		//处理注册的逻辑
	default:
		fmt.Println("请正确输入消息类型，消息类型不存在，无法处理")
	}
	return
}

//处理和客户端之间的通讯
func process(conn net.Conn) {
	//这里延时关闭conn
	defer conn.Close()
	//循环读取客户端发送的消息
	for {
		//这里我们将读取数据包，直接封装成一个函数ReadPackage,返回一个（Message,error）
		fmt.Println("读取客户端发送的消息数据")
		mes, err := ReadPackage(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了，服务器也退出")
				return
			} else {
				fmt.Println("read package faild err=", err)
				return
			}
		}
		err = ServerProcessMessage(conn, &mes)
		if err != nil {
			return
		}
		//输出消息的内容
		fmt.Println("mes=", mes)
	}

}

func main() {
	//提示信息
	fmt.Println("服务器监听8687端口")
	listen, err := net.Listen("tcp", "127.0.0.1:8687")
	if err != nil {
		fmt.Println("net listen faild err=", err)
		return
	}
	defer listen.Close()
	//一旦监听成功，就等待客户端连接服务器
	for {
		fmt.Println("等待客户端连接服务器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("net listen faild err=", err)
			return
		}
		go process(conn)
	}
}
