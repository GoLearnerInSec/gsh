package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handle(conn net.Conn) {
	for {
		fmt.Println("输入传输数据:  ")
		reader := bufio.NewReader(os.Stdin)
		res, _, err := reader.ReadLine()

		if nil != err {
			fmt.Println("reader.ReadLine() error:", err)
		}

		cnt, err := conn.Write(res)
		//创建容器用于接收读取到的数据

		fmt.Println("客户端回显")
		for {
			buf := make([]byte, 1024) //使用make创建切片
			//cnt是实际读到的字符数
			cnt, err = conn.Read(buf)
			if err != nil {
				fmt.Println("conn.Read err:", err)
				break
			}
			fmt.Println("length:", cnt, "data:\n", string(buf[:cnt]))
			if cnt < 1024 {
				break
			}
		}

		//服务器业务 小写转成大写
		//upperData := strings.ToUpper(string(buf[:cnt])) //只处理前cnt个字符
		//cnt, err = conn.Write([]byte(test_send))

		//fmt.Println("Server to Client...")
		//fmt.Println("length:", cnt, "data", res)
	}
}

func main() {

	//创建监听
	ip := "127.0.0.1"
	port := 9999
	address := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	//主GO程负责监听处理连接
	//子GO程负责处理数据 实现客户端与服务端互动的具体逻辑

	//主GO程提取多个连接
	for {
		fmt.Println("listening...")
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}

		fmt.Println("Accept success!")

		//启动一个子GO程 执行业务处理函数
		go handle(conn)
	}
}
