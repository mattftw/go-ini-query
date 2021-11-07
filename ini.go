package main

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gopkg.in/ini.v1"
)

func readConfig(c *cli.Context) (*ini.File, error) {

	if c.String("file") == "-" { // read from stdin
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to read from stdin")
		}

		cfg, err := ini.Load(bytes)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to process ini from stdin")
		}

		return cfg, nil
	} else { // read from file
		cfg, err := ini.Load(c.String("file"))
		if err != nil {
			return nil, errors.Wrapf(err, "unable to read ini file")
		}
		return cfg, nil
	}

}

func saveConfig(c *cli.Context, cfg *ini.File) error {
	if c.String("file") == "" {
		_, err := cfg.WriteTo(os.Stdout)
		return err
	} else {
		return cfg.SaveTo(c.String("file"))
	}
}
