package main

import (
	"flag"

	// This Service
	"hello-gokit/svc/server"
	"hello-gokit/svc/server/cli"

	_ "hello-gokit/di"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	server.Run(cli.Config)
}
