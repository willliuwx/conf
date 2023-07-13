package parse

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func configInit(filename string) (*ini.File, error) {
	fmt.Printf("Init config from %s\n", filename)
	file, err := ini.Load(filename)
	if err != nil {
		fmt.Println("config file read Error, please check file path correct:config.ini", err)
	}
	return file, err
}

func ConfigParse(filename string) {
	f, err := configInit(filename)
	if err != nil {
		return
	} else {
		if f != nil {
			x := f.Section("mongodb").Key("MongoUsername").String()
			fmt.Println(x)
		} else {
			fmt.Println("file hander is nil.")
		}
	}
}

type (
	MongoConf struct {
	}
	RabbitmqConf struct {
	}
	SystemConf struct {
	}
)
