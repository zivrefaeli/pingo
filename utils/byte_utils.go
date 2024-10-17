package utils

func ConcatBytes(higher, lower byte) uint16 {
	return (uint16(higher) << 8) | uint16(lower)
}
