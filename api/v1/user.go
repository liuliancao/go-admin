package v1

import (
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	"log"
	"net/http"

	"go-admin/pkg/util"
	userservice "go-admin/service/user"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type Auth struct {
	Username string `valid: "Required; MaxSize(50)" json:"username"`
	Password string `valid: "Required; Maxsize(50)" json:"password"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	a := Auth{}
	c.BindJSON(&a)
	log.Printf("%v", &a)

	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
	}

	userService := userservice.Auth{Username: a.Username, Password: a.Password}
	isExist, err := userService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(a.Username, a.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	userService := userservice.User{ID: uint(id)}
	exists, err := userService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	user, err := userService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, user)
}

func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	userService := userservice.User{}

	users, err := userService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USERS_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["users"] = users

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddUserForm struct {
	DID       uint   `form:"d_id" valid:"Required;" json:"d_id"`
	Username  string `form:"username" valid:"Required;MaxSize(100)" json:"username"`
	Nickname  string `form:"nickname" json"nickname"`
	Password  string `form:"password" valid:"Required;MaxSize(100)" json:"password"`
	Gender    string `form:"gender" valid:"MaxSize(100)" json:"gender"`
	Phone     string `form:"phone" valid:"MaxSize(100)" json:"phone"`
	Mail      string `form:"mail" valid:"MaxSize(100)" json:"mail"`
	Token     string `form:"token" valid:"MaxSize(100)" json:"token"`
	Status    uint   `form:"status"  json:"status"`
	CreatedBy uint   `form:"created_by"  json:"created_by"`
}

func AddUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	u := AddUserForm{}
	c.BindJSON(&u)
	log.Printf("%v", &u)

	ok, _ := valid.Valid(&u)
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors[0])
		return
	}

	uid := c.MustGet("jwtuid")

	token := util.EncodeMD5(u.Username)
	userService := userservice.User{
		DID:       u.DID,
		Username:  u.Username,
		Nickname:  u.Nickname,
		Password:  u.Password,
		Gender:    u.Gender,
		Phone:     u.Phone,
		Mail:      u.Mail,
		Token:     token,
		Status:    u.Status,
		CreatedBy: uid.(uint),
	}

	if err := userService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditUserForm struct {
	DID       uint   `form:"d_id" valid:"Required" json:"d_id"`
	Username  string `form:"username" valid:"Required;MaxSize(100)" json:"username"`
	Nickname  string `form:"nickname" json:"nickname"`
	Password  string `form:"password" valid:"Required;MaxSize(100)" json:"password"`
	Gender    string `form:"gender" valid:"MaxSize(100)" json:"gender"`
	Phone     string `form:"phone" valid:"MaxSize(100)" json:"phone"`
	Mail      string `form:"mail" valid:"MaxSize(100)" json:"mail"`
	Token     string `form:"token" valid:"MaxSize(100)" json:"token"`
	Status    uint   `form:"status" valid:"Required;Range(0,1)" json:"status"`
	UpdatedBy uint   `form:"updated_by" valid:"MaxSize(50)" json:"updated_by"`
}

func EditUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	u := EditUserForm{}
	c.BindJSON(&u)
	log.Printf("%v", u)

	ok, _ := valid.Valid(&u)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	userService := userservice.User{
		ID:        id,
		DID:       u.DID,
		Username:  u.Username,
		Nickname:  u.Nickname,
		Password:  u.Password,
		Gender:    u.Gender,
		Phone:     u.Phone,
		Mail:      u.Mail,
		Token:     u.Token,
		Status:    u.Status,
		UpdatedBy: u.UpdatedBy,
	}
	exists, err := userService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	err = userService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userService := userservice.User{ID: uint(id)}
	exists, err := userService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	err = userService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
