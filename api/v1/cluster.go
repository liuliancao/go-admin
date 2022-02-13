package v1

import (
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	clusterservice "go-admin/service/cluster"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetCluster(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	clusterService := clusterservice.Cluster{ID: uint(id)}
	exists, err := clusterService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CLUSTER_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CLUSTER, nil)
		return
	}

	cluster, err := clusterService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CLUSTER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, cluster)
}

type GetClustersForm struct {
	AID uint `json:"a_id"`
}

func GetClusters(c *gin.Context) {
	appG := app.Gin{C: c}

	cl := GetClustersForm{
		AID: uint(com.StrTo(c.DefaultQuery("a_id", "0")).MustInt()),
	}
	c.ShouldBind(&cl)
	valid := validation.Validation{}
	ok, _ := valid.Valid(&cl)
	if !ok {
		app.MarkErrors(valid.Errors)
		log.Println(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}
	clusterService := clusterservice.Cluster{}

	clusters, err := clusterService.GetAll(cl.AID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CLUSTERS_FAIL, nil)
	}

	data := make(map[string]interface{})
	data["clusters"] = clusters

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddClusterForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)" json:"name"`
	Description string `form:"description" valid:"MaxSize(100)" json:"description"`
	Status      uint   `form:"status" valid:"Required;Range(0,1)" json:"status"`
	CreatedBy   uint   `form:"created_by"  json:"created_by"`
}

func AddCluster(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	cl := AddClusterForm{}
	c.BindJSON(&cl)
	log.Printf("%v", &cl)

	uid := c.MustGet("jwtuid")

	ok, _ := valid.Valid(&cl)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	clusterService := clusterservice.Cluster{Name: cl.Name, Description: cl.Description, Status: cl.Status, CreatedBy: uid.(uint)}

	id, err := clusterService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CLUSTER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

type EditClusterForm struct {
	Name        string `form:"name" valid:"Required;MaxSize(100)"`
	Description string `form:"description" valid:"MaxSize(100)"`
	Status      uint   `form:"status" valid:"Required;Range(0,1)"`
	UpdatedBy   uint   `form:"updated_by" valid:"Required;"`
}

func EditCluster(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	cl := EditClusterForm{}
	c.BindJSON(&cl)
	log.Printf("%v", cl)

	ok, _ := valid.Valid(&cl)
	if !ok {
		app.MarkErrors(valid.Errors)
	}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	clusterService := clusterservice.Cluster{
		ID:          id,
		Name:        cl.Name,
		Description: cl.Description,
		Status:      cl.Status,
		UpdatedBy:   cl.UpdatedBy,
	}
	exists, err := clusterService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CLUSTER_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CLUSTER, nil)
		return
	}

	err = clusterService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_CLUSTER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteCluster(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	clusterService := clusterservice.Cluster{ID: uint(id)}
	exists, err := clusterService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CLUSTER_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CLUSTER, nil)
		return
	}

	err = clusterService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_CLUSTER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type AddClusterHostForm struct {
	CID uint `form:"c_id" valid:"Required" json:"c_id"`
	HID uint `form:"h_id" valid:"Required" json:"h_id"`
}

func AddClusterHost(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	clh := AddClusterHostForm{}
	c.ShouldBind(&clh)
	log.Printf("%v", &clh)

	ok, _ := valid.Valid(&clh)
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}

	uid := c.MustGet("jwtuid")
	clusterService := clusterservice.ClusterHost{CID: clh.CID, HID: clh.HID, CreatedBy: uid.(uint)}

	id, err := clusterService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CLUSTER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

type AddClusterUserForm struct {
	CID uint `form:"c_id" valid:"Required" json:"c_id"`
	UID uint `form:"u_id" valid:"Required" json:"u_id"`
}

func AddClusterUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	cu := AddClusterUserForm{}
	c.ShouldBind(&cu)
	log.Printf("%v", &cu)

	ok, _ := valid.Valid(&cu)
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}

	uid := c.MustGet("jwtuid")

	clusterUserService := clusterservice.ClusterUser{CID: cu.CID, UID: cu.UID, CreatedBy: uid.(uint)}

	if _, err := clusterUserService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CLUSTERUSER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type GetClusterUsersForm struct {
	CID uint `form:"c_id" json:"c_id"`
}

func GetClusterUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	cu := GetClusterUsersForm{
		CID: uint(com.StrTo(c.DefaultQuery("c_id", "0")).MustInt()),
	}
	c.ShouldBind(&cu)
	ok, _ := valid.Valid(&cu)
	if !ok {
		app.MarkErrors(valid.Errors)
		log.Println(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}
	clusterUserService := clusterservice.ClusterUser{}
	users, err := clusterUserService.Get(cu.CID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CLUSTERUSER_FAIL, err)
		return
	}

	data := make(map[string]interface{})
	data["users"] = users

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
