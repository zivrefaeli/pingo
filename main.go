package main

import (
	"fmt"
	"net"
	"pingo/packet"
	"pingo/utils"
)

const IP_HEADER_SIZE = 20

func SendPacket(conn *net.Conn, request *packet.EchoICMP) {
	sentBytes, err := (*conn).Write(request.Parse())
	if err != nil {
		fmt.Println(err)
		return
	}
	buffer := make([]byte, IP_HEADER_SIZE+sentBytes)

	_, err = (*conn).Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	ipHeader := buffer[:IP_HEADER_SIZE]
	icmpHeader := buffer[IP_HEADER_SIZE:]

	response := packet.EchoICMP{
		ICMP: packet.ICMP{
			Type:     icmpHeader[0],
			Code:     icmpHeader[1],
			Checksum: utils.ConcatBytes(icmpHeader[2], icmpHeader[3]),
			Data:     icmpHeader[8:],
		},
		Identifier: utils.ConcatBytes(icmpHeader[4], icmpHeader[5]),
		Sequence:   utils.ConcatBytes(icmpHeader[6], icmpHeader[7]),
	}

	if response.Checksum == response.CalcChecksum() {
		fmt.Printf("Reply from %s: bytes=%d time=? TTL=%d\n", (*conn).RemoteAddr().String(), len(response.Data), ipHeader[8])
	} else {
		fmt.Printf("Invalid checksum response 0x%x\n", response.Checksum)
	}
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
