package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Exifer"
	app.Usage = "📸 A small CLI tool when use tweets your boast photos."
	app.Version = "0.0.1"

	app.Run(os.Args)
}
