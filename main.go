package main

import (
	"pingo/packet"
)

func main() {
	packet.StartPinging("google.com", 4, 32)
}
