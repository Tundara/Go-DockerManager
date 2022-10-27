package routes

import (
	"dockman/src/env"
	"dockman/src/utils"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// DeleteAllImagesFromRoute is the function for delete all docker images from routes
func DeleteAllImagesFromRoute(c *gin.Context) {
	resp, err := utils.DeleteAllImages(env.GetCtx(), env.GetCli())
	if err != nil {
		c.JSON(400, gin.H{
			"data": "No image found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": resp,
		})
	}
}

// ListAllImagesFromRoute is a function for list all docker images installed
func ListAllImagesFromRoute(c *gin.Context) {
	list, err := utils.ListAllImages(env.GetCtx(), env.GetCli())
	if err != nil {
		c.JSON(400, gin.H{
			"data": "No image found on your system",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": list,
		})
	}

}

// PullImageFromRoute is the function for pull a docker image from route
func PullImageFromRoute(c *gin.Context) {
	repo := c.Param("repo")
	name := c.Param("name")

	reader, err := utils.PullImage(env.GetCtx(), env.GetCli(), repo+"/"+name)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)
}

// DeleteImageFromRoute is the function for delete a docker image from route
func DeleteImageFromRoute(c *gin.Context) {
	id := c.Param("id")
	_, err := utils.DeleteImage(env.GetCtx(), env.GetCli(), id)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "not deleted",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "deleted",
		})
	}
}
