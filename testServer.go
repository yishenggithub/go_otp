package main

import "fmt"

type testServer struct {
	Server    genServer
	SupServer genServer
}

func (t testServer) init() {
	fmt.Println("init")
	t.startLink()
	go t.startLink()
}

func (t testServer) startLink() {
	fmt.Println("start_link")
}
