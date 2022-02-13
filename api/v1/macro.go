package v1

import (
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	macroservice "go-admin/service/macro"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetMacro(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	macroService := macroservice.Macro{ID: uint(id)}
	exists, err := macroService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MACRO_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_MACRO, nil)
		return
	}

	macro, err := macroService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_MACRO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, macro)
}

func GetMacros(c *gin.Context) {
	appG := app.Gin{C: c}

	namespace := c.DefaultQuery("namespace", "")
	mType := c.DefaultQuery("m_type", "")
	log.Println(namespace)
	log.Println(mType)
	macroService := macroservice.Macro{}

	macros, err := macroService.GetAll(namespace, mType)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_MACROS_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["macros"] = macros

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddMacroForm struct {
	Namespace   string `form:"namespace" valid:"Required;MaxSize(100)" json:"namespace"`
	Key         string `form:"key" valid:"Required;MaxSize(100)" json:"key"`
	MType       string `form:"m_type" valid:"Required;" json:"m_type"`
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	Value       string `form:"value" valid:"Required" json:"value"`
	Encrypt     uint   `form:"encrypt" valid:"Required" json:"encrypt"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	CreatedBy   uint   `form:"created_by" valid:"Required;" json:"created_by"`
}

func AddMacro(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	uid := c.MustGet("jwtuid")

	m := AddMacroForm{}
	c.BindJSON(&m)
	log.Printf("%v", &m)

	ok, _ := valid.Valid(&m)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	macroService := macroservice.Macro{Namespace: m.Namespace, Key: m.Key, Name: m.Name, MType: m.MType, Description: m.Description, Value: m.Value, CreatedBy: uid.(uint)}

	if err := macroService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_MACRO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditMacroForm struct {
	Namespace   string `form:"namespace" valid:"Required;MaxSize(100)" json:"namespace"`
	Key         string `form:"key" valid:"Required;MaxSize(100)" json:"key"`
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	MType       string `form:"m_type" valid:"Required;" json:"m_type"`
	Value       string `form:"value" valid:"Required" json:"value"`
	Encrypt     uint   `form:"encrypt" valid:"Required" json:"encrypt"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	UpdatedBy   uint   `form:"updated_by" valid:"Required;"`
}

func EditMacro(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	m := EditMacroForm{}
	c.BindJSON(&m)
	log.Printf("%v", m)

	uid := c.MustGet("jwtuid")

	ok, _ := valid.Valid(&m)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	macroService := macroservice.Macro{
		ID:          id,
		Namespace:   m.Namespace,
		Key:         m.Key,
		Name:        m.Name,
		MType:       m.MType,
		Encrypt:     m.Encrypt,
		Value:       m.Value,
		Description: m.Description,
		UpdatedBy:   uid.(uint),
	}
	exists, err := macroService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MACRO_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_MACRO, nil)
		return
	}

	err = macroService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_MACRO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type AddOrUpdateMacroForm struct {
	Namespace   string `form:"namespace" valid:"Required;MaxSize(100)" json:"namespace"`
	Key         string `form:"key" valid:"Required;MaxSize(100)" json:"key"`
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	MType       string `form:"m_type" valid:"Required;" json:"m_type"`
	Value       string `form:"value" valid:"Required" json:"value"`
	Encrypt     uint   `form:"encrypt" valid:"Required" json:"encrypt"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	UpdatedBy   uint   `form:"updated_by" valid:"Required;"`
}

func AddOrUpdateMacro(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	uid := c.MustGet("jwtuid")

	m := AddOrUpdateMacroForm{}
	c.BindJSON(&m)
	log.Printf("%v", m)

	ok, _ := valid.Valid(&m)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	macroService := macroservice.Macro{
		Namespace:   m.Namespace,
		Key:         m.Key,
		Name:        m.Name,
		MType:       m.MType,
		Encrypt:     m.Encrypt,
		Value:       m.Value,
		Description: m.Description,
		UpdatedBy:   uid.(uint),
	}
	mid, err := macroService.ExistByNamespaceKey()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MACRO_FAIL, nil)
		return
	}
	if mid == 0 {
		if err := macroService.Add(); err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_ADD_MACRO_FAIL, nil)
			return
		}
	} else {
		macroService.ID = mid
		err = macroService.Edit()
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_MACRO_FAIL, nil)
			return
		}
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteMacro(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	macroService := macroservice.Macro{ID: uint(id)}
	exists, err := macroService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MACRO_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_MACRO, nil)
		return
	}

	err = macroService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_MACRO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
