package main

import (
	"fmt"
	"go-admin/models"
	"go-admin/pkg/logging"
	"go-admin/pkg/setting"
	"go-admin/routers"
	"log"
	"net/http"
	"time"
)

func init() {
	fmt.Printf("initing %s", "go-admin")
	setting.Setup()
	models.Setup()

}

// @title op运维平台
// @version 1.0
// @description 基于golang的运维后端
// @contact.name liuliancao
// @contact.email liuliancao@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost
func main() {
	//gin.SetMode()
	logging.InitLogger()
	routerInit := routers.InitRouter()
	endPoint := "0.0.0.0:5120"
	readTimeout := 20 * time.Second
	writeTimeout := 20 * time.Second
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routerInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
