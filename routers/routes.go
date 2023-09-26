package routers

import (
	"app/controllers"
	// session "app/middleware"

	// "github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func BeforeUserRouter(g *gin.RouterGroup) {
	g.GET("/", controllers.HomeGetHanler)
	g.GET("/login", controllers.LoginGetHandler)
	g.GET("/reg", controllers.RegistrationGetHandler)
	g.GET("/help", controllers.HelpGetHandler)
	g.GET("/error", controllers.ErrorGetHandler)
	g.POST("/login", controllers.LoginPostHandler)
	g.POST("/reg", controllers.RegistrationPostHandler)

}

func AfterUserRouter(g *gin.RouterGroup) {
	g.GET("/dashboard", controllers.DashboardGetHandler)
	g.GET("/logout", controllers.LogoutPostHandler)

}
