package main

import "github.com/NoF0rte/oh-my-posh/cli"

var (
	Version = "development"
)

func main() {
	cli.Execute(Version)
}
