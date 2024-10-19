package packet

import (
	"fmt"
	"math"
	"math/rand/v2"
	"net"
	"pingo/utils"
	"time"
)

func sendPingRequest(conn *net.Conn, echoRequest *EchoICMP) (int64, error) {
	startTime := time.Now()
	icmp, ttl, err := sendICMPPacket(conn, echoRequest)
	timeDiff := time.Since(startTime)
	if err != nil {
		return -1, err
	}

	identifier := utils.ConcatBytes(icmp.Data[0], icmp.Data[1])
	sequence := utils.ConcatBytes(icmp.Data[2], icmp.Data[3])
	icmp.Data = icmp.Data[4:]

	echoReply := EchoICMP{
		ICMP:       icmp,
		Identifier: identifier,
		Sequence:   sequence,
	}
	if echoReply.Checksum != echoReply.CalcChecksum() {
		return -1, fmt.Errorf("invalid checksum response 0x%x", echoReply.Checksum)
	}

	ms := timeDiff.Milliseconds()
	if ms == 0 {
		fmt.Printf("Reply from %s: bytes=%d time<1ms TTL=%d\n", (*conn).RemoteAddr().String(), len(echoReply.Data), ttl)
	} else {
		fmt.Printf("Reply from %s: bytes=%d time=%dms TTL=%d\n", (*conn).RemoteAddr().String(), len(echoReply.Data), ms, ttl)
	}
	return ms, nil
}

func StartPinging(targetName string, echoRequestsCount int, bufferSize uint16) error {
	echoRequest := EchoICMP{
		ICMP: ICMP{
			Type: ECHO_REQUEST_TYPE,
			Code: 0,
			Data: utils.GeneratePingData(bufferSize),
		},
		Identifier: uint16(rand.IntN(10)),
		Sequence:   uint16(rand.IntN(1000)),
	}

	conn, err := net.Dial("ip:icmp", targetName)
	if err != nil {
		return err
	}
	defer conn.Close()

	targetIp := conn.RemoteAddr().String()
	if targetName != targetIp {
		fmt.Printf("\nPinging %s [%s] with %d bytes of data:\n", targetName, targetIp, bufferSize)
	} else {
		fmt.Printf("\nPinging %s with %d bytes of data:\n", targetName, bufferSize)
	}

	lostPackets := 0
	var minMs, maxMs, sumMs int64 = math.MaxInt64, math.MinInt64, 0

	for i := 0; i < echoRequestsCount; i++ {
		ms, err := sendPingRequest(&conn, &echoRequest)
		if err != nil {
			fmt.Println(err)
			lostPackets++
		}

		if ms > maxMs {
			maxMs = ms
		} else if ms < minMs {
			minMs = ms
		}
		sumMs += ms

		if i < echoRequestsCount-1 {
			time.Sleep(time.Second)
			echoRequest.Sequence++
		}
	}

	lossPercentage := int(float64(lostPackets) / float64(echoRequestsCount) * 100)
	msAvg := sumMs / int64(echoRequestsCount)

	fmt.Printf("\nPing statistics for %s:\n", targetIp)
	fmt.Printf("    Packets: Sent = %d, Received = %d, Lost = %d (%d%% loss),\n", echoRequestsCount, echoRequestsCount-lostPackets, lostPackets, lossPercentage)
	fmt.Println("Approximate round trip times in milli-seconds:")
	fmt.Printf("    Minimum = %dms, Maximum = %dms, Average = %dms\n", minMs, maxMs, msAvg)

	return nil
}
