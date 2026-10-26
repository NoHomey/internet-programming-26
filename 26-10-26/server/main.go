package main

import (
	"fmt"
	"net"
	"os"
)

type Server struct {
	data     [16]uint8
	listener net.Listener
}

func NewServer() Server {
	return Server{
		listener: nil,
	}
}

func (server *Server) Listen(port string) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return fmt.Errorf("failed to tcp listen on port %s: %w", port, err)
	}
	server.listener = listener
	fmt.Println("Server listening at port", port)
	return nil
}

func (server *Server) Accept() (net.Conn, error) {
	return server.listener.Accept()
}

func (server *Server) Close() {
	err := server.listener.Close()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully closed the server")
}

func main() {
	port := os.Args[1]
	server := NewServer()
	err := server.Listen(port)
	if err != nil {
		panic(err.Error())
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err.Error())
		}
		cmd := make([]byte, 1)
		n, err := conn.Read(cmd)
		if err != nil {
			panic(err.Error())
		}
		if n != 1 {
			panic("failed to receive a command")
		}
		fmt.Println("Received Command", cmd[0])
		n, err = conn.Write(cmd)
		if err != nil {
			panic(err.Error())
		}
		if n != 1 {
			panic("failed to send a command")
		}
		fmt.Println("Sent response")
		os.Exit(1)
	}
}
