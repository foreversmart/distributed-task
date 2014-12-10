
package gonet

import (
	"testing"
	"fmt"
	"distributed-task/gonet"
)

func TestServer(t *testing.T){
	fmt.Println("starting server...")
	gonet.ServerRun()
}