package main

import (
	"fmt"
	"os"

	"github.com/woogea/id3_set/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
