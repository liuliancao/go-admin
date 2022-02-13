package routers

import (
	v1 "go-admin/api/v1"
	_ "go-admin/docs"
	jwt "go-admin/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello, golang",
	})
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.POST("/auth", GetAuth)

	r.POST("/auth", v1.GetAuth)
	r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/hello", sayHello)

		apiv1.GET("/product/:id", v1.GetProduct)
		apiv1.POST("/product", v1.AddProduct)
		apiv1.GET("/products", v1.GetProducts)
		apiv1.PUT("/product/:id", v1.EditProduct)
		apiv1.DELETE("/product/:id", v1.DeleteProduct)
		apiv1.GET("/productuser", v1.GetProductUsers)
		apiv1.POST("/productuser", v1.AddProductUser)

		apiv1.GET("/app/:id", v1.GetApp)
		apiv1.POST("/app", v1.AddApp)
		apiv1.GET("/apps", v1.GetApps)
		apiv1.PUT("/app/:id", v1.EditApp)
		apiv1.DELETE("/app/:id", v1.DeleteApp)
		apiv1.GET("/appuser", v1.GetAppUsers)
		apiv1.POST("/appuser", v1.AddAppUser)

		apiv1.POST("/appcluster", v1.AddAppCluster)

		apiv1.GET("/appenv/:id", v1.GetAppEnv)
		apiv1.POST("/appenv", v1.AddAppEnv)
		apiv1.GET("/appenvs", v1.GetAppEnvs)
		apiv1.PUT("/appenv/:id", v1.EditAppEnv)
		apiv1.DELETE("/appenv/:id", v1.DeleteAppEnv)

		apiv1.GET("/department/:id", v1.GetDepartment)
		apiv1.POST("/department", v1.AddDepartment)
		apiv1.GET("/departments", v1.GetDepartments)
		apiv1.PUT("/department/:id", v1.EditDepartment)
		apiv1.DELETE("/department/:id", v1.DeleteDepartment)

		apiv1.GET("/role/:id", v1.GetRole)
		apiv1.POST("/role", v1.AddRole)
		apiv1.GET("/roles", v1.GetRoles)
		apiv1.PUT("/role/:id", v1.EditRole)
		apiv1.DELETE("/role/:id", v1.DeleteRole)

		apiv1.GET("/cluster/:id", v1.GetCluster)
		apiv1.POST("/cluster", v1.AddCluster)
		apiv1.GET("/clusters", v1.GetClusters)
		apiv1.PUT("/cluster/:id", v1.EditCluster)
		apiv1.DELETE("/cluster/:id", v1.DeleteCluster)

		apiv1.GET("/clusteruser", v1.GetClusterUsers)
		apiv1.POST("/clusteruser", v1.AddClusterUser)

		apiv1.POST("/cluster/host", v1.AddClusterHost)

		apiv1.GET("/host/:id", v1.GetHost)
		apiv1.POST("/host", v1.AddHost)
		apiv1.GET("/hosts", v1.GetHosts)
		apiv1.PUT("/host/:id", v1.EditHost)
		apiv1.DELETE("/host/:id", v1.DeleteHost)

		apiv1.GET("/guarders", v1.GetGuarders)
		apiv1.POST("/guarder", v1.AddGuarder)
		apiv1.PUT("/guarder/:id", v1.EditGuarder)
		apiv1.DELETE("/guarder/:id", v1.DeleteGuarder)

		apiv1.GET("/tag/:id", v1.GetTag)
		apiv1.POST("/tag", v1.AddTag)
		apiv1.GET("/tags", v1.GetTags)
		apiv1.PUT("/tag/:id", v1.EditTag)
		apiv1.DELETE("/tag/:id", v1.DeleteTag)

		apiv1.GET("/user/:id", v1.GetUser)
		apiv1.POST("/user", v1.AddUser)
		apiv1.GET("/users", v1.GetUsers)
		apiv1.PUT("/user/:id", v1.EditUser)
		apiv1.DELETE("/user/:id", v1.DeleteUser)

		apiv1.GET("/resource/:id", v1.GetResource)
		apiv1.POST("/resource", v1.AddResource)
		apiv1.GET("/resources", v1.GetResources)
		apiv1.PUT("/resource/:id", v1.EditResource)
		apiv1.DELETE("/resource/:id", v1.DeleteResource)

		apiv1.GET("/resourcetype/:id", v1.GetResourceType)
		apiv1.POST("/resourcetype", v1.AddResourceType)
		apiv1.GET("/resourcetypes", v1.GetResourceTypes)
		apiv1.PUT("/resourcetype/:id", v1.EditResourceType)
		apiv1.DELETE("/resourcetype/:id", v1.DeleteResourceType)

		apiv1.GET("/macro/:id", v1.GetMacro)
		apiv1.POST("/macro", v1.AddMacro)
		apiv1.POST("/macro/addorupdate", v1.AddOrUpdateMacro)
		apiv1.GET("/macros", v1.GetMacros)
		apiv1.PUT("/macro/:id", v1.EditMacro)
		apiv1.DELETE("/macro/:id", v1.DeleteMacro)

		apiv1.GET("/host/deploy/:id", v1.GetHostDeploy)
		apiv1.POST("/host/deploy", v1.AddHostDeploy)
		apiv1.GET("/host/deploys", v1.GetHostDeploys)
		apiv1.PUT("/host/deploy/:id", v1.EditHostDeploy)
		apiv1.DELETE("/host/deploy/:id", v1.DeleteHostDeploy)

	}
	return r
}
