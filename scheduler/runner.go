// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

package scheduler

import (
	"distributed-task/gocommand"
	"distributed-task/gonet"
	// "fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	Path string = "/Users/hong/goworkspace/src/distributed-task/scheduler/config.xml"
)

var LocalConfig map[string]string
var NodeConfig map[string]*Node

type Node struct {
	Config map[string]string
}

/*
	load config file & start client or server
	r define task, userfunc like map (define task itself),
	serverReduce server data reduce
*/
func Runner(r func(), userfunc UserExecuteFunc, serverReduce UserReduceFunc) {
	switch LocalConfig["ClientType"] {
	case "":
		log.Printf("config xml is wrong \n")
	case "client":
		log.Printf("starting client...\n")
		gonet.ClientInit()
		go r()
		// for {
		// 	var input string
		// 	if input == "stopclient" {
		// 		fmt.Scanln(&input)
		// 		fmt.Println("close client!")
		// 	}
		// }
		// gonet.ClientRead(func (msg string){
		// 	fmt.Printf("client read:", msg)
		// })
	case "server":
		log.Printf("starting server...\n")
		//init server reduce
		InitReducer(serverReduce)
		go manager(userfunc)
		go gonet.ServerRead(func(msg string) {
			log.Printf("recive msg:%v \n", msg)
			command := gocommand.GetCommand(msg)
			switch command.Method {
			case "reduce":
				//客户端发起规约
				ReduceFlag = true
			default:
				AddExcution(command.Method, command.Data, command.Type)
			}
		})
		gonet.ServerRun()

	}
}

/*
	load config
*/
func LoadConfig(path string) {

	if path == "" {
		path = Path
	}

	LocalConfig = make(map[string]string)
	NodeConfig = make(map[string]*Node)
	log.Printf("loading...success~ %v\n", LocalConfig)
	log.Printf("loading..remote nodes %v\n", NodeConfig)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("%s\n", err)
		panic(err)
	}
	content := string(f)
	contentLines := strings.Split(content, "\n")

	var isServerArea = false
	var tempName string

	for _, value := range contentLines {
		if strings.HasPrefix(value, "#") || value == "" {
			//注释
		} else {
			if value == "NodeServer" {
				//
				isServerArea = true
				continue
			}
			temps := strings.Split(value, ":")
			if len(temps) > 2 {
				for i := 2; i < len(temps); i++ {
					temps[1] = temps[1] + ":" + temps[i]
				}
			}
			if isServerArea {
				//server config area
				if temps[0] == "NodeName" {
					tempName = temps[1]
					NodeConfig[tempName] = &Node{}
					NodeConfig[tempName].Config = make(map[string]string)
				} else {
					NodeConfig[tempName].Config[temps[0]] = temps[1]
				}
			} else {
				//local config area
				LocalConfig[temps[0]] = temps[1]
			}

		}
	}
}
