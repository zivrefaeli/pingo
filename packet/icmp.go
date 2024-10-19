package packet

import (
	"net"
	"pingo/utils"
)

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

func SendICMPPacket(conn *net.Conn, icmp ICMPPacket) (ICMP, TTL, error) {
	sentBytes, err := (*conn).Write(icmp.Parse())
	if err != nil {
		return ICMP{}, 0, err
	}

	buffer := make([]byte, IP_HEADER_SIZE+sentBytes)

	_, err = (*conn).Read(buffer)
	if err != nil {
		return ICMP{}, 0, err
	}

	ipHeader := buffer[:IP_HEADER_SIZE]
	icmpHeader := buffer[IP_HEADER_SIZE:]

	return ICMP{
		Type:     icmpHeader[0],
		Code:     icmpHeader[1],
		Checksum: utils.ConcatBytes(icmpHeader[2], icmpHeader[3]),
		Data:     icmpHeader[4:],
	}, TTL(ipHeader[8]), nil
}
