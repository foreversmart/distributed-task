// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

/*
	Package provide method alloc task to Node server(base the Node execution speed)
*/
package scheduler

import (
	"github.com/foreversmart/distributed-task/gocommand"
	"github.com/foreversmart/distributed-task/gonet"
	"log"
	"strconv"
)

var nodeChan chan *Node

/*
	allocate data task with node performance
	//define the user collect data
*/
func AllocateData(method string, commandType string, data map[string]string, userCollect func(msg string)) {
	loadNodeChan()
	switch commandType {
	case gocommand.TypeSequence:
		//TO-DO performance
		count := 0
		avg := len(data) / len(NodeConfig)
		tempMap := make(map[string]string)
		for key, value := range data {
			tempMap[key] = value
			count++
			// log.Printf("datalength: %s \n", len(tempMap))
			if len(tempMap) >= avg {
				command := &gocommand.Command{method, commandType, tempMap}
				commandString := command.GetCommandString()
				content := gocommand.EnCode(commandString)
				sendNode(content, userCollect)
				// log.Printf("datalength1: %s \n", len(tempMapOther))
				tempMap = make(map[string]string)
				// log.Printf("datalength2: %s \n", len(tempMap))
			} else {
				if count >= len(data) {
					//at the end
					command := &gocommand.Command{method, commandType, tempMap}
					commandString := command.GetCommandString()
					content := gocommand.EnCode(commandString)
					sendNode(content, userCollect)
				}
			}
		}

	case gocommand.TypeStartEnd:
		//TO-DO performance

		start, err := strconv.ParseInt(data["start"], 10, 64)
		if err != nil {
			log.Printf("scheduler, execute start end at start type wrong: %v\n", err)
		}
		end, err1 := strconv.ParseInt(data["end"], 10, 64)
		if err1 != nil {
			log.Printf("scheduler, execute start end at end type wrong: %v\n", err)
		}

		avg := (end - start) / int64(len(NodeConfig))
		tempMap := make(map[string]string)

		for current := int64(start); current <= end; current = current + int64(avg) + 1 {
			tempMap["start"] = strconv.FormatInt(current, 10)
			tempEnd := current + int64(avg)
			if tempEnd > end {
				tempEnd = end
			}
			tempMap["end"] = strconv.FormatInt(tempEnd, 10)
			command := &gocommand.Command{method, commandType, tempMap}
			commandString := command.GetCommandString()
			content := gocommand.EnCode(commandString)
			sendNode(content, userCollect)
			tempMap = make(map[string]string)
		}
	case gocommand.TypeCurrentNext:
		command := &gocommand.Command{method, commandType, data}
		commandString := command.GetCommandString()
		content := gocommand.EnCode(commandString)
		sendNode(content, userCollect)
	}

}

/*
	load the nodeconfig in the nodechan sequence
*/
func loadNodeChan() {
	nodeChan = make(chan *Node, len(NodeConfig))
	for _, config := range NodeConfig {
		nodeChan <- config
	}
}

func sendNode(content string, userCollect func(msg string)) {
	node := <-nodeChan
	log.Printf("send msg to node:%s with message: %s\n", node.Config["NodeAddr"], content)
	msg := gonet.Message{node.Config["NodeAddr"], content}
	gonet.Send(msg, userCollect)
	nodeChan <- node
}

/*
	user define the task data and task function name
*/
// func Scheduler(){
// 	var taskData[string]string
// 	var taskType string
// 	var methodName string
// 	//user define data

// 	//
// 	allocateData(methodName, taskType, taskData)
// }
