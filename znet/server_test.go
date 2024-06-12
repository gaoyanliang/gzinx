package znet

import (
	"fmt"
	"gzinx/ziface"
	"net"
	"testing"
	"time"
)

/*
自定义 router
*/
//ping test 自定义路由
type PingRouter struct {
	BaseRouter //一定要先基础BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

/*
模拟客户端
*/
func ClientTest() {

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		_, err := conn.Write([]byte("Zinx V0.3"))
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error ")
			return
		}

		// buf 是一个固定大小的字节数组，当读取的字节数少于缓冲区大小时，未使用的部分会包含默认值（0 值），这些值可能会显示为乱码。
		// 要解决这个问题，可以使用 buf[:cnt] 只打印实际读取的字节数
		fmt.Printf("[Client] server call back : %s, cnt = %d\n", string(buf[:cnt]), cnt)
		time.Sleep(1 * time.Second)
	}
}

// Server 模块的测试函数
func TestServer(t *testing.T) {

	/*
		服务端测试
	*/
	// 1. 创建一个server 句柄 s
	s := NewServer()
	// 1.2. 添加测试路由
	s.AddRouter(&PingRouter{})

	/*
		客户端测试
	*/
	go ClientTest()

	// 2. 开启服务
	s.Serve()
}
