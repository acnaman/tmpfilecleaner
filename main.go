package main

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Target TargetConfig `yaml:"target"`
}

type TargetConfig struct {
	Folders []string `yaml:"folders"`
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

	// 対象ディレクトリ内のファイルを削除する
	for _, f := range targetFolders.Folders {
		os.Chdir(f)

		files, err := ioutil.ReadDir(f)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			os.RemoveAll(file.Name())
		}

	}
}
