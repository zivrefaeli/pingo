package packet

type EchoICMP struct {
	ICMP
	Identifier uint16
	Sequence   uint16
}

func (p *EchoICMP) Parse() []byte {
	return p.Data
}

func (p *EchoICMP) CalcChecksum() {
	p.Checksum = 1
}
