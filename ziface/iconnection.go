package ziface

import "net"

// 定义连接接⼝口
type IConnection interface {
	//启动连接，让当前连接开始工作
	Start()
	//停止连接，结束当前连接状态M
	Stop()

	//从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	//获取当前连接ID
	GetConnID() uint32
	//获取远程客户端地址信息
	RemoteAddr() net.Addr

	//直接将Message数据发送数据给远程的TCP客户端(无缓冲)
	SendMsg(msgId uint32, data []byte) error
	//直接将Message数据发送给远程的TCP客户端(有缓冲)
	SendBuffMsg(msgId uint32, data []byte) error

	//设置链接属性
	SetProperty(key string, value interface{})
	//获取链接属性
	GetProperty(key string) (interface{}, error)
	//移除链接属性
	RemoveProperty(key string)
}

// 所有 conn 链接在处理理业务的函数接口，
// 参数1 socket原⽣链接
// 参数2 客户端请求的数据
// 参数3 客户端请求的数据⻓长度
// 想要指定⼀一个conn的处理理业务，只要定义 一个HandFunc类型的函数，然后和该链接绑定就可以了。
type HandFunc func(*net.TCPConn, []byte, int) error
