package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Target TargetConfig `yaml:"target"`
}

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
		configfile := c.String("file")

		data, err := ioutil.ReadFile(configfile)
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
		return nil
	}

	app.Run(os.Args)
}
