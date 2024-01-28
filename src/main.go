package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	defer listener.Close()
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Could NOT accept connection", err.Error())
			os.Exit(1)
		}
		handleConnection(connection)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Read data
	input := make([]byte, 1024)
	num, err := conn.Read(input)
	if err != nil {
		log.Printf("Couldn't read input")
		return
	}

	log.Println("Received input: ", input[:num])
	conn.Write([]byte("+PONG\r\n"))
	//if bytes.Equal(input, []byte("PING")) {
	//	conn.Write([]byte("+PONG\r\n"))
	//}
}
