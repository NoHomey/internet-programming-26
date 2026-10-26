package main

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
)

func validateArgs(args []string) (hostPort string, err error) {
	if len(args) != 3 {
		return "", errors.New(fmt.Sprintf("expected 2 arguments got: %d", len(args)-1))
	}

	ipArg, portArg := args[1], args[2]

	ip := net.ParseIP(ipArg)
	if ip == nil {
		return "", fmt.Errorf("Argument 1 is not a valid ip address got: %s", ipArg)
	}
	hostPort = net.JoinHostPort(ipArg, portArg)
	_, err = url.Parse(fmt.Sprintf("http://%s/", hostPort))
	if err != nil {
		return "", fmt.Errorf("Argument 2 is not a valid port number got: %s, error: %w", portArg, err)
	}
	return hostPort, nil
}

func main() {
	hostPort, err := validateArgs(os.Args)
	if err != nil {
		panic(err.Error())
	}

	conn, err := net.Dial("tcp", hostPort)
	if err != nil {
		panic(err.Error())
	}

	n, err := conn.Write([]byte{0b11001100})
	if err != nil {
		panic(err.Error())
	}
	if n != 1 {
		panic("failed to send command")
	}
	result := make([]byte, 1)
	n, err = conn.Read(result)
	if err != nil {
		panic("failed to receive result")
	}
	if n != 1 {
		panic("failed to get result")
	}
	fmt.Println(result[0])
}
