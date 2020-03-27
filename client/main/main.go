package main

import (
	"ChatRoom/client/process"
	"fmt"
	"os"
)

var userID int     //用户的账号
var userPwd string //用户密码

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

/*
	//依据用户的输入，显示新的信息
	if key == 1 {
		//说明用户要登陆
		fmt.Println("请输入用户的ID")
		fmt.Scanf("%d\n", &userID)
		fmt.Println("请输入您的用户密码：")
		fmt.Scanf("%s\n", &userPwd)

		//先把登录的函数，写到另外一个文件，比如login.go
		//这里我们会重新调用
		//login(userID, userPwd)
		// if err != nil {
		// 	fmt.Println("登陆失败")
		// } else {
		// 	fmt.Println("登陆成功")
		// }

		//因为使用了新的结构

	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑")
	}
*/
