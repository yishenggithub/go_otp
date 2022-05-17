package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type app struct {
	Description  string
	Vsn          string
	Modules      []string
	Registered   []string
	Applications []string
	Mod          []string
}

func main() {
	testApp := app{}

	// 读配置yml,转成app结构
	// https://stackoverflow.com/questions/68562229/cannot-unmarshal-seq-into-string-in-go
	yamlFile, err := ioutil.ReadFile("simple_cache.yml")
	if err != nil {
		log.Fatalf("yamlFile.Get err %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &testApp)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// 启动gen_server1
	var t genServer
	t = &testServer{ServerName: "testServer"}
	testServerCh := make(chan string)
	go t.init(testServerCh)

	// 启动gen_server2
	var t2 genServer
	t2 = &testServer{ServerName: "testServer2"}
	testServerCh2 := make(chan string)
	go t2.init(testServerCh2)

	// 怎么和其他模块交互呢, 除了在main中, 1标准输入输出, 2tcp

	// 给t发消息
	testClientToChan(testServerCh)
	// testClientToServer(t2)  // t2还是个genServer型的interface, 不能用testServer来接

	// <-testServerCh  // 进展：会造成 fatal error: all goroutines are asleep - deadlock!
	// <-testServerCh2

}

// main在协程前结束了,参考书上的处理方法
// 如果可以做一个全局的映射表,比较方便给t发消息

func testClientToChan(c chan string) {
	// 给t发消息
	c <- "handle_info_test" // 进展：可以给c发消息了,但有概率main先结束,接收不到,
}

func testClientToServer(t testServer) {
	t.ch <- "chan from test_client"
}
