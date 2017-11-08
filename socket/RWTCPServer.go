package main

import (
	"fmt"
	"net"
	"os"
	"sync"
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
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		handleRead(conn)
	}()

	go func() {
		defer wg.Done()
		handleWrite(conn)
	}()
	wg.Wait()
	conn.Close()
}

func handleRead(conn *net.TCPConn) {
	request := make([]byte, 1024)

	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(string(request[:read_len]))
		request = make([]byte, 1024)
	}
}

func handleWrite(conn *net.TCPConn) {
	for {
		time.Sleep(time.Second * 2)
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
