package main

import (
	"fmt"
	"pingo/packet"
	"pingo/utils"
)

func main() {
	request := packet.EchoICMP{
		ICMP: packet.ICMP{
			Type: 8,
			Code: 0,
			Data: utils.GeneratePingData(30),
		},
		Identifier: 1,
		Sequence:   10,
	}

	fmt.Println(request)
	fmt.Println(request.Parse())
	fmt.Printf("Checksum = 0x%x\n", request.Checksum)
}
