package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("Error starting server: %s", err)
    }
    defer listener.Close()
    fmt.Println("Chat server started on :8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Error accepting connection: %s", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        // Print received message from client
        fmt.Printf("Client: %s\n", scanner.Text())
    }

    if scanner.Err() != nil {
        log.Printf("Error reading from client: %s", scanner.Err())
        return
    }
}
