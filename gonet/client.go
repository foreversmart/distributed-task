// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

/*
	Package provide network socket and message define (include server and client)
*/
package gonet

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"
)

type GoClient struct {
	sendChan    chan string
	receiveChan chan Message
}

var ClientMap map[string]*GoClient

func ClientInit() {
	ClientMap = make(map[string]*GoClient)
}

/*
	send message
*/
func Send(message Message, callback func(msg string)) {

	remote := message.Addr
	if ClientMap[remote] == nil {
		ClientMap[remote] = &GoClient{sendChan: make(chan string, 100), receiveChan: make(chan Message, 100)}
		//client keeper
		go ClientMap[remote].clientKeeper(remote)
		//bind hand client read
		// var a = func(msg Message){
		// 	fmt.Println("get msg", msg.Content)
		// }
		go ClientMap[remote].ClientRead(callback)
		ClientMap[remote].sendChan <- message.Content
		fmt.Println("nil")
	} else {
		//
		ClientMap[remote].sendChan <- message.Content
		fmt.Println("not nil")
	}
}

/*
	start a dial to server
	and keep communication
*/
func (goClient *GoClient) clientKeeper(remote string) {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		m := recover()
		fmt.Println(m)
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		sendChan <- "&**& quit this client &**&" //退出协程并且删掉主对象
	}()

	tcpAddr, err := net.ResolveTCPAddr("tcp4", remote)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// conn.SetKeepAlive(true)
	checkError(err)

	var msg = make([]byte, 1024)

	// start send routine 降低读写冲突
	go func() {
		for {
			sendMsg := <-goClient.sendChan
			if sendMsg != "&**& quit this client &**&" {
				fmt.Println("client sengding:", sendMsg)
				_, err = conn.Write([]byte(sendMsg))
				checkError(err)
			} else {
				//崩溃退出关闭发送协程
				fmt.Println("quit client Keeper success")
				goClient = nil //删除掉主对象
				return
			}
		}
	}()

	//hear beat
	go func() {
		for {
			time.Sleep(time.Second * 60)
			if goClient != nil {
				fmt.Println("heart")
				goClient.sendChan <- "&**& !heart beat! &**&"
			} else {
				return
			}
		}
	}()

	// listen message
	for {
		_, err = conn.Read(msg)
		checkError(err)
		var tempMsg = Message{remote, string(msg)}
		goClient.receiveChan <- tempMsg
		msg = make([]byte, 1024)
	}

}

func (goClient *GoClient) ClientRead(f func(msg string)) {
	tempbytes := bytes.Buffer{}
	for {
		tempMsg := <-goClient.receiveChan
		// fmt.Println("client get:", tempMsg.Content)
		/*
			user custom function deal message when the client

			receive message
		*/
		tempbytes.WriteString(tempMsg.Content)
		if strings.Count(tempbytes.String(), "&**&") >= 2 {
			tempArr := strings.SplitN(tempbytes.String(), "&**&", 3)
			f(tempArr[1])
			tempbytes = bytes.Buffer{}
			tempbytes.WriteString(tempArr[3])
		} else {

		}
		// tempMsg.Content = tempMsg.Content + "2"
		// Send(tempMsg)
	}
}
