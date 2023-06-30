# pulsatio

pulsatio is a simple gRPC based terminal application.

#### Example 

##### Server
```sh
$ go run main.go server --address 0.0.0.0:9991 --secure --ca ./.tls/ca-cert.pem --key ./.tls/server-key.pem  --cert ./.tls/server-cert.pem --insecure-tls-verify --timeout 14000
2023/06/30 04:25:39 [pulsatio] Timeout has been set to 14000
2023/06/30 04:25:39 [pulsatio] gRPC server is using tls files: ./.tls/server-cert.pem, ./.tls/server-key.pem, ./.tls/ca-cert.pem
2023/06/30 04:25:39 [pulsatio] gRPC server is not using tls verification
2023/06/30 04:25:39 [pulsatio] gRPC server listening on tcp://[::]:9991
2023/06/30 04:25:49 [pulsatio] Received: go run main.go
```

##### Client
```sh
$ go run main.go client --address 0.0.0.0:9991 --secure --ca ./.tls/ca-cert.pem --key ./.tls/server-key.pem  --cert ./.tls/server-cert.pem --insecure-tls-verify
pulsatio-0.0.1# go run main.go
NAME:
   pulsatio - pulsatio is a simple gRPC based terminal application.

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