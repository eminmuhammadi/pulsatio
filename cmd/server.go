package cmd

import (
	"errors"

	config "github.com/eminmuhammadi/pulsatio/config"
	grpc "github.com/eminmuhammadi/pulsatio/grpc"
	logger "github.com/eminmuhammadi/pulsatio/logger"

	cli "github.com/urfave/cli/v2"
)

func SecureServer(tlsFiles grpc.CertManager, address string, verify bool) error {
	return grpc.StartSecureServer(tlsFiles, address, verify)
}

func Server(address string) error {
	return grpc.StartServer(address)
}

func ServerCMD() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Run pulsatio server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "address",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "cert",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "key",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "ca",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "insecure-tls-verify",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "secure",
				Required: false,
			},
			&cli.IntFlag{
				Name:     "timeout",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			return serverFunc(ctx)
		},
	}
}

func serverFunc(ctx *cli.Context) error {
	cert := ctx.String("cert")
	key := ctx.String("key")
	ca := ctx.String("ca")
	address := ctx.String("address")
	tlsVerify := !ctx.Bool("insecure-tls-verify")
	secure := ctx.Bool("secure")
	timeout := ctx.Int("timeout")

	if timeout > 0 {
		config.ServerTimeout = timeout
	}

	logger.Printf("Timeout has been set to %d", config.ServerTimeout)

	if secure {
		tlsFiles := grpc.CertManager{
			CertFile: cert,
			KeyFile:  key,
			CAFile:   ca,
		}

		if cert == "" || key == "" || ca == "" {
			return errors.New("cert, key and ca files are required for secure connection")
		}

		return SecureServer(tlsFiles, address, tlsVerify)
	}

	return Server(address)
}
