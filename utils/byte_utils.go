package utils

const PING_DATA_START, PING_DATA_END = 'a', 'w'
const PING_DATA_DIFF = PING_DATA_END - PING_DATA_START + 1

func ConcatBytes(higher, lower byte) uint16 {
	return (uint16(higher) << 8) | uint16(lower)
}

func GeneratePingData(bufferSize uint16) []byte {
	data := make([]byte, bufferSize)
	var i uint16 = 0

	for ; i < bufferSize; i++ {
		data[i] = PING_DATA_START + byte(i%PING_DATA_DIFF)
	}
	return data
}
