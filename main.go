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
			fmt.Println("An error occured while setting up a connection")
		}

		reader := bufio.NewReader(conn)

		for{
			line, err := reader.ReadString('\n')
			
			if err != nil{
				break			
			}
			fmt.Println(line)

			if line == "\r\n" || line == "\n" || line == "wrap\n"{
				break
			}
			
		}
		body := "<html><body><h1>Hello! You built an HTTP server from scratch!</h1></body></html>"
	
		responseText := "HTTP/1.1 200 Ok\r\n" +
		           		"Content-Type: text/html\r\n" +
						 fmt.Sprintf("Content-Length: %d\r\n", len(body))+
						 "Connection: close\r\n" +
						 "\r\n" +
						  body
   
        _, err = conn.Write([]byte(responseText))

		if err!= nil{
			fmt.Println(("There is an error "))
		}
		conn.Close()
		fmt.Println("Connection has been closed")
	}

}


