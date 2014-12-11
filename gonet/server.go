// Copyright 2014 Ben-Kuang. All rights reserved.
// Use of this source code is governed by The MIT License
// license that can be found in the LICENSE file.

package gonet

import (
	"fmt"
	"net"
	"time"
	// "strconv"
	"os"
)

// the server can have one client to work
// TO-DO 
var receiveChan chan string
var sendChan chan string

func ServerRun() {
	service := ":1256"
	receiveChan = make(chan string, 20)
	sendChan = make(chan string, 20)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		go serverKeeper(conn)
	}
}

func ServerRead(f func(msg string)) {
	for {
		temp := <-receiveChan
		fmt.Println("server get:",temp)
		/*
			user custom function to deal the message when
			server receive message
		*/
		f(temp)
		// time.Sleep(time.Second * 10)
		// temp = temp + "1"
		// ServerSend(temp)
	}
}

/*
	send a content to the client
*/
func ServerSend(content string) {
	sendChan <- content
}

/*
	keep routine for write and read
*/
func serverKeeper(conn *net.TCPConn) {
	// conn.SetKeepAlive(true)
	request := make([]byte, 128) // set maxium request length to 128KB to prevent flood attack
	defer func() {
		sendChan <- "&**& quit this server &**&"//退出协程并且删掉主对象  
	}()  // close connection before exit

	// send
	go func () {
		for {
			sendData := <-sendChan
			if sendData != "&**& quit this server &**&"{
				fmt.Println("server sending...", sendData)
				_, err := conn.Write([]byte(sendData))
				checkError(err)
			}else{
				//崩溃退出关闭发送协程
				conn.Close()
				fmt.Println("quit server Keeper success")
				return
			}
		}
	}()

	//read
	for {
		conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
		read_len, err := conn.Read(request)

		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("this is EOF")
				break
			}
			fmt.Println(err)
			fmt.Println(err.Error())
			//break
		}

		if read_len == 0 {
			break // connection already closed by client
		} else {
			receiveChan <- string(request)
		}
		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
		// fmt.Println(err.Error())
	}
}

