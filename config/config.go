package config

import "flag"

var (
	Host           = "localhost"
	ServerPort     = "8080"
	AlicePort      = "8081"
	BobPort        = "8082"
	Charlie        = "8083"
	TLScertServer  = "certs/hospital.crt"
	TLSkeyServer   = "certs/hospital.key"
	TLScertAlice   = "certs/Alice.crt"
	TLSkeyAlice    = "certs/Alice.key"
	TLScertBob     = "certs/Bob.crt"
	TLSkeyBob      = "certs/Bob.key"
	TLScertCharlie = "certs/Charlie.crt"
	TLSkeyCharlie  = "certs/Charlie.key"
	Port           = flag.String("own-addr", "localhost:1111", "ownAddr")
	CertFile       = flag.String("cert", "certs/alice.crt", "load cert file")
	KeyFile        = flag.String("key", "certs/alice.key", "load key file")
)
