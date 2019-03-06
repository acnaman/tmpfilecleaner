package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Target TargetConfig `yaml:"target"`
}

type TargetConfig struct {
	folders string `yaml:"folders"`
}

func main() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	var config Config
	yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	targetFolders := config.Target

	fmt.Println(targetFolders.folders)
}
