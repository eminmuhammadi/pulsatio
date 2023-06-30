# pulsatio

gRPC connection between a client and a server to send and receive data.

#### Example 

##### Server
```sh
go run main.go server --address 0.0.0.0:9991 --secure --ca ./.tls/ca-cert.pem --key ./.tls/server-key.pem  --cert ./.tls/server-cert.pem --insecure-tls-verify
2023/06/30 03:35:58 [pulsatio] gRPC server is using tls files: ./.tls/server-cert.pem, ./.tls/server-key.pem, ./.tls/ca-cert.pem
2023/06/30 03:35:58 [pulsatio] gRPC server is not using tls verification
2023/06/30 03:35:58 [pulsatio] gRPC server listening on tcp://[::]:9991
2023/06/30 03:37:24 [pulsatio] Received ping message: Ping!
```

##### Client
```sh
$ go run main.go client --address 0.0.0.0:9991 --secure --ca ./.tls/ca-cert.pem --key ./.tls/server-key.pem  --cert ./.tls/server-cert.pem --insecure-tls-verify
Pong!
```