package packet

import (
	"math"
	"pingo/utils"
)

type EchoICMP struct {
	ICMP
	Identifier uint16
	Sequence   uint16
}

func (p *EchoICMP) Parse() []byte {
	return p.Data
}

func (p *EchoICMP) CalcChecksum() {
	var headersSum uint16 = utils.ConcatBytes(p.Type, p.Code) + p.Identifier + p.Sequence
	var sumCarry uint16

	dataSize := len(p.Data)
	var dataHigher, dataLower byte

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

	p.Checksum = ^(headersSum + sumCarry)
}
