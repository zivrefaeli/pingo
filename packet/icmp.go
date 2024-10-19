package packet

import (
	"net"
	"pingo/utils"
)

const IP_HEADER_SIZE = 20

type Packet interface {
	Parse() []byte
}

type ICMPPacket interface {
	Packet
	CalcChecksum() uint16
}

type ICMP struct {
	Type     byte
	Code     byte
	Checksum uint16
	Data     []byte
}

func SendICMPPacket(conn *net.Conn, icmp ICMPPacket) (ICMP, error) {
	sentBytes, err := (*conn).Write(icmp.Parse())
	if err != nil {
		return ICMP{}, err
	}

	buffer := make([]byte, IP_HEADER_SIZE+sentBytes)

	_, err = (*conn).Read(buffer)
	if err != nil {
		return ICMP{}, err
	}

	icmpHeader := buffer[IP_HEADER_SIZE:]

	return ICMP{
		Type:     icmpHeader[0],
		Code:     icmpHeader[1],
		Checksum: utils.ConcatBytes(icmpHeader[2], icmpHeader[3]),
		Data:     icmpHeader[4:],
	}, nil
}
