package utils

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//DeleteImage is the function for delete a docker image by id
func DeleteImage(ctx context.Context, cli *client.Client, imageID string) ([]types.ImageDeleteResponseItem, error) {
	resp, err := cli.ImageRemove(ctx, imageID, types.ImageRemoveOptions{Force: true})
	return resp, err
}

//GetImageAvatarByName is the function for get an avatar from the orgs of docker image
func GetImageAvatarByName(ctx context.Context, cli *client.Client, name string) ([]byte, error) {
	resp, err := http.Get("https://hub.docker.com/v2/orgs/" + name)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body, err
}

//DeleteAllImages is the function for delete all docker images
func DeleteAllImages(ctx context.Context, cli *client.Client) ([]types.ImageSummary, error) {
	listdeleted, err := ListAllImages(ctx, cli)
	if err != nil {
		log.Println("No image on your system")
		return listdeleted, err
	}
	for _, list := range listdeleted {
		_, err := DeleteImage(ctx, cli, list.ID)
		if err != nil {
			return listdeleted, err
		}
	}
	log.Println("All images deleted")

	return listdeleted, err
}

//ListAllImages is the function for list all docker images
func ListAllImages(ctx context.Context, cli *client.Client) ([]types.ImageSummary, error) {
	resp, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if len(resp) < 1 {
		err := errors.New("Nope")
		return resp, err
	}
	return resp, err
}

// PullImage is the function for pull a docker image
func PullImage(ctx context.Context, cli *client.Client, image string) (io.ReadCloser, error) {
	reader, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})

	return reader, err
}
