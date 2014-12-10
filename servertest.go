// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

/*
	network and message define (include server and client)
*/
package main

import (
	"fmt"
	// "time"
    "distributed-task/gonet"
)


func main() {
	fmt.Println("server running...")
	go gonet.ServerRead()
	gonet.ServerRun()

}

