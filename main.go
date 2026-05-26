package main

import(
	"net"
	"fmt"
	"bufio"
	"strings"
)

func main(){

    // set up a tcp socket



	listener, err := net.Listen("tcp", "localhost:8080")

	if err!= nil {
		fmt.Println("An error occured while trying to open the server")
	}

	for{

		var clientMethod string
		var clientPath string

		conn, err := listener.Accept()
		if err!= nil{
			fmt.Println("An error occured while setting up a connection")
		}

		reader := bufio.NewReader(conn)

		// Detrmining the memthod and path of the client
		requestLine , err := reader.ReadString('\n')

		if err != nil{
			fmt.Println("There was an error reading the requestLine")
		}

		parts := strings.Fields(requestLine)

		if len(parts) >= 2{
			clientMethod = parts[0]
			clientPath = parts[1]

			fmt.Printf("The client just made a %s request on path %s", clientMethod, clientPath)
		}

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

		// implementing dynamic routing 
        var body string
		var status string
		
		switch clientPath{
		case "/" : 
			status = "HTTP/1.1 200 Ok\r\n"
			body = "<html><body><h1>Hello! You built an HTTP server from scratch!</h1></body></html>"

        case "/about" : 
			status = "HTTP/1.1 200 Ok\r\n"	
			body = "<html><body><h1>This is the about page</h1></body></html>"

        default : 
			status = "HTTP/1.1 404 NOT FOUND\r\n"
			body = "<html><body><h1>Page could not be found</h1></body></html>"
		}

		// Set up the HTTP response format
		responseText :=  status +
		           		"Content-Type: text/html\r\n" +
						 fmt.Sprintf("Content-Length: %d\r\n", len(body))+
						//  "Connection: close\r\n" +
						 "\r\n" +
						  body
        
		// Here the server rites to the client
        _, err = conn.Write([]byte(responseText))

		if err!= nil{
			fmt.Println(("There is an error "))
		}
		conn.Close()
		fmt.Println("Connection has been closed")
	}

}


