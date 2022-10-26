package initutil

import (
	"context"
	"log"
	"pmdocker/src/env"
	"pmdocker/src/routes"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func CheckDocker(cli *client.Client, ctx context.Context) error {
	_, err := cli.Info(ctx)
	if err != nil {
		return err
	}

	return nil
}

// InitServer for init the server
func InitServer() {
	err := CheckDocker(env.GetCli(), env.GetCtx())
	if err != nil {
		log.Fatalln("Failed to connect to the docker daemon, Docker is installed ?")
	}
	c := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	routes.InitRoutes(c)

	c.Run(":" + env.PORT)
}
