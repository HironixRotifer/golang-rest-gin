package main

import (
	// "net/http"
	"app/middlewares"
	"app/routers"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	store := middlewares.SetSession()

	// Initializing group router
	beforeAuthorization := router.Group("/")
	afterAuthorization := router.Group("/")

	// Using middleware
	beforeAuthorization.Use(sessions.Sessions("session-name", store))
	afterAuthorization.Use(sessions.Sessions("session-name", store))
	afterAuthorization.Use(middlewares.AuthSession)

	// Routing
	routers.BeforeUserRouter(beforeAuthorization)
	routers.AfterUserRouter(afterAuthorization)

	// Load files
	router.LoadHTMLGlob("./ui/html/*")
	router.StaticFile("style.css", "./ui/static/style.css")
	router.StaticFile("img1.jpg", "./ui/static/img1.jpg")
	router.StaticFile("img2.jpg", "./ui/static/img2.jpg")
	router.StaticFile("img3.jpg", "./ui/static/img3.jpg")
	router.StaticFile("gif3.gif", "./ui/static/gif3.gif")
	router.StaticFile("gif1.gif", "./ui/static/gif1.gif")
	router.StaticFile("badGatwey.gif", "./ui/static/badGatwey.gif")

	// Start server
	router.Run(":8080")
}
