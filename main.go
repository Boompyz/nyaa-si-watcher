package main

import (
	"fmt"

	"github.com/Boompyz/nyaa-si-watcher/torrentoptions"
)

func main() {
	options, err := torrentoptions.GetAllOptions()
	if err != nil {
		panic(err.Error())
	}

	for _, option := range options {
		fmt.Println(option.GetId())
		fmt.Println(option)
	}
}
