package utils

import (
	"ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// Transfer 将这些方法关联到结构体当中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //这是传输时使用到的缓冲
}

// ReadPackage 读取客户端发过来的数据
func (tran *Transfer) ReadPackage() (mes message.Message, err error) {
	// buf := make([]byte, 8096)
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了connm,就不会阻塞
	_, err = tran.Conn.Read(tran.Buf[:4])
	if err != nil {
		// fmt.Println("conn read datd faild,err=", err)
		// errors.New("read package hander faild")
		return
	}
	//根据buf[:4]转成一个uint32类型
	pkgLen := binary.BigEndian.Uint32(tran.Buf[:4])
	//一句pkgLen读取消息数据
	con, err := tran.Conn.Read(tran.Buf[:pkgLen])
	if con != int(pkgLen) || err != nil {
		fmt.Println("conn Read Package faild err=", err)
		return
	}
	//将pkgLen反序列化成->message.Message
	err = json.Unmarshal(tran.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("PkgLen json.Unmarshal faild err=", err)
		return
	}
	return
}

// WritePackage 读取客户端发过来的数据
func (tran *Transfer) WritePackage(data []byte) (err error) {
	//先发送一个长度到客户端
	pkgLen := uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(tran.Buf[0:4], pkgLen)

	//发送长度
	num, err := tran.Conn.Write(tran.Buf[0:4])
	//验证data长度是否发送成功
	if num != 4 || err != nil {
		fmt.Println("conn write(bytes) err=", err)
		return
	}
	//发送消息本身
	num, err = tran.Conn.Write(data)
	if num != int(pkgLen) || err != nil {
		fmt.Println("conn write(data) err=", err)
		return
	}
	return
}
