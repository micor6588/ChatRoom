//主要功能：1.显示登录成功的页面
//2.保持和服务器通讯
//3.当读取服务器发送的消息后就会显示在页面

package process

import (
	"ChatRoom/client/utils"
	"ChatRoom/common/message"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

// ShowMenu 显示登录成功的页面
func ShowMenu() {
	fmt.Println("---------------恭喜xxx登录成功-----------------")
	fmt.Println("-------------1.显示用户在线用户列表--------------")
	fmt.Println("----------------2.发送消息(群发)----------------")
	fmt.Println("-----------------3.点对点聊天(私聊)-----------------")
	fmt.Println("-----------------4.信息列表-----------------")
	fmt.Println("-----------------5.退出系统-----------------")
	fmt.Println("-----------------“请选择(1-5):”-----------------")
	var key int
	var content string
	//因为聊天必须使用SmsProcess实例，因此将其定义在外部
	smsProcess := &SmsProcess{}
	osIn := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("---显示用户在线用户列表----")
		ShowOnlineUsers()
		//TODO:群聊
	case 2:
		fmt.Println("--发送消息--")
		fmt.Println("请输入您想对大家说的话")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMessage(content)
		// TODO:私聊
	case 3:
		//显示在线用户
		ShowOnlineUsers()
		fmt.Println("请输入想私聊的用户ID")
		readline, _, err := osIn.ReadLine()
		if err != nil {
			fmt.Println("终端输入有误，请重新输入")
			return
		}
		sendToUserID, err := strconv.Atoi(string(readline))
		if err != nil {
			fmt.Println("输入用户的id错误，err", err)
			return
		}
		fmt.Println("请输入需要私聊的内容")
		sendContent, _, err := osIn.ReadLine()
		if err != nil {
			fmt.Println("终端输入有误，请重新输入")
			return
		}
		smsProcess.SendPrivateMessage(sendToUserID, string(sendContent))
	case 4:
		fmt.Println("--查看信息列表---")
	case 5:
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
		switch mes.MessageType {
		case message.NotifyUserStatusMessageType: //提示;有人已经上线
			//1.取出用户的NotifyUserStatusMessage
			var notifyUserStatusMessage message.NotifyUserStatusMessage
			err = json.Unmarshal([]byte(mes.MessageData), &notifyUserStatusMessage)
			if err != nil {
				fmt.Println("notifyUserStatusMessage  json.Unmarshal err=", err)
				return
			}
			//2.把该用户的信息和状态保持到客户端Map当中,Map格式：map[int]User.
			UpdataUserStatus(&notifyUserStatusMessage)
		case message.SmsMessageType: //有人群发消息
			OutPutGroupMessage(&mes)
		case message.SmsPrivateMessageType: //
			OutPutPrivateMessage(&mes)
		default:
			fmt.Println("服务器端返回了未知消息类型")

		}
		fmt.Printf("mes=%v\n", mes)
	}
}
