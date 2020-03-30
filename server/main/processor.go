package main

import (
	"ChatRoom/common/message"
	process2 "ChatRoom/server/process"
	"ChatRoom/server/utils"
	"fmt"
	"io"
	"net"
)

// Processor 先创建一个处理器Processor的结构体
type Processor struct {
	Conn net.Conn
}

// ServerProcessMessage 编写一个ServerProcessMessage
//功能：依据客户端发送的消息种类，决定调用哪个函数处理
func (pro *Processor) ServerProcessMessage(mes *message.Message) (err error) {
	// //看看是否能接受到服务端的聊天信息
	// fmt.Printf("%s\n", mes)
	switch mes.MessageType {
	case message.LoginMessageType:
		//处理登录的逻辑
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: pro.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesssageType:
		//处理注册的逻辑
		up := &process2.UserProcess{
			Conn: pro.Conn,
		}
		err = up.ServerProcessRegist(mes)
	default:
		fmt.Println("请正确输入消息类型，消息类型不存在，无法处理")
	}
	return
}

//处理和客户端之间的通讯
func (pro *Processor) process2() (err error) {
	//循环读取客户端发送的消息
	for {
		//这里我们将读取数据包，直接封装成一个函数ReadPackage,返回一个（Message,error）
		fmt.Println("读取客户端发送的消息数据")
		//创建一个Transfer实例完成读包的任务
		tf := &utils.Transfer{
			Conn: pro.Conn,
		}
		mes, err := tf.ReadPackage()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了，服务器也退出")
				return err
			} else {
				fmt.Println("read package faild err=", err)
				return err
			}
		}
		err = pro.ServerProcessMessage(&mes)
		if err != nil {
			return err
		}
		//输出消息的内容
		fmt.Println("mes=", mes)
	}

}
