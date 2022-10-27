package env

import (
	"context"

	"github.com/docker/docker/client"
)

// PORT is the variable who contain port of Gin
const PORT string = "1337"

// GottyPORT is the variable who contain gotty port
const GottyPORT string = "8181"

// GottyURL is the variable who contain gotty url
const GottyURL string = "http://localhost:" + GottyPORT

// GetCtx is the functions for init the context
func GetCtx() context.Context {
	ctx := context.Background()
	return ctx
}

// GetCli is the function for init Cli
func GetCli() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return cli
}
