package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	tmplPath = "../templates/*"

	indexPath = "/"
	indexTmpl = "index.tmpl"

	healthzPath = "/healthz"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob(tmplPath)

	r.GET(healthzPath, healthz)
	r.GET(indexPath, home)

	r.Run()
}

func home(c *gin.Context) {
	c.HTML(200, indexTmpl, gin.H{
		"title": "GoGinAir",
	})
}

func healthz(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
