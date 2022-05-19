package system

import (
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
    servAddr := "localhost:6666"
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }

    _, err = conn.Write([]byte(msg))
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

	println("pid ->", os.Getpid())
    println("write to server = ", msg)

    // reply := make([]byte, 1024)

    // _, err = conn.Read(reply)
    // if err != nil {
    //     println("Write to server failed:", err.Error())
    //     os.Exit(1)
    // }

    // println("reply from server=", string(reply))

    conn.Close()

	// grab my IP address and port sent from
	// myIP := GetMyIP()

	// return network info from process, id of process
	return NetworkData{
		// DestinationIP: destinationAddr,
		// DestinationPort: destinationPort,
		// SourceIP: fmt.Sprintf("%s", myIP.IP.To4()),	// errors in linter but requires instead of string()
		// SourcePort: fmt.Sprintf("%d", myIP.Port),
		// BytesSent: fmt.Sprintf("%d", byteCount),
		// Protocol: "UDP",
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