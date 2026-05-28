package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

)
type serverResponse struct{
	    status string
		contentType string
		body []byte
}


func (r *serverResponse) createServerResponse() string {
    return fmt.Sprintf(
        "%s\r\n"+
        "Content-Type: %s\r\n"+
        "Content-Length: %d\r\n"+
        "\r\n"+
        "%s",
        r.status, r.contentType, len(r.body), string(r.body),
    )
}


func readFiles(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getFilePath(clientPath string) string {
    if clientPath == "/" {
        return "public/index.html"
    }

    fullPath := "public" + clientPath

    if _, err := os.Stat(fullPath); err == nil {
        return fullPath
    }

    htmlPath := fullPath + ".html"
    if _, err := os.Stat(htmlPath); err == nil {
        return htmlPath
    }

    return ""
}



func handleClientConnection(conn net.Conn) {
    defer conn.Close()

    reader := bufio.NewReader(conn)

    requestLine, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("There was an error reading the requestLine")
        return
    }

    var clientMethod, clientPath string
    parts := strings.Fields(requestLine)
    if len(parts) >= 2 {
        clientMethod = parts[0]
        clientPath = parts[1]
        fmt.Printf("The client just made a %s request on path %s\n", clientMethod, clientPath)
    }

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        fmt.Print(line)
        if line == "\r\n" || line == "\n" || line == "wrap\n" {
            break
        }
    }

    var response serverResponse
    fullPath := getFilePath(clientPath)
    bytes, err := readFiles(fullPath)

    if err != nil {
        response = serverResponse{
            status:      "HTTP/1.0 404 NOT FOUND",
            contentType: "text/html",
            body:        []byte("<html><body><h1>404: Page not found</h1></body></html>"),
        }
    } else {
        response = serverResponse{
            status:      "HTTP/1.0 200 OK",
            contentType: "text/html", // (Swap to getContentType(fullPath) later)
            body:        bytes,
        }
    }

    _, err = conn.Write([]byte(response.createServerResponse()))
    if err != nil {
        fmt.Println("There is an error writing response")
    }
    
    fmt.Println("Connection has been closed")
}
