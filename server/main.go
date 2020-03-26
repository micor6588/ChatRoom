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
