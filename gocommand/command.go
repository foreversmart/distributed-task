// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

/*
	Package gocommand is command deal package for distributed-task
*/
package gocommand

import (
	"strings"
	"log"
	"encoding/json"
)

const (
	CommandSplit string = "&**&"
	TypeSequence string = "sequence"//序列数据
	TypeStartEnd string = "startend"//起始数据
)


type Command struct {
	Method string
	Type string
	Data map[string] string
}

/* remove the command split*/
func Decode (content string) (string) {
	if IsCommand(content) {
		return strings.TrimPrefix(strings.TrimSuffix(content, CommandSplit), CommandSplit)
	}else {
		log.Printf("this is not command string: %s\n!", content)
		return ""
	}
}

/* add the command split*/
func Encode (content string) string {
	return CommandSplit + content + CommandSplit
}

/*get the Command struct from the command json string*/
func GetCommand(content string) *Command {
	res := &Command{}
	if err := json.Unmarshal([]byte(content), &res); err != nil {
        panic(err)
    }
    return res
}

/*get the command json string from the command*/
func (command *Command) GetCommandString() string {
	content, err := json.Marshal(command)
	if err!=nil {
		log.Printf("get command string err: %s\n!")
		return ""
	}
	return string(content)
}

/*
	check the content is command or others
*/
func IsCommand (content string) bool {
	if strings.HasPrefix(content, CommandSplit) &&
	 strings.HasSuffix(content, CommandSplit) {
		return true
	}else{
		return false
	}
}
