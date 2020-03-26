package main

import (
	"ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// ReadPackage 读取客户端发过来的数据
func ReadPackage(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了connm,就不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		// fmt.Println("conn read datd faild,err=", err)
		// errors.New("read package hander faild")
		return
	}
	//根据buf[:4]转成一个uint32类型
	pkgLen := binary.BigEndian.Uint32(buf[:4])
	//一句pkgLen读取消息数据
	con, err := conn.Read(buf[:pkgLen])
	if con != int(pkgLen) || err != nil {
		fmt.Println("conn Read Package faild err=", err)
		return
	}
	//将pkgLen反序列化成->message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("PkgLen json.Unmarshal faild err=", err)
		return
	}
	return
}

// WritePackage 读取客户端发过来的数据
func WritePackage(conn net.Conn, data []byte) (err error) {
	//先发送一个长度到客户端
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//发送长度
	num, err := conn.Write(buf[0:4])
	//验证data长度是否发送成功
	if num != 4 || err != nil {
		fmt.Println("conn write(bytes) err=", err)
		return
	}
	//发送消息本身
	num, err = conn.Write(data)
	if num != int(pkgLen) || err != nil {
		fmt.Println("conn write(data) err=", err)
		return
	}
	return
}
