package main

import (
	"./protocol"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func send(conn net.Conn) {
	for i := 0; i < 100; i++ {
		session :=GetSession()
		words := "{\"ID\":"+ strconv.Itoa(i) +"\",\"Session\":"+session +"2015073109532345\",\"Meta\":\"golang\",\"Content\":\"message\"}"
		conn.Write(protocol.Enpack([]byte(words)))
	}
	fmt.Println("send over")
	defer conn.Close()
}

func GetSession() string {
	gs1:=time.Now().Unix()
	gs2:=strconv.FormatInt(gs1,10)
	return gs2
}

func main() {
	server := "127.0.0.1:1024"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	send(conn)
}