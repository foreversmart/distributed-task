// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

/*
	only server use
*/
package scheduler

import(
	"distributed-task/gonet"
	"distributed-task/gocommand"
	"time"
	"fmt"
)


var ReduceMap string
var reduceChan chan string

type UserReduceFunc func(rm, res string) string

func InitReducer(userReduce UserReduceFunc){
	reduceChan = make(chan string, 100)
	go reduceResult(userReduce)
}

func reduceResult(userReduce UserReduceFunc){
	for{
		select {
		case res := <- reduceChan:
			ReduceMap = userReduce(ReduceMap, res)
			fmt.Println("ReduceMap", ReduceMap)
		case <- time.After(time.Second * 20):
			//30 秒没有更新数据reduce data to client
			if ReduceMap!= ""{
				reduceToClient()
			}
		}
	}
}

/*
	call by server
	to send the res to client
*/
func reduceToClient(){
	tempMap := make(map[string]string)
	tempMap["result"] = ReduceMap
	command := &gocommand.Command{"reduce", "reduce", tempMap}
	temp := gocommand.EnCode(command.GetCommandString())
	gonet.ServerSend(temp)
	ReduceMap = ""
}
