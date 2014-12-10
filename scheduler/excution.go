// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

package scheduler

import (

)

const (
	TypeSequence string = "sequence"//序列数据
	TypeStartEnd string = "startend"//起始数据
)

type Execution struct {
	dataItem map[string]string
	dataType string
}

type Unit struct {
	key string
	value string
}

var executeChan chan *Execution
var executeUnitChan chan *Unit
var executeControlChan chan bool

func manager() {
	//分配
	go func(){
		for{
			execute <- executeChan
			switch dataType {
			case TypeSequence:
				for key,value := range execute.dataItem{
					executeUnitChan <- &executeUnit{key, value}
				}
			case TypeStartEnd

			}

		}
	}()
	//动态协程增量执行
	go func(){
		for{
			go Excute()

		}
	}
}

func Execute(){

}

func 