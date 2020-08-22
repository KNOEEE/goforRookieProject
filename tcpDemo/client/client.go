package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("dial err =", err)
		return
	}
	fmt.Println("conn =", conn)
	reader := bufio.NewReader(os.Stdin) //标准输入

	for {
		//从终端读取一行用户输入 发送给server
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString err =", err)
		}
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("Client exit...")
			break
		}
		//将line发送给server
		//客户端发送了n字节的数据", \n必须加双引号
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err =", err)
		}
	}
}