package message

//确定消息类型
const (
	LoginMessageType            = "LoginMessage"
	LoginResponceMessageType    = "LoginResponMessage"
	RegisterMesssageType        = "RegisterMesssage"
	RegisterResMessageType      = "RegisterResMessage"
	NotifyUserStatusMessageType = "NotifyUserStatusMessage"
	SmsMessageType              = "SmsMessage"
)

//这里我们定义几个用户在线的状态常量
const (
	UserOnline = iota
	UserOffLine
	UserBusyStutas
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

// RegisterMesssages  定义用户注册信息结构体
type RegisterMesssages struct {
	User User `json:"user"` //类型就是User结构体
}

// RegisterResMessage 定义用户注册反馈信息结构体
type RegisterResMessage struct {
	Code  int    `json:"code"`  // 返回状态码 400 表示该用户已经占有 200表示注册成功
	Error string `json:"error"` // 返回错误信息
}

// NotifyUserStatusMessage 为了配合服务器端，推送用户状态变化的消息，定义该结构体
type NotifyUserStatusMessage struct {
	UserID     int `json:"userID"`     //用户id
	UserStatus int `json:"userStatus"` //用户状态
}

// SmsMessage 增加一个处理消息发送的结构体
type SmsMessage struct {
	Content string `json:"content"` //消息内容
	User           //匿名结构体，继承user.go的User结构体
}
