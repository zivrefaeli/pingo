package main

import (
	"fmt"
	"net"
	"pingo/packet"
	"pingo/utils"
)

func SendPacket(conn *net.Conn, request *packet.EchoICMP) {
	sentBytes, err := (*conn).Write(request.Parse())
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 100)

	readBytes, err := (*conn).Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sentBytes, readBytes)
	fmt.Println(buffer)
}

func main() {
	address := "127.0.0.1"
	var size uint16 = 32

	request := packet.EchoICMP{
		ICMP: packet.ICMP{
			Type: 8,
			Code: 0,
			Data: utils.GeneratePingData(size),
		},
		Identifier: 1,
		Sequence:   10,
	}

	conn, err := net.Dial("ip:icmp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Printf("\nPinging %s with %d bytes of data:\n", address, size)

	SendPacket(&conn, &request)
}
