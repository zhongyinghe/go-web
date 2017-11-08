package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	service := "127.0.0.1:9999"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	defer listener.Close()
	checkError(err)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("request err: ", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn *net.TCPConn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(request[:read_len]))

		if read_len == 0 {
			break
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
