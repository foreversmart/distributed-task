package gonet

import (
	"testing"
	"fmt"
	"time"
    "distributed-task/gonet"
)

func TestClient(t *testing.T){
	fmt.Println("starting client...")
	gonet.ClientInit()
	var msg = gonet.Message{Host:"127.0.0.1", Port:"1256", Content:"haha"}
	for {
		gonet.Send(msg)
		time.Sleep(time.Second * 5)
	}
}