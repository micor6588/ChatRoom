//主要功能：1.显示登录成功的页面
//2.保持和服务器通讯
//3.当读取服务器发送的消息后就会显示在页面

package process

import (
	"ChatRoom/client/utils"
	"fmt"
	"net"
	"os"
)

// ShowMenu 显示登录成功的页面
func ShowMenu() {
	fmt.Println("---------------恭喜xxx登录成功-----------------")
	fmt.Println("-------------1.显示用户在线用户列表--------------")
	fmt.Println("-----------------2.发送消息-----------------")
	fmt.Println("-----------------3.信息列表-----------------")
	fmt.Println("-----------------4.退出系统-----------------")
	fmt.Println("-----------------“请选择(1-4):”-----------------")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("---显示用户在线用户列表----")
	case 2:
		fmt.Println("--发送消息--")
	case 3:
		fmt.Println("--查看信息列表---")
	case 4:
		fmt.Println("--您已经推出系统--")
		os.Exit(0)
	default:
		fmt.Println("您的选项不正确，请重新输入")
	}

}

// ServerProcessMes 和服务器端保持通讯
func ServerProcessMes(conn net.Conn) {
	//创建一个Transfer实例
	tf := utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待服务器发送消息")
		mes, err := tf.ReadPackage()
		if err != nil {
			fmt.Println("tf.ReadPackage faild err=", err)
			return
		}
		//如果读取到消息，就要进行下一步的逻辑处理
		fmt.Printf("mes=%v\n", mes)
	}
}
