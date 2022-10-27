package main

import (
	"dockman/src/env"
	"dockman/src/initutil"
	"dockman/src/utils"
	"fmt"
	"log"
	"os"
)

func init() {
	if !utils.CheckIfLinux() {
		log.Fatalln("DockerManager are not compatible for Windows at this time")
	}
	if !utils.CheckRoot() {
		log.Fatalln("You are not Root")
	}
}

func main() {
	if len(os.Args) > 1 {
		a := string(os.Args[1])
		switch a {
		case "serve":
			initutil.InitServer()
		case "version":
			fmt.Println(utils.GetVersion(env.GetCli()))
		}
	} else {
		fmt.Println("Missing arguments, correct syntax :\n", "<serve> To start server\n", "<version> to get the Docker API version")
	}
}
