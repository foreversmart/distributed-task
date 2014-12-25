// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

package scheduler

import (
	"strconv"
	"sync"
	"log"
	"time"
	"fmt"

)

const (
	TypeSequence string = "sequence"//序列数据
	TypeStartEnd string = "startend"//起始数据
)

type Execution struct {
	method string
	dataItem map[string]string
	dataType string
}

type ExecutionUnit struct {
	method string
	key string
	value string
}

var executeChan chan *Execution
var executeUnitChan chan *ExecutionUnit
var executeControlChan chan bool

//performance analyze data
var timeTotal int64
var taskNum int64
var performance float64

func manager(userfunc func(key, vlaue string)) {
	//init
	executeChan = make(chan *Execution, 0)
	executeUnitChan = make(chan *ExecutionUnit, 1000)
	executeControlChan = make(chan bool, 100000)
	timeTotal = 0
	taskNum = 0
	performance = 0
	lock := &sync.Mutex{}

	//分配
	go func(){
		for{
			execute :=<- executeChan
			switch execute.dataType {
			case TypeSequence:
				for key,value := range execute.dataItem{
					executeUnitChan <- &ExecutionUnit{execute.method, key, value}
				}
			case TypeStartEnd:
				if execute.dataItem["start"]!="" &&
				 execute.dataItem["end"] != ""{
				 	start,err := strconv.ParseInt(execute.dataItem["start"], 10, 64)
				 	if err !=nil {
				 		log.Printf("manager, execute start end at start type wrong: %v\n", err)
				 	}
				 	end,err1 := strconv.ParseInt(execute.dataItem["end"], 10, 64)
				 	if err1 !=nil {
				 		log.Printf("manager, execute start end at end type wrong: %v\n", err)
				 	}
				 	for i := start; i <= end; i++ {
				 		value := strconv.FormatInt(i, 10)
				 		executeUnitChan <- &ExecutionUnit{execute.method, value, value}
				 	}

				}
			}

		}
	}()

	//起始并发量
	func(total int){
		for i := 0; i < total; i++ {
			executeControlChan <- true
		}

	}(10)

	//excute
	go func(){
		for{
			//并发控制
			<- executeControlChan
			// 执行单元
			unit := <- executeUnitChan
			go doExecute(userfunc, lock, unit)
		}
	}()
}

func AddExcution(method string, dataItem map[string]string, dataType string){
	executeChan <- &Execution{method, dataItem, dataType}
}

func doExecute(userfunc func(key, vlaue string), lock *sync.Mutex, unit *ExecutionUnit){

	t1 := time.Now()
	// UserExecute(unit.key, unit.value)
	res := userfunc(unit.key, unit.value)
	reduceChan <- res
	t2 := time.Now()

	//动态协程增量执行
	// change time unit to microsecond
	lock.Lock()
	timeTotal = timeTotal + int64(t2.Sub(t1))
	log.Printf("time total: %v \n", timeTotal)
	taskNum := taskNum + 1
	oldPerformance := performance
	performance := float64(taskNum) / float64(taskNum)
	if oldPerformance < performance {
		executeControlChan <- true
		executeControlChan <- true
	}
	if oldPerformance == performance {
		executeControlChan <- true
	}
	lock.Unlock()

}

/*
	define user functions
*/

func UserExecute(key, value string){
	fmt.Println("UserExecute :", key, value)
}