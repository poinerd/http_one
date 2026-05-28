package main

import (
	"fmt"
	"net"
)



func main() {

	fmt.Println("Server is up, listening")

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("An error occured while trying to open the server")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
            fmt.Println("An error occurred while setting up a connection")
            continue 
        }
		go handleClientConnection(conn)
		fmt.Println("Connection has been closed")
	}

}



