package main

import (
	"ChatRoom/server/model"
	_ "errors"
	"fmt"
	"net"
	"time"
)

//处理和客户端之间的通讯
func process(conn net.Conn) {
	//这里延时关闭conn
	defer conn.Close()
	//这里创建一个总控的实例
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器之间的协程出问题,err", err)
		return
	}
}

//这里我们编写一个函数完成对UserDao的初始化任务
func initUserDao() {
	//这里需要注意初始化顺序的问题
	//initPool,在initUserDao
	model.MyUserDao = model.NewUserDao(pool) //这里的pool是已经在redis.go里面定义的全局变量pool
}

func main() {
	//当服务器启动时，我们就初始化连接池
	initPool("127.0.0.1:6379", 16, 0, 300*time.Second)
	initUserDao()
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
