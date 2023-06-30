package cmd

import (
	"errors"

	grpc "github.com/eminmuhammadi/pulsatio/grpc"
	lib "github.com/eminmuhammadi/pulsatio/lib"
	"github.com/eminmuhammadi/pulsatio/terminal"

	cli "github.com/urfave/cli/v2"
)

func SecureClient(msg string, tlsFiles grpc.CertManager, address string, verify bool) (*lib.PongMessage, error) {
	conn, err := grpc.SecureConnect(tlsFiles, address, verify)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := grpc.Client(conn)

	return grpc.Ping(client, msg)
}

func Client(msg string, address string) (*lib.PongMessage, error) {
	conn, err := grpc.Connect(address)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := grpc.Client(conn)

	return grpc.Ping(client, msg)
}

func ClientCMD() *cli.Command {
	return &cli.Command{
		Name:  "client",
		Usage: "Run pulsatio client",
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
		},
		Action: func(ctx *cli.Context) error {
			for {
				command := terminal.Stdin()

				if command == "exit" {
					break
				}

				clientFunc(ctx, command)
			}

			return nil
		},
	}
}

func clientFunc(ctx *cli.Context, command string) error {
	if command == "" {
		return errors.New("command is required")
	}

	var err error
	var msg *lib.PongMessage

	cert := ctx.String("cert")
	key := ctx.String("key")
	ca := ctx.String("ca")
	address := ctx.String("address")
	tlsVerify := !ctx.Bool("insecure-tls-verify")
	secure := ctx.Bool("secure")
	if secure {
		tlsFiles := grpc.CertManager{
			CertFile: cert,
			KeyFile:  key,
			CAFile:   ca,
		}

		if cert == "" || key == "" || ca == "" {
			return errors.New("cert, key and ca files are required for secure connection")
		}

		msg, err = SecureClient(command, tlsFiles, address, tlsVerify)
	} else {
		msg, err = Client(command, address)
	}

	if err == nil {
		print(msg.Message)
	}

	return err
}
