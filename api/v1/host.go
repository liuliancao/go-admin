package v1

import (
	"fmt"
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	hostservice "go-admin/service/host"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetHost(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	hostService := hostservice.Host{ID: uint(id)}
	exists, err := hostService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_HOST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_HOST, nil)
		return
	}

	host, err := hostService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_HOST_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, host)
}

type GetHostsForm struct {
	Hostname  string `json:"hostname" valid:"MaxSize(100)"`
	OS        string `json:"os" valid:"MaxSize(20)"`
	CID       uint   `json:"cid"`
	AID       uint   `json:"aid"`
	GID       uint   `json:"gid"`
	Cores     uint   `json:"cores"`
	IP        string `json:"ip" valid:"MaxSize(32)"`
	Status    uint   `json:"status"`
	CreatedBy uint   `json:"created_by"`
	UpdatedBy uint   `json:"updated_by"`
	PageNum   uint   `json:"pagenum"`
	PageSize  uint   `json:"pagesize"`
	Key       string `json:"key"`
}

func GetHosts(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	h := GetHostsForm{
		Hostname:  c.DefaultQuery("hostname", "null"),
		OS:        c.DefaultQuery("os", "null"),
		IP:        c.DefaultQuery("ip", "null"),
		CID:       uint(com.StrTo(c.DefaultQuery("cid", "0")).MustInt()),
		AID:       uint(com.StrTo(c.DefaultQuery("aid", "0")).MustInt()),
		GID:       uint(com.StrTo(c.DefaultQuery("gid", "0")).MustInt()),
		Cores:     uint(com.StrTo(c.DefaultQuery("cores", "0")).MustInt()),
		Status:    uint(com.StrTo(c.DefaultQuery("status", "9999")).MustInt()),
		CreatedBy: uint(com.StrTo(c.DefaultQuery("created_by", "0")).MustInt()),
		UpdatedBy: uint(com.StrTo(c.DefaultQuery("updated_by", "0")).MustInt()),
		PageNum:   uint(com.StrTo(c.DefaultQuery("pagenum", "0")).MustInt()),
		PageSize:  uint(com.StrTo(c.DefaultQuery("pagesize", "0")).MustInt()),
		Key:       c.DefaultQuery("key", "null"),
	}
	c.ShouldBind(&h)
	ok, _ := valid.Valid(&h)
	if !ok {
		app.MarkErrors(valid.Errors)
		log.Println(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}
	hostService := hostservice.Host{
		Hostname:  h.Hostname,
		OS:        h.OS,
		Cores:     h.Cores,
		CID:       h.CID,
		AID:       h.AID,
		GID:       h.GID,
		IP:        h.IP,
		Status:    h.Status,
		CreatedBy: h.CreatedBy,
		UpdatedBy: h.UpdatedBy,
		PageSize:  h.PageSize,
		PageNum:   h.PageNum,
		Key:       h.Key,
	}
	/*total, err := hostService.GetTotal()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TOTAL_FAIL, err)
	}*/

	total, hosts, err := hostService.GetVerboseHostsWithPage()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_HOSTS_FAIL, err)
		return
	}

	data := make(map[string]interface{})
	data["hosts"] = hosts
	data["total"] = total
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetGuarders(c *gin.Context) {
	appG := app.Gin{C: c}
	guarderService := hostservice.Guarder{}
	guarders, err := guarderService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INTERNAL_ERROR, err)
	}

	data := make(map[string]interface{})
	data["guarders"] = guarders

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddHostForm struct {
	Hostname    string `json:"hostname" valid:"Required;MaxSize(255)"`
	OS          string `json:"os" valid:"Required;MaxSize(20)"`
	Cores       uint   `json:"cores"`
	GID         uint   `json:"gid"`
	MemorySize  uint64 `json:"memory_size"`
	IP          string `json:"ip" valid:"Required;MaxSize(255)"`
	Status      uint   `json:"status" valid:"Required;"`
	Description string `json:"description" valid:"MaxSize(255)"`
	Extras      string `json:"extras" valid:"MaxSize(3000)"`
	Uptime      uint64 `json:"uptime"`
	CreatedBy   uint   `json:"created_by"`
}

