package v1

import (
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	roleservice "go-admin/service/role"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetRole(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	roleService := roleservice.Role{ID: uint(id)}
	exists, err := roleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ROLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ROLE, nil)
		return
	}

	role, err := roleService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, role)
}

func GetRoles(c *gin.Context) {
	appG := app.Gin{C: c}

	roleService := roleservice.Role{}

	roles, err := roleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROLES_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["roles"] = roles

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddRoleForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	Status      uint   `form:"status" valid:"Required;Range(0,1)" json:"status"`
	CreatedBy   uint   `form:"created_by" valid:"Required;" json:"created_by"`
}

func AddRole(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	r := AddRoleForm{}
	c.BindJSON(&r)
	log.Printf("%v", &r)

	ok, _ := valid.Valid(&r)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	roleService := roleservice.Role{Name: r.Name, Description: r.Description, Status: r.Status, CreatedBy: r.CreatedBy}

	if err := roleService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ROLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditRoleForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)"`
	Description string `form:"description" valid:"MaxSize(100)"`
	Status      uint   `form:"status" valid:"Required;Range(0,1)"`
	UpdatedBy   uint   `form:"updated_by" valid:"Required"`
}

func EditRole(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	r := EditRoleForm{}
	c.BindJSON(&r)
	log.Printf("%v", r)

	ok, _ := valid.Valid(&r)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	roleService := roleservice.Role{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
		Status:      r.Status,
		UpdatedBy:   r.UpdatedBy,
	}
	exists, err := roleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ROLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ROLE, nil)
		return
	}

	err = roleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ROLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteRole(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	roleService := roleservice.Role{ID: uint(id)}
	exists, err := roleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ROLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ROLE, nil)
		return
	}

	err = roleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ROLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
