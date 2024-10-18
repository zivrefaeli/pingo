package packet

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
