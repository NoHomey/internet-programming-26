package command

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

type Command uint8

func (cmd Command) IsGet() bool {
	return !cmd.IsSet()
}

func (cmd Command) IsSet() bool {
	return uint8(cmd)&uint8(1<<7) > 0
}

func (cmd Command) Key() uint8 {
	return (uint8(cmd) & uint8(15<<3)) >> 3
}

func (cmd Command) Value() uint8 {
	return uint8(cmd) & 7
}

func GetCommand(key uint8) Command {
	var b uint8
	b |= key << 3
	return Command(b)
}

func SetCommand(key uint8, val uint8) Command {
	var b uint8
	b |= 1 << 7
	b |= key << 3
	b |= val
	return Command(b)
}

func (cmd Command) IsOk() bool {
	return cmd.IsSet()
}

func (cmd Command) IsFail() bool {
	return !cmd.IsOk()
}

func (cmd Command) IsEOF() bool {
	return cmd.IsFail() && (cmd&3) > 0
}

func GetOkRespCommand(key uint8, val uint8) Command {
	return SetCommand(key, val)
}

func SetOkRespCommand(key uint8, val uint8) Command {
	return SetCommand(key, val)
}

func FailRespCommand(key uint8, val uint8) Command {
	cmd := SetOkRespCommand(key, val)
	cmd &= 0b01111111
	return cmd
}

func FailToReadResp() Command {
	return FailRespCommand(0, 1)
}

func FailToWriteResp() Command {
	return FailRespCommand(0, 2)
}

func FailConnClosedResp() Command {
	return FailRespCommand(0, 3)
}

func SuccessResp() Command {
	return SetOkRespCommand(0, 0)
}

func ReceiveCommand(ctx context.Context, conn net.Conn) Command {
	b := make([]byte, 1)
	if deadline, ok := ctx.Deadline(); ok {
		conn.SetReadDeadline(deadline)
	}
	n, err := conn.Read(b)
	if err != nil {
		if errors.Is(err, io.EOF) {
			fmt.Println("Connection is closed")
			return FailConnClosedResp()
		}
		fmt.Println("Got an error", err)
		return FailToReadResp()
	}
	if n != 1 {
		return FailToReadResp()
	}
	fmt.Println("Received Command", b[0])
	return Command(b[0])
}

func SendCommand(ctx context.Context, cmd Command, conn net.Conn) Command {
	if deadline, ok := ctx.Deadline(); ok {
		conn.SetWriteDeadline(deadline)
	}
	n, err := conn.Write([]byte{uint8(cmd)})
	if err != nil {
		if errors.Is(err, io.EOF) {
			fmt.Println("Connection is closed")
			return FailConnClosedResp()
		}
		fmt.Println("Got an error", err)
		return FailToWriteResp()
	}
	if n != 1 {
		return FailToWriteResp()
	}
	fmt.Println("Sent response")
	return SuccessResp()
}

func ConstPause(dur time.Duration) func(int, time.Duration) time.Duration {
	return func(int, time.Duration) time.Duration {
		return dur
	}
}

func LinearPause(dur time.Duration) func(int, time.Duration) time.Duration {
	return func(_i int, lastPause time.Duration) time.Duration {
		return lastPause + dur
	}
}

func ExpBackOffPause(linearPart time.Duration, exp float64) func(int, time.Duration) time.Duration {
	return func(_i int, lastPause time.Duration) time.Duration {
		return time.Duration(float64(int64(lastPause))*exp) + linearPart
	}
}

func Retry(
	num int,
	pauseFn func(int, time.Duration) time.Duration,
	op func() error,
	shouldRetry func(error) bool,
	fastRetry bool,
) error {
	lastErr := op()
	if lastErr == nil {
		return nil
	}
	pause := newPause(pauseFn)
	if !fastRetry {
		pause()
	}
	for i := 1; i < num; i++ {
		err := op()
		if err == nil {
			return nil
		}
		lastErr = err
		if !shouldRetry(err) {
			break
		}
		pause()
	}
	return lastErr
}

func newPause(pauseFn func(int, time.Duration) time.Duration) func() {
	var currentDelay time.Duration
	currentPause := 0
	return func() {
		currentDelay = pauseFn(currentPause, currentDelay)
		<-time.After(currentDelay)
		currentPause++
	}
}
