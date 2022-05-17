package main

import (
	"fmt"
)

type testServer struct {
	ch         chan string
	ServerName string
	Server     genServer
	SupServer  genServer
}

func (t *testServer) init(ch chan string) {
	// 这里的放入操作可以改变t吗
	t.ch = ch
	fmt.Println("init")

	t.startLink(ch)
}

func (t *testServer) startLink(ch chan string) {
	fmt.Println("start_link")
	t.handleInfo(ch)
}

func (t *testServer) handleInfo(ch chan string) {
	fmt.Println("handleInfo")
	var info string
	info = <-ch
	fmt.Println("test")
	fmt.Println(info)
	switch "handle_info_test" {
	case "handle_info_test":
		fmt.Println("handle_info_test")
	}

}
