package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(c *gin.Engine) {
	c.LoadHTMLGlob("./static/*.html")
	c.Static("/assets", "./static/assets/")
	c.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Posts",
		})
	})
	c.GET("/exec/:containername", func(c *gin.Context) {
		c.HTML(http.StatusOK, "exec.html", gin.H{
			"title":     "Posts",
			"container": c.Param("containername"),
		})
	})

	c.GET("/containers/delete", DeleteAllContainersFromRoute)
	c.GET("/containers/list", ListAllContainersFromRoute)
	c.GET("/container/create/:containername", CreateContainerFromRoute)
	c.GET("/container/delete/:id", DeleteContainerFromRoute)
	c.GET("/container/exec/:containername", OpenContainerInATermFromRoute)
	c.GET("/container/pull/:repo/:name", PullImageFromRoute)
	c.GET("/image/delete/:id", DeleteImageFromRoute)
	c.GET("/images/list", ListAllImagesFromRoute)
	c.GET("/images/delete", DeleteAllImagesFromRoute)
	c.GET("/images/getavatar/:name", GetImageAvatarByNameFromRoute)
	c.GET("/redirectToTerm", RedirectToTermFromRoute)
}
