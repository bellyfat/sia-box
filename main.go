package main

import (
	"github.com/jay-dee7/sia-box/cmd"
	"github.com/jay-dee7/sia-box/config"
)

func init() {
	config.Setup()
}

func main() {
	cmd.Execute()
}
