package tcpSocket

import (
	"bytes"
	"encoding/binary"
)

type TcpSocket struct {

}

func (item *TcpSocket)BytesToInt(bts []byte)int{
	byteBuffer:=bytes.NewBuffer(bts)
	var data int64
	binary.Read(byteBuffer,binary.BigEndian,&data)
	return int(data)
}
func (item *TcpSocket)IntToBytes(n int)[]byte{
	data:=int64(n)
	byteBuffer:=bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer,binary.BigEndian,data)
	return byteBuffer.Bytes()
}
