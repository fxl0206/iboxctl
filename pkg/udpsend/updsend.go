package udpsend

import (
	"fmt"
	"net"
)

func Exec(command string) {
	//connect server
	conn, err := net.DialUDP("udp", &net.UDPAddr{
		Port: 7676,
	}, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 1, 253),
		Port: 5053,
	})

	if err != nil {
		fmt.Printf("connect failed, err: %v\n", err)
		return
	}

	//send data
	_, err = conn.Write([]byte(command))
	if err != nil {
		fmt.Printf("send data failed, err : %v\n", err)
		return
	}

	//receive data from server
	result := make([]byte, 16)
	_, remoteAddr, err := conn.ReadFromUDP(result)
	if err != nil {
		fmt.Printf("receive data failed, err: %s\n", err)
		return
	}
	fmt.Printf("receive from addr: %v  data: %b\n", remoteAddr, result)
	/*fmt.Printf("receive from addr: %v  data: %b\n", remoteAddr, result[8])
	fmt.Printf("receive from addr: %v  data: %b\n", remoteAddr, result[9])
	fmt.Printf("receive from addr: %v  data: %b\n", remoteAddr, result[10])*/

	/*fmt.Printf("sts: %v  data: %b\n", remoteAddr, result[11])
	fmt.Printf("sts: %v  data: %b\n", remoteAddr, result[12])
	fmt.Printf("sts: %v  data: %b\n", remoteAddr, result[13])
	fmt.Printf("receive from addr: %v  data: %x\n", remoteAddr, result[14])
	fmt.Printf("receive from addr: %v  data: %x\n", remoteAddr, result[15])*/


}
