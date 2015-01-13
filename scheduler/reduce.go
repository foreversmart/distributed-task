// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

/*
	only server use
*/
package scheduler

import (
	"distributed-task/gocommand"
	"distributed-task/gonet"
	"log"
	"time"
)

//merge result
var ReduceMap string
var reduceChan chan string

//flag the reduce client
var ReduceFlag bool

type UserReduceFunc func(rm, res string) string

func InitReducer(userReduce UserReduceFunc) {
	reduceChan = make(chan string, 100)
	ReduceFlag = false
	go reduceResult(userReduce)
}

func reduceResult(userReduce UserReduceFunc) {
	for {
		select {
		case res := <-reduceChan:
			ReduceMap = userReduce(ReduceMap, res)
			log.Printf("routine number:%d \n", ExecutionRoutineNum)
		case <-time.After(time.Second * 10):
			//30 秒没有更新数据reduce data to client
			// log.Printf("routine number after:%d \n", ExecutionRoutineNum)
			if ExecutionRoutineNum == 0 {
				if ReduceMap != "" && ReduceFlag {
					reduceToClient()
				}
			}
		}
	}
}

/*
	call by server
	to send the res to client
*/
func reduceToClient() {
	tempMap := make(map[string]string)
	tempMap["result"] = ReduceMap
	command := &gocommand.Command{"reduce", "reduce", tempMap}
	temp := gocommand.EnCode(command.GetCommandString())
	gonet.ServerSend(temp)
	ReduceMap = ""
}

/*
	call by client
	to send the reduce command to client
*/
func CollectReduceToClient() {
	for _, config := range NodeConfig {
		tempMap := make(map[string]string)
		tempMap["result"] = ReduceMap
		command := &gocommand.Command{"reduce", "reduce", tempMap}
		temp := gocommand.EnCode(command.GetCommandString())
		msg := gonet.Message{config.Config["NodeAddr"], temp}
		gonet.Send(msg, func(msg string) {})
	}
}
