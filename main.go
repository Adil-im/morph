package main

import (
	"fmt"
	"github.com/urfave/cli/v3"
	"morph/commands"
	"log"
	"os"
	"context"
)

func main() {
	cmd := &cli.Command{
		Name:     "morph",
		Usage:    "simple and elegant CLI tool for image format conversion",
		Commands: []*cli.Command{
			commands.img(),	
		},

	}
	
	if err := cmd.Run(context.Background(), os.Args); err !=  nil {
		log.Fatal(err)
	}		 		

}
