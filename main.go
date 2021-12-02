package main

import (
	"net/http"

	c "github.com/bsdpunk/goApp/controller"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()
	c.LoadDetroit()
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("html/templates/*")

	router.Use(favicon.New("./favicon.ico"))
	//router.GET("/ping", func(c *gin.Context) {
	//	c.String(http.StatusOK, "Hello favicon.")
	//})
	//router.GET("/pingi2", func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong")
	//})

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})

	router.GET("/this", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"this.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":    "Home Page",
				"bullshit": "This",
			},
		)

	})

	//	m := autocert.Manager{
	//		Prompt:     autocert.AcceptTOS,
	//		HostPolicy: autocert.HostWhitelist("localhost"),
	//		Cache:      autocert.DirCache("/var/www/.cache"),
	//	}
	//	log.Fatal(autotls.RunWithManager(router, &m))
	// Start serving the application
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "wHaT yOu Do? INTERROBANG"})
		// Call the HTML method of the Context to render a template

	})
	router.Run()

}
