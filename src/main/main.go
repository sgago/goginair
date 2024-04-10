package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	templatesPath = "../templates/*"

	homePath     = "/"
	homeTemplate = "home.tmpl"

	healthzPath = "/healthz"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob(templatesPath)

	r.GET(healthzPath, healthz)
	r.GET(homePath, home)

	r.Run()
}

func home(c *gin.Context) {
	c.HTML(200, homeTemplate, gin.H{
		"title": "GoGinAir",
	})
}

func healthz(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
