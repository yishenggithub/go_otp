package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type app struct {
	Description  string
	Vsn          string
	Modules      []string
	Registered   []string
	Applications []string
	Mod          []string
}

// https://stackoverflow.com/questions/68562229/cannot-unmarshal-seq-into-string-in-go

func main() {
	testApp := app{}

	// 读配置yml，转成app结构
	yamlFile, err := ioutil.ReadFile("simple_cache.yml")

	if err != nil {
		log.Fatalf("yamlFile.Get err %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &testApp)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Println("conf", testApp)

	fmt.Println(testApp.Registered)

	// 启动sup
	// 怎么根据string获得这个struct

	// 怎么把值放进去
	var t genServer
	t = &testServer{}

	ch := make(chan string)

	go t.init(ch)

	// 怎么交互测试效果, 想象好像能go了，要么标准输入输出，要么tcp
}
