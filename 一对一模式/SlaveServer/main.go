package main

import (
	"fmt"
	"net"
	"tcpSocket"
	"spyder"
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
		// 根据收到的url进行爬虫收集邮箱地址，将邮箱地址返回给masterServer
		url:=string(buffer1)
		var htmlHandler spyder.Html
		mailList:=htmlHandler.GetMail(url)
		for _,email:=range mailList{
			lenReply:=len(email)
			bytesLen:=socket.IntToBytes(lenReply)
			conn.Write(bytesLen)
			conn.Write([]byte(email))
		}
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
