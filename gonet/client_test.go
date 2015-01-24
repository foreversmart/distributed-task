package gonet

import (
	"fmt"
	"github.com/foreversmart/distributed-task/gonet"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	fmt.Println("starting client...")
	gonet.ClientInit()
	var msg = gonet.Message{Addr: "127.0.0.1:1256", Content: "haha"}
	fmt.Println("runging")
	for {
		fmt.Println("runging")
		gonet.Send(msg)
		time.Sleep(time.Second * 5)
	}
}
