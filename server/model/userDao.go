package model

import (
	"ChatRoom/common/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//我们在服务器启动后，就立马启动一个UserDao实例
//把它做成一个全局的变量，在需要redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

// UserDao 定义一个UserDao的结构体
type UserDao struct {
	pool *redis.Pool
}

// NewUserDao 使用工厂模式创建一个UserDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//思考UserDao需要提供哪些方法给我们
//完成对User的各种操作(增删该查)
func (dao *UserDao) getUserByID(conn redis.Conn, id int) (user *User, err error) {
	//通过给定的ID，去Redis里面查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误提示;
		if err == redis.ErrNil { //表示在users哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXITS
		}
		return
	}
	user = &User{}
	//这里需要把res反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("res json Unmarshal faild err", err)
		return
	}
	return
}

// LoginVerify 完成登录校验
//1.LoginVerify完成对用户的验证
//2.如果用户id和pwd都正确就返回一个User实例
//3.如果用户id和pwd有错误就返回错误的信息
func (dao *UserDao) LoginVerify(userID int, userPwd string) (user *User, err error) {
	//先从UserDao链接池当中取出一根链接
	conn := dao.pool.Get()
	defer conn.Close()
	user, err = dao.getUserByID(conn, userID)
	if err != nil {
		return
	}
	//这个时候证明用户获取到了
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return

}

// RegistVerify 完成注册校验
//1.RegistVerify完成对用户的验证
//2.如果用户id和pwd都正确就返回一个User实例
//3.如果用户id和pwd有错误就返回错误的信息
func (dao *UserDao) RegistVerify(user *message.User) (err error) {
	//先从UserDao链接池当中取出一根链接
	conn := dao.pool.Get()
	defer conn.Close()
	_, err = dao.getUserByID(conn, user.UserID)
	if err == nil {
		err = ERROR_USER_EXITS //用户已经存在
		return
	}
	//这个时候说明此用户ID还没有被注册
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("user json.Marshal faild err=", err)
		return
	}
	//没有序列化错误，就写入数据库
	_, err = conn.Do("HSet", "users", user.UserID, string(data))
	if err != nil {
		fmt.Println("user 保存注册用户 faild err=", err)
		return
	}
	return
}
