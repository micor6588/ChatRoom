package process

import (
	"ChatRoom/common/message"
	"ChatRoom/server/utils"
	"encoding/json"
	"fmt"
)

// SmsProcess 短消息处理
type SmsProcess struct {
}

// SendGroupMessage 发送群聊消息
func (pro *SmsProcess) SendGroupMessage(content string) (err error) {
	//创建一个mes
	var mes message.Message
	mes.MessageType = message.SmsMessageType
	//2.创建一个SmsMessage实例
	var smsMes message.SmsMessage
	smsMes.Content = content
	smsMes.UserID = CurrentUser.UserID
	smsMes.UserStatus = CurrentUser.UserStatus

	//3.序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("smsMes json.Marshal faild err=", err)
		return
	}
	mes.MessageData = string(data)
	//对mes进行再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.Marshal faild err=", err)
		return
	}
	//5.将序列化后的mes发送给服务器
	tf := &utils.Transfer{
		Conn: CurrentUser.Conn,
	}
	//发送群聊消息
	err = tf.WritePackage(data)
	if err != nil {
		fmt.Println("Send GroupMessage  faild err=", err)
		return
	}
	return

}
