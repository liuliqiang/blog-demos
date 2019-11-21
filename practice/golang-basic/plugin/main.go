package main

import (
	"fmt"
	"os"
	"plugin"

	"github.com/liuliqiang/log4go"
)

func main() {
	mod := "./plugin.so"

	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	add, err := plug.Lookup("Add")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	addFunc := add.(func(i, j int) int)

	log4go.Info("Result for 1 add 1 is: %d", addFunc(1, 1))
}
