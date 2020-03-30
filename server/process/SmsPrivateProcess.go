package process2

import (
	"ChatRoom/common/message"
	"ChatRoom/server/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

// SmsPrivateProcess 处理私聊消息结构体
type SmsPrivateProcess struct {
	//暂不需要任何字段
}

// SendPrivateMessage 处理私聊消息转发
func (pro *SmsPrivateProcess) SendPrivateMessage(mes *message.Message) (err error) {
	//功能:将私聊消息转发出去
	var smsPriMes message.SmsPrivateMessage
	err = json.Unmarshal([]byte(mes.MessageData), &smsPriMes)
	if err != nil {
		fmt.Println("smsMes son.Unmarshal faild err:", err)
		return
	}
	//遍历服务器端的MAP,将消息转发出去
	//1.取出mes当中的内容,为了后面避免将消息转发给在线特定用户
	privateUser, ok := userManager.OnlineUsers[smsPriMes.ChatUserID]
	if !ok {
		fmt.Println("查找不到想要私聊的用户,用户ID:", smsPriMes.ChatUserID)
		return errors.New("查找不到私聊用户")
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.Marshal faild err=", err)
		return
	}
	pro.SendPrivateMessageToOnlineUser(data, privateUser.Conn)
	return
}

// SendPrivateMessageToOnlineUser 将消息发送给每一个在线用户
func (pro *SmsPrivateProcess) SendPrivateMessageToOnlineUser(data []byte, conn net.Conn) {
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePackage(data)
	if err != nil {
		fmt.Println("转发消息失败，err", err)
		return
	}
}
