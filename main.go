package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	service := "127.0.0.1:7777"
	fmt.Printf("service: %v\n", service)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	fmt.Printf("err: %v\r\n", err)
	for {
		conn, err := listener.Accept()
		if err != nil { //等一秒钟
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	ret := "return from server: " + daytime + "\r\n"
	conn.Write([]byte(ret))
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
	}
}
