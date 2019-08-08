package main

import (
	"os"

	"github.com/makotia/exifer/modules"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Exifer"
	app.Usage = "📸 A small CLI tool when use tweets your boast photos."
	app.Version = "1.0.0"

	app.Action = func(c *cli.Context) error {
		if c.Bool("image") {
			modules.Decode(c.Args().Get(0))
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "image, i",
			Usage: "Add watermark to image.",
		},
	}

	app.Run(os.Args)
}
