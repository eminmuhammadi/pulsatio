# pulsatio

[![Go Reference](https://pkg.go.dev/badge/github.com/eminmuhammadi/pulsatio.svg)](https://pkg.go.dev/github.com/eminmuhammadi/pulsatio) [![Test](https://github.com/eminmuhammadi/pulsatio/actions/workflows/cli-test.yaml/badge.svg)](https://github.com/eminmuhammadi/pulsatio/actions/workflows/cli-test.yaml) [![goreleaser](https://github.com/eminmuhammadi/pulsatio/actions/workflows/release.yaml/badge.svg)](https://github.com/eminmuhammadi/pulsatio/actions/workflows/release.yaml)

pulsatio is a simple gRPC based remote terminal application.

## Installation

The easiest way to install pulsatio is to grab a precompiled binary from the [releases page](https://github.com/eminmuhammadi/pulsatio/releases).

### Go Install
```sh
$ go install github.com/eminmuhammadi/pulsatio@latest
```

### From binary
```sh
$ curl -LO https://github.com/eminmuhammadi/pulsatio/releases/download/v0.0.1/pulsatio_0.0.1_windows_amd64.tar.gz && tar -xvf pulsatio_0.0.1_windows_amd64.tar.gz
```

```sh
$ mv pulsatio /usr/local/bin
```

### From source
If you want to compile it yourself, you'll need to have a working Go environment with [version 1.16 or greater installed](https://golang.org/doc/install).

```sh
$ git clone https://github.com/eminmuhammadi/pulsatio.git && cd pulsatio
```

```sh
$ go build -o pulsatio && mv pulsatio /usr/local/bin
```

## Usage

The following commands are available:

### Server
Start a pulsatio server.

```sh
$ pulsatio server --address 0.0.0.0:9991 --secure --ca ./.tls/ca-cert.pem --key ./.tls/server-key.pem  --cert ./.tls/server-cert.pem --insecure-tls-verify --timeout 14000
2023/06/30 04:25:39 [pulsatio] Timeout has been set to 14000
2023/06/30 04:25:39 [pulsatio] gRPC server is using tls files: ./.tls/server-cert.pem, ./.tls/server-key.pem, ./.tls/ca-cert.pem
2023/06/30 04:25:39 [pulsatio] gRPC server is not using tls verification
2023/06/30 04:25:39 [pulsatio] gRPC server listening on tcp://[::]:9991
2023/06/30 04:25:49 [pulsatio] Received: pulsatio
```

### Client
Connect to a pulsatio server, and run remote commands.

```sh
$ pulsatio client --address 0.0.0.0:9991 --secure --ca ./.tls/ca-cert.pem --key ./.tls/server-key.pem  --cert ./.tls/server-cert.pem --insecure-tls-verify
pulsatio-0.0.1# pulsatio
NAME:
   pulsatio - pulsatio is a simple gRPC based remote terminal application.

USAGE:
   pulsatio [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   server   Run pulsatio server
   client   Run pulsatio client
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
pulsatio-0.0.1#
```

Note: Adding `PULSATIO_DEBUG=false` to the environment will disable debug logging.