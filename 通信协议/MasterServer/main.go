package main

import (
	"fmt"
	"net"
	"strconv"
	"tcpSocket"
	"time"
)

func MsgHandler(conn net.Conn){
	defer conn.Close()
	for {
		buffer:=make([]byte,8)
		n,err:=conn.Read(buffer)
		if err !=nil || n!=8{
			fmt.Println("连接关闭",conn)
			return
		}
		var socket tcpSocket.TcpSocket
		length:=socket.BytesToInt(buffer)
		buffer1:=make([]byte,length)
		n,err=conn.Read(buffer1)
		if err!=nil || n!=length{
			fmt.Println("连接关闭",conn)
			return
		}
		fmt.Println("收到：",string(buffer1))
	}
}

func main(){
	addr,err:=net.ResolveTCPAddr("tcp4","192.168.2.209:8848")
	if err !=nil{
		panic(err)
	}
	conn,err:=net.DialTCP("tcp",nil,addr)
	if err !=nil{
		panic(err)
	}
	go MsgHandler(conn) //处理接收消息
	var socket tcpSocket.TcpSocket
	for i:=0;i<10;i++{
		str:="message "+strconv.Itoa(i)
		length:=len(str)
		btsLen:=socket.IntToBytes(length)
		conn.Write(btsLen)
		conn.Write([]byte(str))
	}
	time.Sleep(10*time.Second)
}
