package main

import (
	"./protocol"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnectionProtocol(conn)
	}
}

func handleConnectionProtocol(conn net.Conn) {
	tmpBuffer := make([]byte, 0)

	readerChannel := make(chan []byte, 16)
	go reader(readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error", err)
			return
		}
		tmpBuffer = protocol.Depack(append(tmpBuffer, buffer[:n]...))
		Log(string(tmpBuffer))
	}
	defer conn.Close()
}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			Log(string(data))
		}
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	}
}

func Log(v ...interface{})  {
	log.Println(v...)
}

func CheckError(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}