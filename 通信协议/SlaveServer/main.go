package main

import (
	"fmt"
	"net"
	"tcpSocket"
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
		reply:="收到："+string(buffer1)
		lenReply:=len(reply)
		bytesLen:=socket.IntToBytes(lenReply)
		conn.Write(bytesLen)
		conn.Write([]byte(reply))
	}
}

func main(){
	server,err:=net.Listen("tcp","192.168.2.209:8848")
	if err !=nil{
		panic(err)
	}
	defer server.Close()
	for {
		conn,err:=server.Accept()
		if err!=nil{
			panic(err)
		}
		go MsgHandler(conn)
	}
}
