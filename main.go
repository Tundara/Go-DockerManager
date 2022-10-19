package main

import (
	"log"
	"os"
	"pmdocker/src/initutil"
	"pmdocker/src/utils"
)

func init() {
	if !utils.CheckRoot() {
		log.Fatal("You are not Root")
		os.Exit(1)
	}
}

func main() {
	initutil.Utils()
}
