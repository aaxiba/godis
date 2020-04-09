package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkClientError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkClientError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkClientError(err)

	_ = conn.SetReadDeadline(time.Now().Add(time.Second))
	for {

		response := make([]byte, 512)
		_, err = conn.Read(response)
		if err != nil {
			_ = fmt.Errorf("read err: %v", err)
			os.Exit(0)
		}
		fmt.Println(string(response))
	}

}
func checkClientError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s \r\n", err.Error())
	}
}
