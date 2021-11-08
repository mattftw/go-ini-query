package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var (
	gitHash        = ""
	releaseVersion = ""
	gitState       = ""
)

var sectionFlag = &cli.StringFlag{
	Name:    "section",
	Aliases: []string{"s"},
	Usage:   "ini section to work within",
}

var propertyFlag = &cli.StringFlag{
	Name:    "property",
	Aliases: []string{"p"},
	Usage:   "get property",
}

func printOutput(c *cli.Context, output string) error {
	if c.Bool("witout-newline") {
		_, err := c.App.Writer.Write([]byte(output))
		return err
	} else {
		_, err := c.App.Writer.Write([]byte(output + "\n"))
		return err
	}
}

func getVersion() string {
	if gitState == "clean" && releaseVersion != "" {
		return releaseVersion
	}

	if gitState != "clean" {
		return fmt.Sprintf("%s (dirty)", gitHash)
	}

	return gitHash
}

var app = &cli.App{
	Name:    "go-ini-query",
	Usage:   "query and set ini values",
	Reader:  os.Stdin,
	Writer:  os.Stdout,
	Version: getVersion(),
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "file to process (specify '-' to read from stdin)",
		},
		&cli.BoolFlag{
			Name:    "without-newline",
			Aliases: []string{"w"},
			Usage:   "output text with newline",
		},
	},
	Commands: []*cli.Command{
		// handle retrieval of ini keys
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get a value",
			Flags: []cli.Flag{
				sectionFlag,
				propertyFlag,
			},
			Subcommands: []*cli.Command{},
			Action: func(c *cli.Context) error {

				cfg, err := readConfig(c)
				if err != nil {
					return errors.Wrapf(err, "unable to read config")
				}

				property := cfg.Section(c.String("section")).Key(c.String("property")).String()
				return printOutput(c, property)
			},
		},

		// handle addition of ini keys
		{
			Name:    "set",
			Aliases: []string{"s"},
			Usage:   "set a value",
			Flags: []cli.Flag{
				sectionFlag,
				propertyFlag,
				&cli.StringFlag{
					Name:    "value",
					Aliases: []string{"v"},
					Usage:   "set value",
				},
			},
			Subcommands: []*cli.Command{},
			Action: func(c *cli.Context) error {

				cfg, err := readConfig(c)
				if err != nil {
					return errors.Wrapf(err, "unable to read config")
				}

				cfg.Section(c.String("section")).Key(c.String("property")).SetValue(c.String("value"))
				return saveConfig(c, cfg)
			},
		},

		// handle deletion of ini keys
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete a value",
			Flags: []cli.Flag{
				sectionFlag,
				propertyFlag,
			},
			Subcommands: []*cli.Command{},
			Action: func(c *cli.Context) error {

				cfg, err := readConfig(c)
				if err != nil {
					return errors.Wrapf(err, "unable to read config")
				}

				cfg.Section(c.String("section")).DeleteKey(c.String("property"))

				if len(cfg.Section(c.String("section")).Keys()) == 0 {
					cfg.DeleteSection(c.String("section"))
				}

				return saveConfig(c, cfg)
			},
		},
	},
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
