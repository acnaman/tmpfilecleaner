package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
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
		cli.BoolFlag{
			Name:  "Yes, Y",
			Usage: "Unshow confirm message",
		},
		cli.BoolFlag{
			Name:  "trash, t",
			Usage: "move files into trash can",
		},
	}

	app.Action = func(c *cli.Context) error {
		DeleteFile(c.String("file"), c.Bool("Yes"), c.Bool("trash"))
		return nil
	}

	app.Run(os.Args)
}

// DeleteFile :設定ファイル情報を元にデータを削除する
func DeleteFile(f string, skipconfirm bool, trashmode bool) {
	configfile := f
	t := time.Now()

	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Fatalf("cannot read config file: %v", err)
	}
	var config Config
	yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	
	// ゴミ箱モードの時は退避用ディレクトリ(名前は日付)を作成する
	if trashmode {
		exe := os.Executable()
		exedir := filepath.Dir(exe)
		os.Chdir(exedir)
		if err := os.MkdirAll("_trash/" + t.Day(), 0777); err != nil {
			fmt.Println(err)
		}
	
	}

	targetFolders := config.Target

	// 対象ディレクトリ内のファイルを削除する
	for _, f := range targetFolders.Folders {
		fmt.Printf("ディレクトリ[" + f + "]内のファイルをすべて削除します。")
		if skipconfirm {
			fmt.Println("")
		} else {
			fmt.Println("よろしいですか？[Y/n]")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			if scanner.Text() != "Y" {
				continue
			}
		}
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
