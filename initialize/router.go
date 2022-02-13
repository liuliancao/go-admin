package initialize

import (
	"go-admin/middleware"
	"go-admin/router"

	"github.com/gin-gonic/gin"
)

// init routes
func Routers() *gin.Engine {
	var Router = gin.Default()

	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup)
	}

	PrivateGroup := Router.Group()
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitApiRouter(PrivateGroup) // register api routes
		//router.InitJwtRouter(PrivateGroup) // register jwt routes
		//router.InitUserRouter(PrivateGroup) // register user
	}
}
