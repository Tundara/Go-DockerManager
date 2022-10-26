package main

import (
	"fmt"
	"log"
	"os"
	"pmdocker/src/env"
	"pmdocker/src/initutil"
	"pmdocker/src/utils"
)

func init() {
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
