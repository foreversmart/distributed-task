// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

/*
	Package provide method alloc task to Node server(base the Node execution speed)
*/
package scheduler

import (
	"distributed-task/gocommand"
	"distributed-task/gonet"
)

/*
	allocate data task with node performance
*/
func AllocateData(method string, commandType string, data map[string]string){
	switch commandType {
	case gocommand.TypeSequence:
		//TO-DO performance
		
		avg := Math.Ceil(len(data) / len(NodeConfig))
		for nodeName, config := NodeConfig {
			tempMap := make(map[string]string)
			for key,value := data {
				tempMap[key]=value
				if len(tempMap)>avg{
					command := &gocommand.Command{method, commandType, tempMap}
					commandString := command.GetCommandString()
					msg := &gonet.Message{value.Config["NodeAddr"], 
					gocommand.EnCode(commandString)}
					gonet.Send(msg)
					break
				}
			}
		}
	case gocommand.TypeStartEnd:
		//TO-DO performance
		avg := Math.Ceil(len(data) / len(NodeConfig))
		start,err := strconv.ParseInt(data["start"], 10, 64)
		if err !=nil {
			log.Printf("scheduler, execute start end at start type wrong: %v\n", err)
		}
		end,err1 := strconv.ParseInt(data["end"], 10, 64)
		if err1 !=nil {
			log.Printf("scheduler, execute start end at end type wrong: %v\n", err)
		}
		current := start
		for nodeName, config := NodeConfig {
			tempMap := make(map[string]string)
			if (current <= end) {
				tempMap["start"] = strconv.Itoa(current)
				current = current + avg
				if current > end {
					current = end
				}
				tempMap["end"] = strconv.Itoa(current)

				command := &gocommand.Command{method, commandType, tempMap}
				commandString := command.GetCommandString()
				msg := &gonet.Message{value.Config["NodeAddr"], 
				gocommand.EnCode(commandString)}
				gonet.Send(msg)
			}else{
				break
			}
		}
	}
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









