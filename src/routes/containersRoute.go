package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pmdocker/src/env"
	"pmdocker/src/utils"

	"github.com/gin-gonic/gin"
)

// Avatar is the struct to put json output
type Avatar struct {
	Img string `json:"gravatar_url"`
}

// DeleteAllContainersFromRoute is the function for delete all containers from a route
func DeleteAllContainersFromRoute(c *gin.Context) {
	//var list ListAllContainersJson
	_, err := utils.DeleteAllContainers(env.GetCtx(), env.GetCli())
	if err != nil {
		c.JSON(400, gin.H{
			"status": "No container running",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "All Containers deleted",
		})
	}
}

// CreateContainerFromRoute is the function for create container from a route
func CreateContainerFromRoute(c *gin.Context) {
	query := c.Param("containername")
	_, err := utils.CreateNewContainer(env.GetCtx(), env.GetCli(), query)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "Not Created",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "Created",
		})
	}
}

// ListAllContainersFromRoute is the function for list all containers from a route
func ListAllContainersFromRoute(c *gin.Context) {
	list, err := utils.ListAllContainers(env.GetCtx(), env.GetCli())
	if err != nil {
		c.JSON(400, gin.H{
			"data": "No container found or other",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": list,
		})
	}
}

// OpenContainerInATermFromRoute is a function for open a container in a terminal on the web
func OpenContainerInATermFromRoute(c *gin.Context) {
	query := c.Param("containername")
	err := utils.OpenContainerInATerm(query)
	if err != nil {
		c.JSON(400, gin.H{
			"status": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

// DeleteContainerFromRoute is the function for delete all containers from the route
func DeleteContainerFromRoute(c *gin.Context) {
	id := c.Param("id")
	err := utils.DeleteContainer(env.GetCtx(), env.GetCli(), id)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "Not deleted",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "Deleted",
		})
	}
}

// GetImageAvatarByNameFromRoute is the function for get an avatar from orgs of a docker image
func GetImageAvatarByNameFromRoute(c *gin.Context) {
	name := c.Param("name")
	resp, err := utils.GetImageAvatarByName(env.GetCtx(), env.GetCli(), name)
	var avatar Avatar
	err = json.Unmarshal([]byte(resp), &avatar)
	if err != nil {
		panic(err)
	}
	fmt.Println(avatar.Img)
	if err != nil {
		c.JSON(400, gin.H{
			"data": "No avatar found",
		})
	} else {
		c.String(http.StatusOK, avatar.Img)
	}
}

// RedirectToTermFromRoute is the function for redirect the user to GottyURL
func RedirectToTermFromRoute(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, env.GottyURL)
}
