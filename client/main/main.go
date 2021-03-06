package main

import (
	"ChatRoom/client/process"
	"fmt"
	"os"
)

var userID int      //用户的账号
var userPwd string  //用户密码
var userName string //用户名

func main() {
	//接收用户选择
	var key int
	// 判断是否继续显示菜单
	var loop = true
	for loop {
		fmt.Println("------------------欢迎登陆多人聊天系统--------------")
		fmt.Println("\t\t\t               1.登陆聊天系统")
		fmt.Println("\t\t\t               2.注册用户")
		fmt.Println("\t\t\t               3.退出系统")
		fmt.Println("\t\t\t 请选择(1--3)")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			//说明用户要登陆
			fmt.Println("请输入用户的ID")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入您的用户密码：")
			fmt.Scanf("%s\n", &userPwd)

			//完成登录
			//1.创建一个UserProcess结构体
			up := &process.UserProcess{}
			up.Login(userID, userPwd)
			//loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入新用户的ID")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入您设置用户的密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入您设置用户的名字（nickName）：")
			fmt.Scanf("%s\n", &userName)
			//完成登录
			//1.创建一个UserProcess结构体
			//调用UserProcess.go里面的Regist函数完成注册
			up := &process.UserProcess{}
			up.Regist(userID, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("退出聊天室")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
	}

}
