package router

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"log"
)

// Create function called to create a router for Restful APIs.
func Create() *gin.Engine {
	r := gin.Default()

	if gin.Mode() != gin.DebugMode {
		log.Println("Setup SSL for the API server...")
		r.Use(secure.New(secure.Config{
			AllowedHosts:          []string{"acm.uestc.edu.cn"},
			SSLRedirect:           true,
			SSLHost:               "acm.uestc.edu.cn",
			STSSeconds:            315360000,
			STSIncludeSubdomains:  true,
			FrameDeny:             true,
			ContentTypeNosniff:    true,
			BrowserXssFilter:      true,
			ContentSecurityPolicy: "default-src 'self'",
			IENoOpen:              true,
			ReferrerPolicy:        "strict-origin-when-cross-origin",
			SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		}))
	}

	registerUserRouter(r)

	// TODO(ruinshe): initalize the router.
	// initRouter(r)

	return r
}
