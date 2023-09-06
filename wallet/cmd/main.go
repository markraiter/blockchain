package main

import (
	"os"

	"github.com/markraiter/blockchain/wallet/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
