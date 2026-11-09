package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"tcp/command"
)

const (
	SendCommandTimeout   = 500 * time.Millisecond
	ReciveCommandTimeout = 2 * SendCommandTimeout
)

type Server struct {
	lock     *sync.Mutex
	data     [16]uint8
	listener net.Listener
}

func NewServer() Server {
	return Server{
		lock:     &sync.Mutex{},
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

func (server *Server) Handle(ctx context.Context, conn net.Conn) {
	handleCtx, cancelHandleCtx := context.WithTimeout(ctx, time.Hour)
	defer cancelHandleCtx()
	for {
		var cmd command.Command
		op := func() error {
			receiveCtx, cancelReceiveCtx := context.WithTimeout(handleCtx, ReciveCommandTimeout)
			cmd = command.ReceiveCommand(receiveCtx, conn)
			err := receiveCtx.Err()
			cancelReceiveCtx()
			return err
		}

		err := command.Retry(
			10,
			command.ConstPause(3*time.Second),
			op,
			func(err error) bool {
				return true
			},
			true,
		)

		if err != nil {
			fmt.Println("Got an error while trying to receive a command", err)
			continue
		}

		if cmd.IsFail() {
			fmt.Println("Failed to Receive command")
			if cmd.IsEOF() {
				fmt.Println("Closing connection")
				err := conn.Close()
				if err != nil {
					fmt.Println("Failed to close a connection", err)
				}
				fmt.Println("Terminating Connection loop")
				return
			}
		}
		var res command.Command
		if cmd.IsGet() {
			key := cmd.Key()
			server.lock.Lock()
			val := server.data[key]
			server.lock.Unlock()
			resp := command.GetOkRespCommand(key, val)
			sendCtx, cancelSendCtx := context.WithTimeout(handleCtx, SendCommandTimeout)
			res = command.SendCommand(sendCtx, resp, conn)
			cancelSendCtx()
		} else {
			key := cmd.Key()
			val := cmd.Value()
			server.lock.Lock()
			server.data[key] = val
			server.lock.Unlock()
			resp := command.SetOkRespCommand(key, val)
			sendCtx, cancelSendCtx := context.WithTimeout(handleCtx, SendCommandTimeout)
			res = command.SendCommand(sendCtx, resp, conn)
			cancelSendCtx()
		}
		if res.IsFail() {
			fmt.Println("Fail to Send Response")
		}
	}
}

func main() {
	port := os.Args[1]
	server := NewServer()
	err := server.Listen(port)
	if err != nil {
		panic(err.Error())
	}
	defer server.Close()

	serverCtx, cancelServerCtx := context.WithTimeout(
		context.Background(),
		6*time.Hour,
	)
	defer cancelServerCtx()

	for {
		fmt.Println("Waiting to accept connection")
		conn, err := server.Accept()
		fmt.Println("Accepted connection")
		if err != nil {
			panic(err.Error())
		}
		go server.Handle(serverCtx, conn)
	}
}
