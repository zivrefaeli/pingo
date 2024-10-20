package packet

import (
	"math"
	"pingo/utils"
)

const ECHO_REQUEST_TYPE = 8

type EchoICMP struct {
	ICMP
	Identifier uint16
	Sequence   uint16
}

func (p *EchoICMP) Parse() []byte {
	p.Checksum = p.CalcChecksum()

	icmpHeader := make([]byte, 8)
	icmpHeader[0] = p.Type
	icmpHeader[1] = p.Code
	icmpHeader[2] = byte(p.Checksum >> 8)
	icmpHeader[3] = byte(p.Checksum)
	icmpHeader[4] = byte(p.Identifier >> 8)
	icmpHeader[5] = byte(p.Identifier)
	icmpHeader[6] = byte(p.Sequence >> 8)
	icmpHeader[7] = byte(p.Sequence)

	icmpHeader = append(icmpHeader, p.Data...)

	return icmpHeader
}

func (p *EchoICMP) CalcChecksum() uint16 {
	dataSize := len(p.Data)
	var dataHigher, dataLower byte
	var sumCarry uint16
	headersSum := utils.ConcatBytes(p.Type, p.Code) + p.Identifier + p.Sequence

	for i := 0; i < dataSize; i += 2 {
		dataHigher = p.Data[i]
		if i+1 < dataSize {
			dataLower = p.Data[i+1]
		} else {
			dataLower = 0
		}
		diff := utils.ConcatBytes(dataHigher, dataLower)

		if int(headersSum)+int(diff) > math.MaxUint16 {
			sumCarry++
		}
		headersSum += diff
	}

	return ^(headersSum + sumCarry)
}
