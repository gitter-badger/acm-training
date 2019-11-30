package router

import (
	"github.com/gin-gonic/gin"
	user "github.com/uestc-acm/acm-training/controller"
	"github.com/uestc-acm/acm-training/db"
	"github.com/uestc-acm/acm-training/model"
)

// Binds all user related controller APIs with specific URL.
// This function maps the url with controller.user subpackage.
func registerUserRouter(r *gin.Engine) {
	db.DB().AutoMigrate(&model.User{})

	r.GET("/api/v1/users/", user.GetUsers)
	r.POST("/api/v1/users/", user.AddUser)
}
