package main

import (
	"fmt"

	"github.com/willliuwx/conf/parse"
)

func main() {
	fmt.Println("Conf package starts.")
}

func init() {
	parse.ConfigParse("./config.ini")
}
