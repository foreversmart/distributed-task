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
)

const (
	CommandSplit string = "&**&"
)

type Command struct {
	Method string
	Data map[string] string
}

/* remove the command split*/
func Decode (content string) (string) {
	if IsCommand(content) {
		return TrimPrefix(TrimSuffix(content, CommandSplit), CommandSplit)
	}else {
		log.Printf("this is not command string: %s\n!", content)
		return ""
	}
}

/* add the command split*/
func Encode (content string) string {
	reuturn CommandSplit + content + CommandSplit
}

/*get the Command struct from the command json string*/
func GetCommand(content string) *Command {
	res := &Command{}
	if err := json.Unmarshal([]byte(str), &res); err != nil {
        panic(err)
    }
    return res
}

/*get the command json string from the command*/
func (command *Command) GetCommandString() string {
	return json.Marshal(command)
}

/*
	check the content is command or others
*/
func IsCommand (content string) bool {
	if strings.HasPrefix(CommandSplit) && strings.HasSuffix(CommandSplit) {
		return true
	}else{
		return false
	}
}
