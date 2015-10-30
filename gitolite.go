package main

import (
	"fmt"
	"os"

	"github.com/jsternberg/gitolite"
)

func realMain() int {
	config := gitolite.DefaultConfig()
	server := gitolite.New(config)

	if err := server.ListenAndServe(":1997"); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(realMain())
}
