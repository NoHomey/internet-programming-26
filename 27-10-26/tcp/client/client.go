package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"os"
	"time"

	"tcp/command"
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

	key := uint8(5)
	val := uint8(rand.Int() % 8)

	cmd := command.SetCommand(key, val)
	time.Sleep(5 * time.Second)
	res := command.SendCommand(cmd, conn)
	if res.IsFail() {
		fmt.Println("Failed to Send a command")
		os.Exit(1)
	}
	time.Sleep(time.Second)
	resp := command.ReceiveCommand(conn)
	if resp.IsOk() {
		key := resp.Key()
		val := resp.Value()
		fmt.Println("Success", "Set", key, val)
	} else {
		fmt.Println("Fail", "Set", key, 3)
		os.Exit(1)
	}

	cmd = command.GetCommand(key)
	time.Sleep(time.Second)
	resp = command.SendCommand(cmd, conn)
	if res.IsFail() {
		fmt.Println("Failed to Send a command")
		os.Exit(1)
	}
	time.Sleep(time.Second)
	resp = command.ReceiveCommand(conn)
	if resp.IsOk() {
		key := resp.Key()
		val := resp.Value()
		fmt.Println("Success", "Get", key, val)
	} else {
		fmt.Println("Fail", "Get", key)
	}
}
