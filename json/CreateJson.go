package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	var s ServerSlice
	s.Servers = append(s.Servers, Server{ServerName: "shanghai_vpn", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "beijin_vpn", ServerIP: "192.168.1.1"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json:err:", err)
	}

	fmt.Println(string(b))

	var tmp []Server
	tmp = append(tmp, Server{ServerName: "shanghai_vpn", ServerIP: "127.0.0.1"})
	tmp = append(tmp, Server{ServerName: "beijin_vpn", ServerIP: "192.168.1.1"})
	bt, _ := json.Marshal(tmp)
	fmt.Println(string(bt))
}
