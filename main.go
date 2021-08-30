package main

import (
	"os"

	"github.com/jupripratama/go-homework/app"
	"github.com/jupripratama/go-homework/cli"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
