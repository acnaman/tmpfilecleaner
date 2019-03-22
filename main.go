package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"

	yaml "gopkg.in/yaml.v2"
)

// Config :YAMLの全体
type Config struct {
	Target TargetConfig `yaml:"target"`
}

// TargetConfig :削除対象についての情報
type TargetConfig struct {
	Folders []string `yaml:"folders"`
}

func main() {
	app := cli.NewApp()
	app.Name = "TMP File Cleaner"
	app.Usage = ""
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "config.yaml",
			Usage: "Config file(YAML)",
		},
	}

	app.Action = func(c *cli.Context) error {
		DeleteFile(c.String("file"))
		return nil
	}

	app.Run(os.Args)
}

// DeleteFile :設定ファイル情報を元にデータを削除する
func DeleteFile(f string) {
	configfile := f

	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Fatalf("cannot read config file: %v", err)
	}
	var config Config
	yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	targetFolders := config.Target

	// 対象ディレクトリ内のファイルを削除する
	for _, f := range targetFolders.Folders {
		fmt.Println("ディレクトリ[" + f + "]内のファイルをすべて削除します")
		os.Chdir(f)

		files, err := ioutil.ReadDir(f)
		if err != nil {
			log.Fatal(err)
			fmt.Println("warning: cannot read directory: " + f)
		} else {
			for _, file := range files {
				fmt.Printf("削除中：" + file.Name() + "\r")
				os.RemoveAll(file.Name())
				fmt.Printf("                                           \r")
			}
			fmt.Println("")
		}
	}
}
