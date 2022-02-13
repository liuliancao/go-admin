package v1

import (
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	departmentservice "go-admin/service/department"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetDepartment(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	departmentService := departmentservice.Department{ID: uint(id)}
	exists, err := departmentService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_DEPARTMENT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_DEPARTMENT, nil)
		return
	}

	department, err := departmentService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DEPARTMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, department)
}

func GetDepartments(c *gin.Context) {
	appG := app.Gin{C: c}

	departmentService := departmentservice.Department{}

	departments, err := departmentService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DEPARTMENTS_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["departments"] = departments

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddDepartmentForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	Status      uint   `form:"status" valid:"Required;Range(0,1)" json:"status"`
	Parent      uint   `form:"parent" json:"parent"`
	CreatedBy   uint   `form:"created_by" valid:"Required;" json:"created_by"`
}

func AddDepartment(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	d := AddDepartmentForm{}
	c.BindJSON(&d)
	log.Printf("%v", &d)

	ok, _ := valid.Valid(&d)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	departmentService := departmentservice.Department{Name: d.Name, Description: d.Description, Status: d.Status, CreatedBy: d.CreatedBy}

	if err := departmentService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_DEPARTMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditDepartmentForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)"`
	Description string `form:"description" valid:"MaxSize(100)"`
	Parent      uint   `form:"parent" json:"parent"`
	Status      uint   `form:"status" valid:"Required;Range(0,1)"`
	UpdatedBy   uint   `form:"updated_by" valid:"Required;"`
}

func EditDepartment(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	p := EditDepartmentForm{}
	c.BindJSON(&p)
	log.Printf("%v", p)

	ok, _ := valid.Valid(&p)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	departmentService := departmentservice.Department{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		Parent:      p.Parent,
		Status:      p.Status,
		UpdatedBy:   p.UpdatedBy,
	}
	exists, err := departmentService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_DEPARTMENT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_DEPARTMENT, nil)
		return
	}

	err = departmentService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_DEPARTMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteDepartment(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	departmentService := departmentservice.Department{ID: uint(id)}
	exists, err := departmentService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_DEPARTMENT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_DEPARTMENT, nil)
		return
	}

	err = departmentService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_DEPARTMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
