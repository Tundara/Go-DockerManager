package utils

import (
	os "os/exec"
	"pmdocker/src/env"
	"syscall"

	"github.com/pbar1/pkill-go"
)

// EndGotty is a function for end gotty process
func EndGotty() error {
	_, err := pkill.Pkill("gotty", syscall.SIGTERM)
	if err != nil {
		return err
	}

	return nil
}

// OpenContainerInATerm is a function for open a docker container in a term
func OpenContainerInATerm(contID string) error {
	feedback := EndGotty()
	if feedback != nil {
		panic(feedback)
	}
	cmd := os.Command(
		"./gottydir/gotty",
		"-w",
		"--port",
		env.GottyPORT,
		"docker",
		"exec",
		"-it",
		contID,
		"bash",
	)
	err := cmd.Start()

	return err
}
