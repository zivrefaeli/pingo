package main

import (
	"fmt"
	"pingo/packet"
	"pingo/utils"
)

func main() {
	var size uint16 = 100

	pRequest := packet.EchoICMP{
		ICMP: packet.ICMP{
			Type: 8,
			Code: 0,
			Data: utils.GenerateData(size),
		},
		Identifier: 1,
		Sequence:   54,
	}

	pRequest.CalcChecksum()

	fmt.Printf("0x%x\n", pRequest.Checksum)
}
