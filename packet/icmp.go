package packet

type Packet interface {
	Parse() []byte
}

type ICMPPacket interface {
	Packet
	CalcChecksum()
}

type ICMP struct {
	Type     byte
	Code     byte
	Checksum uint16
	Data     []byte
}
