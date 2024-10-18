package main

import (
	"fmt"
	"net"
	"pingo/packet"
	"pingo/utils"
)

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
}
