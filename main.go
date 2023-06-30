package main

import (
	"log"
	"os"

	cmd "github.com/eminmuhammadi/pulsatio/cmd"
	config "github.com/eminmuhammadi/pulsatio/config"
	cli "github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	cmd.ServerCMD(),
	cmd.ClientCMD(),
}

func main() {
	app := &cli.App{
		Name:                 config.Name,
		Usage:                config.Usage,
		Version:              config.Version,
		Commands:             Commands,
		EnableBashCompletion: true,
		Suggest:              true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
