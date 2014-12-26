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
    "distributed-task/scheduler"
)


func main() {

	//collect in client
	var collect = func(msg string) {
		fmt.Println("client collect message:", msg)
	}

	scheduler.Runner(
		//define data and task and collect
		func(){
			var data  = map[string]string{"1": "111", "2": "122111"}
			scheduler.AllocateData("", scheduler.TypeSequence, data, collect)
		}, 
		//define what the task
		func(key, value string) string{
			fmt.Println("key:", key)
			fmt.Println("value:", value)
			return "aaaa"
		},
		//server reduce
		func(rm, res string) string{
			return rm + res
		},
	)
}
