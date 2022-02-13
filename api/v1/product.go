package v1

import (
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	productservice "go-admin/service/product"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	productService := productservice.Product{ID: uint(id)}
	exists, err := productService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PRODUCT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}

	product, err := productService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_PRODUCT_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, product)
}

func GetProducts(c *gin.Context) {
	appG := app.Gin{C: c}

	productService := productservice.Product{}

	products, err := productService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_PRODUCTS_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["products"] = products

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddProductForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	Parent      uint   `form:"parent" json:"parent"`
	Status      uint   `form:"status" valid:"Required;Range(0,1)" json:"status"`
	CreatedBy   uint   `form:"created_by"  json:"created_by"`
}

func AddProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	p := AddProductForm{}
	c.BindJSON(&p)
	log.Printf("%v", &p)

	uid := c.MustGet("jwtuid")

	ok, _ := valid.Valid(&p)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	productService := productservice.Product{Name: p.Name, Description: p.Description, Status: p.Status, CreatedBy: uid.(uint)}

	id, err := productService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

type EditProductForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)"`
	Description string `form:"description" valid:"MaxSize(100)"`
	Parent      uint   `form:"parent" json:"parent"`
	Status      uint   `form:"status" valid:"Required;Range(0,1)"`
	UpdatedBy   uint   `form:"updated_by" valid:"Required;"`
}

func EditProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	p := EditProductForm{}
	c.BindJSON(&p)
	log.Printf("%v", p)

	ok, _ := valid.Valid(&p)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	productService := productservice.Product{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		Status:      p.Status,
		Parent:      p.Parent,
		UpdatedBy:   p.UpdatedBy,
	}
	exists, err := productService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PRODUCT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}

	err = productService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteProduct(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	productService := productservice.Product{ID: uint(id)}
	exists, err := productService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PRODUCT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_PRODUCT, nil)
		return
	}

	err = productService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_PRODUCT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type AddProductUserForm struct {
	PID uint `form:"pid" json:"p_id"`
	UID uint `form:"uid" json:"u_id"`
}

func AddProductUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	pu := AddProductUserForm{}
	c.BindJSON(&pu)
	log.Printf("%v", &pu)

	ok, _ := valid.Valid(&pu)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	uid := c.MustGet("jwtuid")

	productUserService := productservice.ProductUser{PID: pu.PID, UID: pu.UID, CreatedBy: uid.(uint)}

	if err := productUserService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_PRODUCTUSER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type GetProductUsersForm struct {
	PID uint `json:"p_id"`
}

func GetProductUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	pu := GetProductUsersForm{
		PID: uint(com.StrTo(c.DefaultQuery("p_id", "0")).MustInt()),
	}
	c.ShouldBind(&pu)
	ok, _ := valid.Valid(&pu)
	if !ok {
		app.MarkErrors(valid.Errors)
		log.Println(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}
	productUserService := productservice.ProductUser{}
	users, err := productUserService.Get(pu.PID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_PRODUCTUSERS_FAIL, err)
		return
	}

	data := make(map[string]interface{})
	data["users"] = users

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
