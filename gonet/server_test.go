package gonet

import (
	"fmt"
	"github.com/foreversmat/distributed-task/gonet"
	"testing"
)

func TestServer(t *testing.T) {
	fmt.Println("starting server...")
	gonet.ServerRun()
}
