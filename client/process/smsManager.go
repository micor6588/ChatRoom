package process

import (
	"ChatRoom/common/message"
	"encoding/json"
	"fmt"
)

/********************************
     处理群聊和点对点聊天
*********************************/

// OutPutGroupMessage 显示群聊的转发消息
func OutPutGroupMessage(mes *message.Message) { //mes一定是smsMessage
	//将其显示即可
	//1.反序列化mes.MessageData
	var smsMes message.SmsMessage
	err := json.Unmarshal([]byte(mes.MessageData), &smsMes)
	if err != nil {
		fmt.Println("smsMes json.Unmarshal err:", err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户id:\t%d,对您说:\t%s", smsMes.UserID, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}

// OutPutGroupMessage 显示群聊的转发消息
func OutPutPrivateMessage(mes *message.Message) { //mes一定是smsMessage
	//将其显示即可
	//1.反序列化mes.MessageData
	var smsMes message.SmsPrivateMessage
	err := json.Unmarshal([]byte(mes.MessageData), &smsMes)
	if err != nil {
		fmt.Println("smsMes json.Unmarshal err:", err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户id:\t%d,对大家说:\t%s", smsMes.UserID, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
