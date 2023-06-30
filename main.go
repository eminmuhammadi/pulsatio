package main

import (
	"log"
	"os"

	cmd "github.com/eminmuhammadi/pulsatio/cmd"
	cli "github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	cmd.ServerCMD(),
	cmd.ClientCMD(),
}

var PULSATIO_VERSION = "0.0.1"

func main() {
	app := &cli.App{
		Name:                 "pulsatio",
		Usage:                "pulsatio is a simple gRPC client/server application",
		Version:              PULSATIO_VERSION,
		Commands:             Commands,
		EnableBashCompletion: true,
		Suggest:              true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
