package initutil

import (
	"net/http"
	"pmdocker/src/env"
	"pmdocker/src/routes"

	"github.com/gin-gonic/gin"
)

//Utils for init all utils
func Utils() {
	c := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
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

	c.GET("/containers/delete", routes.DeleteAllContainersFromRoute)
	c.GET("/containers/list", routes.ListAllContainersFromRoute)
	c.GET("/container/create/:containername", routes.CreateContainerFromRoute)
	c.GET("/container/delete/:id", routes.DeleteContainerFromRoute)
	c.GET("/container/exec/:containername", routes.OpenContainerInATermFromRoute)
	c.GET("/container/pull/:repo/:name", routes.PullImageFromRoute)
	c.GET("/image/delete/:id", routes.DeleteImageFromRoute)
	c.GET("/images/list", routes.ListAllImagesFromRoute)
	c.GET("/images/delete", routes.DeleteAllImagesFromRoute)
	c.GET("/images/getavatar/:name", routes.GetImageAvatarByNameFromRoute)
	c.GET("/redirectToTerm", routes.RedirectToTermFromRoute)

	c.Run(":" + env.PORT)
}
