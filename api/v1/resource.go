package v1

import (
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	resourceservice "go-admin/service/resource"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetResourceType(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	resourceTypeService := resourceservice.ResourceType{ID: uint(id)}
	exists, err := resourceTypeService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_RESOURCETYPE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_RESOURCETYPE, nil)
		return
	}

	resourceType, err := resourceTypeService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESOURCETYPE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, resourceType)
}

func GetResourceTypes(c *gin.Context) {
	appG := app.Gin{C: c}

	resourceTypeService := resourceservice.ResourceType{}

	resourceTypes, err := resourceTypeService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESOURCETYPES_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["resourceTypes"] = resourceTypes

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddResourceTypeForm struct {
	Name         string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	Nickname     string `form:"nickname" valid:"Required;MaxSize(100)" json:"nickname" `
	Tag          string `form:"tag" json:"tag"`
	HTMLFormJSON string `form:"html_form_json" json:"html_form_json"`
	Description  string `form:"description" valid:"MaxSize(100)" json:"description"`
	CreatedBy    uint   `form:"created_by"  json:"created_by"`
}

func AddResourceType(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	rt := AddResourceTypeForm{}
	c.BindJSON(&rt)
	log.Printf("%v", &rt)

	uid := c.MustGet("jwtuid")

	ok, err := valid.Valid(&rt)
	if !ok {
		app.MarkErrors(valid.Errors)
		log.Println(valid.Errors, err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}

	resourceTypeService := resourceservice.ResourceType{Name: rt.Name, Nickname: rt.Nickname, Tag: rt.Tag, Description: rt.Description, HTMLFormJSON: rt.HTMLFormJSON, CreatedBy: uid.(uint)}

	id, err := resourceTypeService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_RESOURCETYPE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

type EditResourceTypeForm struct {
	Name         string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	Nickname     string `form:"nickname" valid:"Required;MaxSize(100)" json:"nickname" `
	Tag          string `form:"tag" json:"tag"`
	HTMLFormJSON string `form:"html_form_json" json:"html_form_json"`
	Description  string `form:"description" valid:"MaxSize(100)" json:"description"`
}

func EditResourceType(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	rt := EditResourceTypeForm{}
	c.BindJSON(&rt)
	log.Printf("%v", rt)
	uid := c.MustGet("jwtuid")

	ok, _ := valid.Valid(&rt)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	resourceTypeService := resourceservice.ResourceType{
		ID:           id,
		Name:         rt.Name,
		Nickname:     rt.Nickname,
		Description:  rt.Description,
		Tag:          rt.Tag,
		HTMLFormJSON: rt.HTMLFormJSON,
		UpdatedBy:    uid.(uint),
	}
	exists, err := resourceTypeService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_RESOURCETYPE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_RESOURCETYPE, nil)
		return
	}

	err = resourceTypeService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_RESOURCETYPE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteResourceType(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	resourceTypeService := resourceservice.ResourceType{ID: uint(id)}
	exists, err := resourceTypeService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_RESOURCETYPE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_RESOURCETYPE, nil)
		return
	}

	err = resourceTypeService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_RESOURCETYPE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetResource(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	resourceService := resourceservice.Resource{ID: uint(id)}
	exists, err := resourceService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_RESOURCE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_RESOURCE, nil)
		return
	}

	resource, err := resourceService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESOURCE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, resource)
}

func GetResources(c *gin.Context) {
	appG := app.Gin{C: c}

	resourceService := resourceservice.Resource{}

	resources, err := resourceService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESOURCES_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["resources"] = resources

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddResourceForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	TID         uint   `form:"t_id" json:"t_id"`
	DataJSON    string `form:"data_json" json:"data_json"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	CreatedBy   uint   `form:"created_by"  json:"created_by"`
}

func AddResource(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	r := AddResourceForm{}
	c.BindJSON(&r)
	log.Printf("%v", &r)

	uid := c.MustGet("jwtuid")

	ok, _ := valid.Valid(&r)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	resourceService := resourceservice.Resource{
		Name:        r.Name,
		TID:         r.TID,
		Description: r.Description,
		DataJSON:    r.DataJSON,
		CreatedBy:   uid.(uint),
	}

	id, err := resourceService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_RESOURCE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

type EditResourceForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	TID         uint   `form:"t_id" json:"t_id"`
	DataJSON    string `form:"data_json" json:"data_json"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
}

func EditResource(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	r := EditResourceForm{}
	c.BindJSON(&r)
	log.Printf("%v", r)
	uid := c.MustGet("jwtuid")

	ok, _ := valid.Valid(&r)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	resourceService := resourceservice.Resource{
		ID:          id,
		Name:        r.Name,
		TID:         r.TID,
		Description: r.Description,
		DataJSON:    r.DataJSON,
		UpdatedBy:   uid.(uint),
	}
	exists, err := resourceService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_RESOURCE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_RESOURCE, nil)
		return
	}

	err = resourceService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_RESOURCE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteResource(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	resourceService := resourceservice.Resource{ID: uint(id)}
	exists, err := resourceService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_RESOURCE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_RESOURCE, nil)
		return
	}

	err = resourceService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_RESOURCE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
