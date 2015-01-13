// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

/*
	Package gocommand is command deal package for distributed-task
*/
package gocommand

import (
	"encoding/json"
	"log"
	"strings"
)

const (
	CommandSplit    string = "&**&"
	TypeSequence    string = "sequence"    //序列数据
	TypeStartEnd    string = "startend"    //起始数据
	TypeCurrentNext string = "currentnext" //当前和下一个
)

type Command struct {
	Method string
	Type   string
	Data   map[string]string
}

/* remove the command split*/
func DeCode(content string) string {
	if IsCommand(content) {
		temp := strings.Split(content, CommandSplit)[1]
		log.Printf("decode string: %s \n", temp)
		return temp
	} else {
		log.Printf("this is not command string: %s\n!", content)
		return ""
	}
}

/* add the command split*/
func EnCode(content string) string {
	return CommandSplit + content + CommandSplit
}

/*get the Command struct from the command json string*/
func GetCommand(content string) *Command {
	res := &Command{}
	if content != "" {
		if err := json.Unmarshal([]byte(content), &res); err != nil {
			panic(err)
		}
		return res
	} else {
		return nil
	}
}

/*get the command json string from the command*/
func (command *Command) GetCommandString() string {
	content, err := json.Marshal(command)
	if err != nil {
		log.Printf("get command string err: %s\n!")
		return ""
	}
	return string(content)
}

/*
	check the content is command or others
*/
func IsCommand(content string) bool {
	if strings.HasPrefix(content, CommandSplit) &&
		hasCommandSuffix(content) {
		return true
	} else {
		return false
	}
}

/*
	because many strings have empty char at tail
	so this method is the way find valid string len
*/
func hasCommandSuffix(command string) bool {
	lenCommand := len(command)
	lenSplit := len(CommandSplit)
	for i := 2 * lenSplit; i < lenCommand; i++ {
		if command[i:i+lenSplit] == CommandSplit {
			return true
		}
	}
	return false
}
