package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type ServerSlice struct {
	Servers []Server `json:"servers"`
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
}
