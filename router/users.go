package router

import (
	"github.com/gin-gonic/gin"
	user "github.com/uestc-acm/acm-training/controller"
)

// Binds all user related controller APIs with specific URL.
// This function maps the url with controller.user subpackage.
func registerUserRouter(r *gin.Engine) {
	r.GET("/api/v1/users/", user.GetUsers)
	r.POST("/api/v1/users/", user.AddUser)
}
