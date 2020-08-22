package main
import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//循环地接收客户端发送的数据
	defer conn.Close()
	for {
		//创建一个新的slice
		buf := make([]byte, 1024)
		//等待客户端通过conn发送数据
		//若客户端没有Write 则阻塞在这里
		// fmt.Printf("Server is waiting for %s to send msg...\n",
		// 	conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("a connection was closed by the remote host")
			return
		}
		//显示客户发送的内容
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("Server starts listening")
	//使用tcp网络协议 监听本地8888端口
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("listen err =", err)
		return
	}
	defer listen.Close() //延时关闭listen
	//循环等待客户端的连接
	for {
		fmt.Println("waiting cli conn...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err =", err)
		} else {
			fmt.Println("Accept() conn =", conn)
			fmt.Println("Cli ip =", 
				conn.RemoteAddr().String())
		}
		go process(conn)
	}
}