package main

import (
	"fmt"
	"net/http"
	"strconv"

	cd "github.com/bsdpunk/goApp/controller"
	m "github.com/bsdpunk/goApp/models"
	p "github.com/bsdpunk/goApp/page"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

var router *gin.Engine

var db = cd.GetDB()

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()
	cd.LoadDetroit()
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

	router.GET("/show", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"this.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":    "Home Page",
				"bullshit": "Single Crime",
			},
		)

	})
	router.GET("/show/:id", func(c *gin.Context) {

		fmt.Println(cd.GetDetroit(c.Param("id")))
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"show.html",
			// Pass the data that the page uses (in this case, 'title')

			gin.H{
				"title":    "Home Page",
				"bullshit": "Single Crime",
				"crime":    cd.GetDetroit(c.Param("id")),
			},
		)

	})

	router.GET("/create", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"create.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":    "Home Page",
				"bullshit": "Create",
			},
		)

	})
	router.GET("/map", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"map.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":    "Home Page",
				"bullshit": "Map",
			},
		)

	})
	router.GET("/map2", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"map2.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":    "Home Page",
				"bullshit": "Map",
			},
		)

	})

	router.GET("/page", func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "3"))
		var detroits m.Detroits

		paginator := p.Paging(&p.Param{
			DB:      db,
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"rownum"},
			ShowSQL: true,
		}, &detroits)
		c.JSON(200, paginator)
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
