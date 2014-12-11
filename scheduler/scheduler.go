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

func allocateData(method string, commandType string, data map[string]string){
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
	}
}

func scheduler(){
	var TaskData[string]string
	var TaskType string
}









