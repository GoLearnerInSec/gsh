package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	//与服务器建立连接
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	fmt.Println("Dial success")
	//client0发送的内容是 "hello this is client 0"
	//client1发送的内容是 "hello this is client 1"
	//sendData := []byte("hello this is client 0")

	//向服务器多次发送数据
	for {

		fmt.Println("输入传输数据:  ")
		reader := bufio.NewReader(os.Stdin)

		res, _, err := reader.ReadLine()

		if nil != err {
			fmt.Println("reader.ReadLine() error:", err)
		}

		cnt, err := conn.Write(res)

		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}

		fmt.Println("send data to server...")

		//接收服务器返回的数据
		buf := make([]byte, 1024)

		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("Server to Client...")
		fmt.Println("data:   ", string(buf[:cnt])) //只处理前cnt个字符

		//休息5秒再发送
		//time.Sleep(5 * time.Second)
	}
}
