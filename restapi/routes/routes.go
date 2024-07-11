package routes

import (
	"example.com/first-app/secureevent"
	"github.com/gin-gonic/gin"
)

func Registerevent(server *gin.Engine) {
	server.GET("/events", getevents)
	server.GET("/events/:id", getevent)
	authenticated := server.Group("/")
	authenticated.Use(secureevent.Authenticate)
	authenticated.POST("/events", secureevent.Authenticate, createevents)
	authenticated.PUT("/events/:id", updateevent)
	authenticated.DELETE("/events/:id", deleteevent)
	authenticated.POST("/events/:id/register", Eventregistration)
	authenticated.DELETE("/events/:id/register", Cancelregistration)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
