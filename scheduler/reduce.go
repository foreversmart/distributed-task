// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

/*
	only server use
*/
package scheduler

import(
	"distributed-task/gonet"
	"distributed-task/goncommand"
)


var reduceMap string
var reduceChan chan string
var clientReduce

func InitReducer(){
	reduceChan = make(chan string, 100)
	go reduceResult()
}

func reduceResult(userReduce func(reduceMap, res string)){
	for{
		select {
		case res := <- reduceChan:
			userReduc(userReduce, res)
		case <- time.After(time.Second * 30):
			//30 秒没有更新数据reduce data to client
			reduceToClient()
		}
	}
}

/*
	call by server
	to send the res to client
*/
func reduceToClient(){
	tempMap := make(map[string]string)
	tempMap["result"] = reduceMap
	command := &goncommand.Command{"reduce", "reduce", tempMap}
	gocommand.EnCode(command.GetCommandString())
	gonet.ServerSend()
	reduceMap = ""
}
