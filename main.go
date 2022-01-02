package main

import (
	"fmt"
	"os"

	"github.com/mudler/golauncher/gui"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:   "golauncher",
		Author: "Ettore Di Giacinto",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:   "theme",
				EnvVar: "THEME",
			},
			&cli.StringFlag{
				Name:   "plugin-dir",
				EnvVar: "PLUGIN_DIR",
			},
		},
		Action: func(c *cli.Context) error {
			gui.Run(c.String("theme"), c.String("plugin-dir"))
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
