package infect

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func Version() string {
	return "0.0.1"
}
func Execute() {

	app := cli.NewApp()
	app.Name = "Infect"
	app.Usage = "infect"
	app.Version = Version()
	app.Flags = []cli.Flag{
		// cli.BoolFlag{"edit, e", "edit a given shortcut"},
		// cli.BoolFlag{"search, s", "search-ack/ag it"},
		// cli.BoolFlag{"print, p", "display existing shortcuts"},
		cli.BoolFlag{"debug, d", "show all the texts"},
		// cli.StringFlag{"flags, f", "-i", "flags to pass to ag"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "install",
			Usage: "install dependencies",
			Action: func(c *cli.Context) {
				install(c)
			},
		},
	}
	app.Action = func(c *cli.Context) {

		if c.Bool("debug") {
			fmt.Printf("Context %#v\n", c)
		}

		switch true {
		default:
			install(c)
		}
	}

	app.Run(os.Args)
}
