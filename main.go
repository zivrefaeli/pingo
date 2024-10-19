package main

import (
	"fmt"
	"net"
	"pingo/packet"
	"pingo/utils"
	"time"
)

func SendPingRequest(conn *net.Conn, echoRequest *packet.EchoICMP) (int64, error) {
	startTime := time.Now()
	icmp, err := packet.SendICMPPacket(conn, echoRequest)
	timeDiff := time.Since(startTime)
	if err != nil {
		return -1, err
	}

	identifier := utils.ConcatBytes(icmp.Data[0], icmp.Data[1])
	sequence := utils.ConcatBytes(icmp.Data[2], icmp.Data[3])
	icmp.Data = icmp.Data[4:]

	echoReply := packet.EchoICMP{
		ICMP:       icmp,
		Identifier: identifier,
		Sequence:   sequence,
	}
	if echoReply.Checksum != echoReply.CalcChecksum() {
		return -1, fmt.Errorf("invalid checksum response 0x%x", echoReply.Checksum)
	}

	ms := timeDiff.Milliseconds()
	if ms == 0 {
		fmt.Printf("Reply from %s: bytes=%d time<1ms TTL=X\n", (*conn).RemoteAddr().String(), len(echoReply.Data))
	} else {
		fmt.Printf("Reply from %s: bytes=%d time=%dms TTL=X\n", (*conn).RemoteAddr().String(), len(echoReply.Data), ms)
	}
	return ms, nil
}

func StartPinging(targetName string, echoRequestsCount int, bufferSize uint16) error {
	echoRequest := packet.EchoICMP{
		ICMP: packet.ICMP{
			Type: 8,
			Code: 0,
			Data: utils.GeneratePingData(bufferSize),
		},
		Identifier: 1,
		Sequence:   10,
	}

	conn, err := net.Dial("ip:icmp", targetName)
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Printf("\nPinging %s with %d bytes of data:\n", targetName, bufferSize)

	for i := 0; i < echoRequestsCount; i++ {
		_, err = SendPingRequest(&conn, &echoRequest)
		if err != nil {
			fmt.Println(err)
		}

		if i < echoRequestsCount-1 {
			time.Sleep(time.Second)
			echoRequest.Sequence++
		}
	}
	return nil
}

func main() {
	StartPinging("google.com", 4, 32)
	// lostPackets := 0
	// var msMin, msMax, msSum int64 = math.MaxInt64, math.MinInt64, 0

	// for i := 0; i < totalPackets; i++ {
	// 	ms, err := SendPingRequest(&conn, &request)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		lostPackets++
	// 	}

	// 	if ms > msMax {
	// 		msMax = ms
	// 	} else if ms < msMin {
	// 		msMin = ms
	// 	}
	// 	msSum += ms

	// 	if i < totalPackets-1 {
	// 		time.Sleep(time.Second)
	// 		request.Sequence++
	// 	}
	// }

	// lossPercentage := int(float64(lostPackets) / float64(totalPackets) * 100)
	// fmt.Printf("\nPing statistics for %s:\n", conn.RemoteAddr().String())
	// fmt.Printf("    Packets: Sent = %d, Received = %d, Lost = %d (%d%% loss),\n", totalPackets, totalPackets-lostPackets, lostPackets, lossPercentage)
	// fmt.Println("Approximate round trip times in milli-seconds:")
	// fmt.Printf("    Minimum = %dms, Maximum = %dms, Average = %dms\n", msMin, msMax, msSum/int64(totalPackets))
}
