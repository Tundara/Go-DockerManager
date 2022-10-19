package utils

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// CreateNewContainer is the function for create a container
func CreateNewContainer(ctx context.Context, cli *client.Client, imageName string) (string, error) {
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        imageName,
		Cmd:          []string{"sh"},
		AttachStdin:  true,
		AttachStdout: true,
		OpenStdin:    true,
		Tty:          false,
	}, nil, nil, nil, "")
	if err != nil {
		return "nothing", err
	}

	strt := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if strt != nil {
		panic(err)
	}
	log.Printf("Container [%v] Created \n", resp.ID)
	return string(resp.ID), err
}

// DeleteContainer is the function for delete a container
func DeleteContainer(ctx context.Context, cli *client.Client, containerID string) error {
	err := cli.ContainerKill(ctx, string(containerID), "KILL")
	return err
}

// ListAllContainers is the function for list all active containers
func ListAllContainers(ctx context.Context, cli *client.Client) ([]types.Container, error) {
	listoutSec, _ := cli.ContainerList(ctx, types.ContainerListOptions{})
	if len(listoutSec) < 1 {
		err := errors.New("Nope")
		return listoutSec, err
	}
	return listoutSec, nil
}

// DeleteAllContainers is the function for delete all containers
func DeleteAllContainers(ctx context.Context, cli *client.Client) ([]types.Container, error) {
	listoutSec, err := ListAllContainers(ctx, cli)
	if err != nil {
		fmt.Println("No container are running")
		return listoutSec, err
	}
	for _, list := range listoutSec {
		DeleteContainer(ctx, cli, list.ID)
	}
	log.Println("All containers deleted")

	return listoutSec, err
}

// ContainerInfo is the function for get information about container
func ContainerInfo(ctx context.Context, cli *client.Client) {
	//do something
}
