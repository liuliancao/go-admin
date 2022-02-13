package v1

import (
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	deployservice "go-admin/service/deploy"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetHostDeploy(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	deployService := deployservice.HostDeploy{ID: uint(id)}
	exists, err := deployService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_HOSTDEPLOY_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_HOSTDEPLOY, nil)
		return
	}

	hostdeploy, err := deployService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_HOSTDEPLOY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, hostdeploy)
}

type GetHostDeploysForm struct {
	AID       uint   `form:"a_id" json:"a_id"`
	HID       uint   `form:"h_id"  json:"h_id"`
	Status    string `form:"status" json:"status"`
	Version   string `form:"version"  json:"version"`
	Stage     string `form:"stage" json:"stage"`
	CreatedBy uint   `form:"created_by" json:"created_by"`
}

func GetHostDeploys(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	d := GetHostDeploysForm{
		HID:       uint(com.StrTo(c.DefaultQuery("h_id", "0")).MustInt()),
		AID:       uint(com.StrTo(c.DefaultQuery("a_id", "0")).MustInt()),
		Status:    c.DefaultQuery("status", ""),
		Version:   c.DefaultQuery("version", ""),
		Stage:     c.DefaultQuery("stage", ""),
		CreatedBy: uint(com.StrTo(c.DefaultQuery("created_by", "0")).MustInt()),
	}
	c.ShouldBindJSON(&d)
	log.Printf("%v", &d)

	ok, _ := valid.Valid(&d)
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	deployService := deployservice.HostDeploy{
		AID:       d.AID,
		HID:       d.HID,
		Status:    d.Status,
		Version:   d.Version,
		Stage:     d.Stage,
		CreatedBy: d.CreatedBy,
	}

	hostdeploys, err := deployService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_HOSTDEPLOYS_FAIL, err)
		return
	}

	data := make(map[string]interface{})
	data["hostdeploys"] = hostdeploys

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddHostDeployForm struct {
	AID       uint   `form:"a_id" valid:"Required;" json:"a_id"`
	HID       uint   `form:"h_id" valid:"Required;" json:"h_id"`
	Status    string `form:"status" json:"status"`
	Version   string `form:"version"  json:"version"`
	Stage     string `form:"stage" json:"stage"`
	Message   string `form:"message"  json:"message"`
	AppMacros string `form:"app_macros" json:"app_macros"`
	CreatedBy uint   `form:"created_by" valid:"Required;" json:"created_by"`
}

func AddHostDeploy(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	uid := c.MustGet("jwtuid")

	d := AddHostDeployForm{}
	c.BindJSON(&d)
	log.Printf("%v", &d)

	ok, _ := valid.Valid(&d)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	deployService := deployservice.HostDeploy{
		AID:       d.AID,
		HID:       d.HID,
		Status:    d.Status,
		Version:   d.Version,
		Stage:     d.Stage,
		Message:   d.Message,
		AppMacros: d.AppMacros,
		CreatedBy: uid.(uint),
	}

	if err := deployService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_HOSTDEPLOY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditHostDeployForm struct {
	AID       uint   `form:"aid" valid:"Required;" json:"aid"`
	HID       uint   `form:"hid" valid:"Required;" json:"hid"`
	Status    string `form:"status" json:"status"`
	Stage     string `form:"stage" json:"stage"`
	Version   string `form:"version"  json:"version"`
	Message   string `form:"message"  json:"message"`
	AppMacros string `form:"app_macros" json:"app_macros"`
	UpdatedBy uint   `form:"updated_by" valid:"Required;"`
}

func EditHostDeploy(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	uid := c.MustGet("jwtuid")

	h := EditHostDeployForm{}
	c.BindJSON(&h)
	log.Printf("%v", h)

	ok, _ := valid.Valid(&h)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	deployService := deployservice.HostDeploy{
		ID:        id,
		AID:       h.AID,
		HID:       h.HID,
		Status:    h.Status,
		Stage:     h.Stage,
		Version:   h.Version,
		Message:   h.Message,
		AppMacros: h.AppMacros,
		UpdatedBy: uid.(uint),
	}
	exists, err := deployService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_HOSTDEPLOY_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_HOSTDEPLOY, nil)
		return
	}

	err = deployService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_HOSTDEPLOY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteHostDeploy(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	deployservice := deployservice.HostDeploy{ID: uint(id)}
	exists, err := deployservice.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_HOSTDEPLOY_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_HOSTDEPLOY, nil)
		return
	}

	err = deployservice.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_HOSTDEPLOY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
