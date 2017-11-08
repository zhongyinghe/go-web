package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := "127.0.0.1:9998"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	read_len, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		return
	}

	//读到的数据
	fmt.Println(string(data[:read_len]))

	//写数据
	daytime := time.Now().String() + ":response data"
	conn.WriteToUDP([]byte(daytime), remoteAddr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
