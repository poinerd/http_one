package main

import(
	"net"
	"fmt"
	"bufio"
)

func main(){
    // set up a tcp socket

	listener, err := net.Listen("tcp", "localhost:8080")

	if err!= nil {
		fmt.Println("An error occured while trying to open the server")
	}

	for{
		conn, err := listener.Accept()
		if err!= nil{
			fmt.Println("TAn error occured while setting up a connection")
		}

		reader := bufio.NewReader(conn)

		for{
			line, err := reader.ReadString('\n')
			
			if err != nil{
				break			
			}
			fmt.Println(line)

			if line == "\r\n" || line == "\n"{
				break
			}

			
		}
		conn.Close()
		fmt.Println("Connection has been closed")
	}

}


