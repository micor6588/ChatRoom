package message

//定义一个用户的结构体
type User struct {
	//确定字段信息
	//为了序列化和反序列化成功
	//用户信息的json字符串的key和结构体的字段对应的tag名字必须一致
	UserID     int    `json:"userId"`     //用户ID
	UserPwd    string `json:"userPwd"`    //用户密码
	UserName   string `json:"userName"`   //用户名
	UserStatus int    `json:"userStatus"` //用户状态...
}
