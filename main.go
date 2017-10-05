package main

import (
	"os"

	"github.com/denemebirki/denemebeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
