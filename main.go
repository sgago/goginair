package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed templates
	embeddedFiles embed.FS
)

const (
	templPath = "templates/*"

	indexPath  = "/"
	indexTempl = "index.tmpl"

	healthzPath = "/healthz"
)

func main() {
	templ := template.
		Must(template.New("").
			ParseFS(embeddedFiles, templPath))

	r := gin.Default()

	r.SetHTMLTemplate(templ)

	r.GET(healthzPath, healthz)
	r.GET(indexPath, index)

	r.Run()
}

func index(c *gin.Context) {
	c.HTML(200, indexTempl, gin.H{
		"title": "GoGinAircccc",
	})
}

func healthz(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
