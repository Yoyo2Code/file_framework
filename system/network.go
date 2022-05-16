package system

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type NetworkData struct {
	DestinationIP string
	DestinationPort string
	SourceIP string
	SourcePort string
	BytesSent string
	Protocol string
}

// ex. localhost:8080, "hello world"
func SendMessage(address string, msg string) (data NetworkData, pid int) {
	// establish network connection
 	conn, err := net.Dial("udp", address)
	if err != nil {
			panic(err)
	}

	// get separated destination port and address
	destinationAddr, destinationPort := GetIpAndAddr(address)

	defer conn.Close()	// closes connection

	// send message
	byteCount, _ := fmt.Fprint(conn, msg)

	// grab my IP address and port sent from
	myIP := GetMyIP()

	// return network info from process, id of process
	return NetworkData{
		DestinationIP: destinationAddr,
		DestinationPort: destinationPort,
		SourceIP: fmt.Sprintf("%s", myIP.IP.To4()),	// errors in linter but requires instead of string()
		SourcePort: fmt.Sprintf("%d", myIP.Port),
		BytesSent: fmt.Sprintf("%d", byteCount),
		Protocol: "UDP",
	},os.Getpid()
}

// separate localhost:8080 -> localhost 8080
func GetIpAndAddr(s string) (addr string, port string) {
	seperatedString := strings.Split(s, ":")
	return seperatedString[0], seperatedString[1]
}

// get my IP address and port
func GetMyIP() net.UDPAddr {
	conn, err := net.Dial("udp", "8.8.8.8:80") // reach out to google dns

	if err != nil {
		panic("error finding My IP address")
	}

     defer conn.Close()
     localAddr := conn.LocalAddr().(*net.UDPAddr)
	 return *localAddr
}