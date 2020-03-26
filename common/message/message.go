package message

//确定消息类型
const (
	LoginMessageType         = "LoginMessage"
	LoginResponceMessageType = "LoginResponMessage"
	RegisterMesssageType     = "RegisterMesssage"
)

// Message 定义消息结构体
type Message struct {
	MessageType string `json:"type"` //消息类型
	MessageData string `json:"data"` //消息的内容
}

// LoginMessage 定义登陆消息结构体
type LoginMessage struct {
	UserID   int    `json:"userId"`   //用户ID
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名

}

// LoginResponMessage 定义登陆反馈信息结构体
type LoginResponMessage struct {
	Code    int    `json:"code"` //返回状态码 500 表示用户未注册200表示登陆成功
	UsersID []int  //添加字段，保存用户id的切片
	Error   string `json:"error"` //返回错误信息
}

// RegisterResMessage 定义用户注册反馈信息结构体
type RegisterResMessage struct {
	Code  int    `json:"code"`  // 返回状态码 400 表示该用户已经占有 200表示注册成功
	Error string `json:"error"` // 返回错误信息
}
