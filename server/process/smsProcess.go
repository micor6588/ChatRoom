//处理短消息相关的,比如私聊，群发...
package process2

import (
	"ChatRoom/common/message"
	"ChatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

// SmsProcess 处理群聊消息转发
type SmsProcess struct {
}

// SendGroupMessage 发送群聊消息
func (pro *SmsProcess) SendGroupMessage(mes *message.Message) (err error) {
	//遍历服务器端的MAP,将消息转发出去
	//1.取出mes当中的内容,为了后面避免将消息转发给自己
	var smsMes message.SmsMessage
	err = json.Unmarshal([]byte(mes.MessageData), &smsMes)
	if err != nil {
		fmt.Println("smsMes son.Unmarshal faild err:", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.Marshal faild err=", err)
		return
	}
	//过滤掉消息人本身，并转发消息
	for id, up := range userManager.OnlineUsers {
		//这里还需要过滤掉自己，不要将消息发送给自己
		if id == smsMes.UserID {
			continue
		}

		pro.SendMessageToOnlineUser(data, up.Conn)
	}
	return
}

// SendMessageToOnlineUser 将消息发送给每一个在线用户
func (pro *SmsProcess) SendMessageToOnlineUser(data []byte, conn net.Conn) {
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
