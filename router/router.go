package router

import "github.com/gin-gonic/gin"

// Create function called to create a router for Restful APIs.
func Create() *gin.Engine {
	r := gin.Default()

	registerUserRouter(r)

	// TODO(ruinshe): initalize the router.
	// initRouter(r)

	return r
}
