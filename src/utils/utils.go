package utils

import (
	"os"

	"github.com/docker/docker/client"
)

// GetVersion is the function to get the version of the Docker API
func GetVersion(cli *client.Client) string {
	ver := cli.ClientVersion()

	return "Docker API version : " + ver
}

// CheckRoot is the function for check if user are root
func CheckRoot() bool {
	file := "/root/tmpshellcheck.oklm"
	_, err := os.Create(file)
	if err != nil {
		os.Remove(file)
		return false
	}
	os.Remove(file)
	return true
}
