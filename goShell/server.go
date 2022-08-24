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
