package cmd

import (
	"errors"

	grpc "github.com/eminmuhammadi/pulsatio/grpc"
	lib "github.com/eminmuhammadi/pulsatio/lib"

	cli "github.com/urfave/cli/v2"
)

func SecureClient(tlsFiles grpc.CertManager, address string, verify bool) (*lib.PongMessage, error) {
	conn, err := grpc.SecureConnect(tlsFiles, address, verify)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := grpc.Client(conn)

	return grpc.Ping(client)
}

func Client(address string) (*lib.PongMessage, error) {
	conn, err := grpc.Connect(address)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := grpc.Client(conn)

	return grpc.Ping(client)
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
			return clientFunc(ctx)
		},
	}
}

func clientFunc(ctx *cli.Context) error {
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

		msg, err = SecureClient(tlsFiles, address, tlsVerify)
	} else {
		msg, err = Client(address)
	}

	if err == nil {
		println(msg.Message)
	}

	return err
}