func AddHost(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	h := AddHostForm{}
	c.BindJSON(&h)
	log.Printf("%v", &h)

	uid := c.MustGet("jwtuid")

	ok, _ := valid.Valid(&h)
	if !ok {
		log.Println(valid.Errors)
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}

	hostService := hostservice.Host{
		Hostname:    h.Hostname,
		OS:          h.OS,
		Cores:       h.Cores,
		GID:         h.GID,
		MemorySize:  h.MemorySize,
		IP:          h.IP,
		Status:      h.Status,
		Description: h.Description,
		Extras:      h.Extras,
		Uptime:      h.Uptime,
		CreatedBy:   uid.(uint),
	}

	id, err := hostService.Add()
	if err != nil || id == 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_HOST_FAIL, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

// ID is different.
type EditHostForm struct {
	Hostname    string `json:"hostname" valid:"MaxSize(100)" `
	OS          string `json:"os" valid:"MaxSize(20)" `
	Cores       uint   `json:"cores"`
	MemorySize  uint64 `json:"memory_size" `
	IP          string `json:"ip" valid:"MaxSize(16)" `
	GID         uint   `json:"gid"`
	Status      uint   `json:"status"`
	Description string `json:"description" valid:"MaxSize(100)" `
	Extras      string `json:"extras" valid:"MaxSize(3000)" `
	Uptime      uint64 `json:"uptime" `
	UpdatedBy   uint   `json:"updated_by" valid:"MaxSize(50)" `
}

func EditHost(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	h := EditHostForm{}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	c.BindJSON(&h)

	log.Printf("the update raw data is: %v", h)

	ok, _ := valid.Valid(&h)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	hostService := hostservice.Host{
		ID:          id,
		Hostname:    h.Hostname,
		OS:          h.OS,
		Cores:       h.Cores,
		MemorySize:  h.MemorySize,
		IP:          h.IP,
		Status:      h.Status,
		GID:         h.GID,
		Description: h.Description,
		Extras:      h.Extras,
		Uptime:      h.Uptime,
		UpdatedBy:   h.UpdatedBy,
	}

	exists, err := hostService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_HOST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_HOST, nil)
		return
	}

	err = hostService.Edit(h.IP)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_HOST_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteHost(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	hostService := hostservice.Host{ID: uint(id)}
	exists, err := hostService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_HOST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_HOST, nil)
		return
	}

	err = hostService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_HOST_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type MustAddHostForm struct {
	Hostname   string `json:"hostname" valid:"Required;MaxSize(100)"`
	OS         string `json:"os" valid:"Required;MaxSize(20)"`
	Cores      uint   `json:"cores" valid:"Range(1,1000)"`
	MemorySize uint64 `json:"memory_size"`

	IP        string `json:"ip" valid:"Required;MaxSize(32)"`
	Status    uint   `json:"status" valid:"Required;Range(0,1)"`
	Extras    string `json:"extras" valid:"Required;MaxSize(3000)"`
	Uptime    uint64 `json:"uptime"`
	UpdatedBy uint   `json:"updated_by" valid:"Required;"`
	CreatedBy uint   `json:"created_by" valid:"Required;"`
}

type AddGuarderForm struct {
	Datacenter string `json:"datacenter" valid:"Required"`
	Provider   string `json:"provider" valid:"Required"`
	Port       uint   `json:"port" valid:"Required"`
	IP         string `json:"ip" valid:"Required"`
	Extras     string `json:"extras" valid:"Required"`
}

func AddGuarder(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	g := AddGuarderForm{}
	c.BindJSON(&g)

	fmt.Println("g is ", g)
	ok, _ := valid.Valid(&g)
	if !ok {
		log.Println(valid.Errors)
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}

	guarderService := hostservice.Guarder{
		Datacenter: g.Datacenter,
		Provider:   g.Provider,
		Port:       g.Port,
		IP:         g.IP,
		Extras:     g.Extras,
	}

	id, err := guarderService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_GUARDER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

type EditGuarderForm struct {
	Datacenter string `json:"datacenter" valid:"MaxSize(100)" `
	IP         string `json:"ip"`
	Port       uint   `json:"port"`
	Provider   string `json:"provider"`
	Extras     string `json:"extras" valid:"MaxSize(3000)" `
}

func EditGuarder(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	g := EditGuarderForm{}
	id := uint(com.StrTo(c.Param("id")).MustInt())

	c.BindJSON(&g)

	log.Printf("the update raw data is: %v", g)

	ok, _ := valid.Valid(&g)
	if !ok {
		app.MarkErrors(valid.Errors)
	}

	guarderService := hostservice.Guarder{
		ID:         id,
		Datacenter: g.Datacenter,
		IP:         g.IP,
		Port:       g.Port,
		Provider:   g.Provider,
		Extras:     g.Extras,
	}

	exists, err := guarderService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_GUARDER_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_GUARDER, nil)
		return
	}

	err = guarderService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_GUARDER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteGuarder(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	guarderService := hostservice.Guarder{ID: uint(id)}
	exists, err := guarderService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_GUARDER_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_GUARDER, nil)
		return
	}

	err = guarderService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_GUARDER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
