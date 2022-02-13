package v1

import (
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	tag_service "go-admin/service/tag"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTag(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tagService := tag_service.Tag{ID: uint(id)}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_TAG_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	tag, err := tagService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, tag)
}

func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}

	tagService := tag_service.Tag{}

	tags, err := tagService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["tags"] = tags

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddTagForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	CreatedBy   string `form:"created_by" valid:"Required;MaxSize(50)" json:"created_by"`
}

func AddTag(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	t := AddTagForm{}
	c.BindJSON(&t)
	log.Printf("%v", &t)

	ok, _ := valid.Valid(&t)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	tagService := tag_service.Tag{Name: t.Name, Description: t.Description, CreatedBy: t.CreatedBy}

	if err := tagService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditTagForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)"`
	Description string `form:"description" valid:"MaxSize(100)"`
	UpdatedBy   string `form:"updated_by" valid:"Required;MaxSize(50)"`
}

func EditTag(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	t := EditTagForm{}
	c.BindJSON(&t)
	log.Printf("%v", t)

	ok, _ := valid.Valid(&t)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	tagService := tag_service.Tag{
		ID:          id,
		Name:        t.Name,
		Description: t.Description,
		UpdatedBy:   t.UpdatedBy,
	}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_TAG_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = tagService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	tagService := tag_service.Tag{ID: uint(id)}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_TAG_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = tagService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
