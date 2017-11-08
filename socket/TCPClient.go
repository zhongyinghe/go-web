package main

import (
	"fmt"
	"net"
	"time"
)

var closeFlag chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	go RecivedMessage(conn)
	conn.Write([]byte("first Request!"))
	<-closeFlag
}

func RecivedMessage(conn *net.TCPConn) {
	respond := make([]byte, 1024)
	for {
		read_len, err := conn.Read(respond)
		if err != nil {
			fmt.Println(err)
			closeFlag <- true
			break
		}

		fmt.Println(string(respond[:read_len]))

		time.Sleep(time.Second)

		conn.Write([]byte(time.Now().String() + ": send client Request"))
		respond = make([]byte, 1024)
	}
}
