// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

/*
	only server use
	this file retask is retask to client
*/
package scheduler

import (
	"github.com/foreversmart/distributed-task/gocommand"
	"github.com/foreversmart/distributed-task/gonet"
)

/*
	call by server
	send task to client to retask
*/
func Retask(task string) {
	tempMap := make(map[string]string)
	tempMap["task"] = task
	command := &gocommand.Command{"retask", "retask", tempMap}
	temp := gocommand.EnCode(command.GetCommandString())
	gonet.ServerSend(temp)
}
