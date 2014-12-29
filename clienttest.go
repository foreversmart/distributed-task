
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


func mainclient() {
	fmt.Println("starting client...")
	gonet.ClientInit()
	var msg = gonet.Message{Addr:"127.0.0.1:1256", Content:"haha"}
	gonet.Send(msg, nil)
	// time.Sleep(time.Second * 20)	
	var input string
    for {
	    fmt.Scanln(&input)
	    fmt.Println("done")
	    msg = gonet.Message{Addr:"127.0.0.1:1256", Content: input}
	    gonet.Send(msg, nil)
    }
}



