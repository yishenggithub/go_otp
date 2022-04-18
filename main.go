package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type App struct {
	Description  string
	Vsn          string
	Modules      string
	Registered   string
	Applications string
	Mod          string
}

// https://stackoverflow.com/questions/68562229/cannot-unmarshal-seq-into-string-in-go

func main() {
	app := App{}

	// 读配置yml，转成app结构
	yamlFile, err := ioutil.ReadFile("simple_cache.yml")

	if err != nil {
		log.Fatalf("yamlFile.Get err %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &app)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Println("conf", app)

}
