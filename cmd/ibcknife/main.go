package main

import (
	"fmt"
	"os"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/urfave/cli/v2"
)

func main() {
	figure.NewFigure("vIBC KNIFE", "standard", true).Print()
	fmt.Println()
	app := &cli.App{
		Name:                 "vibc_knife",
		Usage:                "A cross-platform command line utility for vIBC.",
		EnableBashCompletion: true,
		Flags:                []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:      "new",
				Usage:     "create a new vIBC project",
				Action:    newCmd,
				Args:      true,
				ArgsUsage: "<Project Name>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "Specify the project root",
						Value:   ".",
					},
					&cli.BoolFlag{
						Name:    "force",
						Aliases: []string{"f"},
						Usage:   "Force create the project, ignore whether the project root is already exists",
					},
					&cli.BoolFlag{
						Name:    "recurse",
						Aliases: []string{"r"},
						Usage:   "Specify whether the submodules should be recurse after the clone is created",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
