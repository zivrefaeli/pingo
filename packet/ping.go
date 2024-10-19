package packet

import (
	"fmt"
	"net"
	"pingo/utils"
	"time"
)

func SendPingRequest(conn *net.Conn, echoRequest *EchoICMP) (int64, error) {
	startTime := time.Now()
	icmp, ttl, err := SendICMPPacket(conn, echoRequest)
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
