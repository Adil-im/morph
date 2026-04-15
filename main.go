package main

import (
	"fmt"
	"github.com/urfave/cli/v3"
	"morph/commands"
)

func main() {
	cmd := &cli.Command{
		Name:     "morph",
		Usage:    "simple and elegant CLI tool for image format conversion",
		Commands: []*cli.Command{},
	}
}
