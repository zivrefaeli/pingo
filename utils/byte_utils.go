package utils

func ConcatBytes(higher, lower byte) uint16 {
	return (uint16(higher) << 8) | uint16(lower)
}

func GenerateData(bufferSize uint16) []byte {
	data := make([]byte, bufferSize)
	var i uint16 = 0

	for ; i < bufferSize; i++ {
		data[i] = byte(97 + i%23)
	}
	return data
}
