package v1

import (
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	appservice "go-admin/service/app"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetApp(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors[0])
		return
	}

	appService := appservice.App{ID: uint(id)}
	exists, err := appService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_APP_FAIL, err)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_APP, nil)
		return
	}

	app, err := appService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_APP_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, app)
}

type GetAppsForm struct {
	PID    uint `json:"pid"`
	EID    uint `json:"eid"`
	Parent uint `json:"parent"`
	Status uint `json:"status"`
}

func GetApps(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	a := GetAppsForm{
		PID:    uint(com.StrTo(c.DefaultQuery("pid", "0")).MustInt()),
		EID:    uint(com.StrTo(c.DefaultQuery("eid", "0")).MustInt()),
		Parent: uint(com.StrTo(c.DefaultQuery("parent", "0")).MustInt()),
		Status: uint(com.StrTo(c.DefaultQuery("status", "0")).MustInt()),
	}
	c.ShouldBind(&a)
	ok, _ := valid.Valid(&a)
	if !ok {
		app.MarkErrors(valid.Errors)
		log.Println(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}

	appService := appservice.App{
		PID:    a.PID,
		EID:    a.EID,
		Parent: a.Parent,
		Status: a.Status,
	}

	apps, err := appService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_APPS_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["apps"] = apps

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddAppForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	PID         uint   `form:"pid" valid:"Required" json:"pid"`
	Parent      uint   `form:"parent" json:"parent"`
	EID         uint   `form:"eid" json:"eid"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	Status      uint   `form:"status" valid:"Required;Range(0,1)" json:"status"`
	CreatedBy   uint   `form:"created_by" valid:"Required" json:"created_by"`
}

func AddApp(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	a := AddAppForm{}
	c.BindJSON(&a)
	log.Printf("%v", &a)

	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		log.Println(valid.Errors)
	}

	appService := appservice.App{Name: a.Name, Description: a.Description, Status: a.Status, PID: a.PID, CreatedBy: a.CreatedBy, EID: a.EID}

	err := appService.Add()
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_APP_FAIL, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditAppForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)"`
	Description string `form:"description" valid:"MaxSize(100)"`
	PID         uint   `form:"pid" valid:"Required" json:"pid"`
	EID         uint   `form:"eid" json:"eid"`
	Parent      uint   `form:"parent" json"parent"`
	Status      uint   `form:"status" valid:"Required;Range(0,1)"`
	UpdatedBy   uint   `form:"updated_by" valid:"Required"`
}

func EditApp(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	a := EditAppForm{}
	c.BindJSON(&a)
	log.Printf("%v", a)

	ok, _ := valid.Valid(&a)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	appService := appservice.App{
		ID:          id,
		Name:        a.Name,
		Description: a.Description,
		Parent:      a.Parent,
		Status:      a.Status,
		UpdatedBy:   a.UpdatedBy,
		PID:         a.PID,
	}
	exists, err := appService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_APP_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_APP, nil)
		return
	}

	err = appService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_APP_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteApp(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	appService := appservice.App{ID: uint(id)}
	exists, err := appService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_APP_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_APP, nil)
		return
	}

	err = appService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_APP_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
func GetAppEnv(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	appenvService := appservice.AppEnv{ID: uint(id)}
	exists, err := appenvService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_APPENV_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_APPENV, nil)
		return
	}

	appenv, err := appenvService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_APPENV_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, appenv)
}

func GetAppEnvs(c *gin.Context) {
	appG := app.Gin{C: c}

	appenvService := appservice.AppEnv{}

	appenvs, err := appenvService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_APPENVS_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["appenvs"] = appenvs

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddAppEnvForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	CreatedBy   uint   `form:"created_by" valid:"Required;" json:"created_by"`
}

func AddAppEnv(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	cl := AddAppEnvForm{}
	c.BindJSON(&cl)
	log.Printf("%v", &cl)

	ok, _ := valid.Valid(&cl)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	appenvService := appservice.AppEnv{Name: cl.Name, Description: cl.Description, CreatedBy: cl.CreatedBy}

	if err := appenvService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_APPENV_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditAppEnvForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)"`
	Description string `form:"description" valid:"MaxSize(100)"`
	UpdatedBy   uint   `form:"updated_by" valid:"Required;"`
}

func EditAppEnv(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	cl := EditAppEnvForm{}
	c.BindJSON(&cl)
	log.Printf("%v", cl)

	ok, _ := valid.Valid(&cl)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	appenvService := appservice.AppEnv{
		ID:          id,
		Name:        cl.Name,
		Description: cl.Description,
		UpdatedBy:   cl.UpdatedBy,
	}
	exists, err := appenvService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_APPENV_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_APPENV, nil)
		return
	}

	err = appenvService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_APPENV_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteAppEnv(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	appenvService := appservice.AppEnv{ID: uint(id)}
	exists, err := appenvService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_APPENV_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_APPENV, nil)
		return
	}

	err = appenvService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_APPENV_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type AddAppClusterForm struct {
	AID       uint `form:"a_id" valid:"required" json:"a_id"`
	CID       uint `form:"c_id" valid:"required" json:"c_id"`
	CreatedBy uint `form:"created_by" valid:"Required" json:"created_by"`
}

func AddAppCluster(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	a := AddAppClusterForm{}
	c.BindJSON(&a)
	log.Printf("%v", &a)

	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		log.Println(valid.Errors)
	}

	uid := c.MustGet("jwtuid")

	appClusterService := appservice.AppCluster{AID: a.AID, CID: a.CID, CreatedBy: uid.(uint)}

	err := appClusterService.Add()
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_APP_CLUSTER_FAIL, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type AddAppUserForm struct {
	AID uint `form:"a_id" valid:"required" json:"a_id"`
	UID uint `form:"u_id" valid:"required" json:"u_id"`
}

func AddAppUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	au := AddAppUserForm{}
	c.BindJSON(&au)
	log.Printf("%v", &au)

	ok, _ := valid.Valid(&au)

	if !ok {
		app.MarkErrors(valid.Errors)
		log.Println(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
	}

	uid := c.MustGet("jwtuid")

	appUserService := appservice.AppUser{AID: au.AID, UID: au.UID, CreatedBy: uid.(uint)}

	err := appUserService.Add()
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_APP_USER_FAIL, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type GetAppUsersForm struct {
	AID uint `json:"a_id"`
}

func GetAppUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	au := GetAppUsersForm{
		AID: uint(com.StrTo(c.DefaultQuery("a_id", "0")).MustInt()),
	}
	c.ShouldBind(&au)
	ok, _ := valid.Valid(&au)
	if !ok {
		app.MarkErrors(valid.Errors)
		log.Println(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}
	appUserService := appservice.AppUser{}
	users, err := appUserService.Get(au.AID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_APPUSERS_FAIL, err)
		return
	}

	data := make(map[string]interface{})
	data["users"] = users

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
