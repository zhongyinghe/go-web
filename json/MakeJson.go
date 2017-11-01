package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ID          int    `json:"-"`
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`
	ServerIP    string `json:"ServerIP,omitempty"`
}

func main() {
	s := Server{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    "",
	}

	b, _ := json.Marshal(s)
	fmt.Println(string(b))
}
