package models

import (
	"errors"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

//App used for product to app, logic product, in game == service
type App struct {
	gorm.Model

	Name        string `json:"name" gorm:"comment:'app名称'"`
	PID         uint   `json:"pid" gorm:"comment:'产品关联id'"` //product id
	EID         uint   `json:"eid" gorm:"comment:'绑定环境id'"` //env id
	Parent      uint   `json:"parent" gorm:"comment:'标记几级app'"`
	Description string `json:"description" gorm:"comment:'描述信息'"`
	Status      uint   `json:"status" gorm:"comment:'app状态'"`
	CreatedBy   uint   `json:"created_by" gorm:"comment:'创建人'"`
	UpdatedBy   uint   `json:"updated_by" gorm:"comment:'更新人'"`
}

//AppEnv expands App with online, test, pretest and so on
type AppEnv struct {
	gorm.Model

	Name        string `json:"name" gorm:"comment:'app环境名, 比如预发测试线上'"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"created_by"`
	UpdatedBy   uint   `json:"updated_by"`
}

//AppUser Association with app and user for many to many
type AppUser struct {
	gorm.Model

	AID       uint `json:"a_id" gorm:"comment:'app对应id'"`
	UID       uint `json:"u_id" gorm:"comment:'用户对用id'"`
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

//AppCluster Association with app and cluster for many to many
type AppCluster struct {
	gorm.Model

	AID       uint `json:"a_id" gorm:"index" gorm:"comment:'app对应id"`
	CID       uint `json:"c_id" gorm:"index" gorm:"comment:'集群对应id'"`
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

func GetApp(id uint) (*App, error) {
	var app App
	err := db.Where("id = ? and deleted_at is NULL", id).First(&app).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &app, nil
}

func GetApps(data map[string]interface{}) ([]*App, error) {
	var apps []*App

	pid := data["pid"].(uint)
	eid := data["eid"].(uint)
	parent := data["parent"].(uint)
	status := data["status"].(uint)
	newdb := db.New()
	if pid != 0 {
		newdb = newdb.Where("p_id = ?", pid)
	}
	if eid != 0 {
		newdb = newdb.Where("e_id = ?", eid)
	}
	if parent != 0 {
		newdb = newdb.Where("parent = ?", parent)
	}
	if status != 0 {
		newdb = newdb.Where("status = ?", status)
	}
	err := newdb.Where("deleted_at is NULL").Find(&apps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return apps, nil
}

func EditApp(id uint, data interface{}) error {
	var app App
	if err := db.Model(&app).Where("id = ? AND deleted_at is ?", id, "NULL").Update(data).Error; err != nil {
		return err
	}

	return nil
}

func AddApp(data map[string]interface{}) error {
	var product Product
	pid := data["pid"].(uint)
	if err := db.Where("id = ? AND deleted_at is NULL", pid).First(&product).Error; err != nil {
		return err
	}
	if exists, _ := ExistAppEnvByID(data["eid"].(uint)); exists == false {
		return errors.New("app's env not exists " + strconv.Itoa(int(data["eid"].(uint))))
	}
	app := App{
		Name:        data["name"].(string),
		PID:         data["pid"].(uint),
		EID:         data["eid"].(uint),
		Parent:      data["parent"].(uint),
		Description: data["description"].(string),
		Status:      data["status"].(uint),
		CreatedBy:   data["created_by"].(uint),
	}
	if err := db.Create(&app).Error; err != nil {
		return err
	}

	return nil
}

// AddAppUser bind app with multi users
// Params
// createdBy: created user
func AddAppUser(aid, uid uint, createdBy uint) error {

	var user User
	if err := db.Where("id = ? AND deleted_at is NULL", uid).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err := db.Where("id = ? AND deleted_at is NULL", createdBy).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	var app App
	if err := db.Where("id = ? AND deleted_at is NULL", aid).First(&app).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	appUser := AppUser{
		AID:       aid,
		UID:       uid,
		CreatedBy: createdBy,
	}
	if err := db.Create(&appUser).Error; err != nil {
		log.Fatalf("create app %d with user %d failed: %v", aid, uid, err)
	}
	return nil
}

func GetAppUsers(aid uint) ([]*User, error) {
	var users []*User
	err := db.Table("user").Where("id = any(select uid from app_user where a_id = ?)", aid).Scan(&users).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}

func DeleteAppUser(aid, uid uint) error {
	if err := db.Where("aid = ? AND uid = ? AND deleted_at is NULL", aid, uid).Delete(&AppUser{}).Error; err != nil {
		return err
	}

	return nil
}

func AddAppCluster(aid, cid, createdBy uint) error {
	var cluster Cluster
	if err := db.Where("id = ? AND deleted_at is NULL", createdBy).First(&cluster).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	var app App
	if err := db.Where("id = ? AND deleted_at is NULL", aid).First(&app).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	appUser := AppCluster{
		AID:       aid,
		CID:       cid,
		CreatedBy: createdBy,
	}
	if err := db.Create(&appUser).Error; err != nil {
		log.Fatalf("create app %d with user %d failed: %v", aid, createdBy, err)
	}
	return nil
}
func DeleteAppCluster(aid, cid uint) error {
	if err := db.Where("aid = ? AND cid = ? AND deleted_at is NULL", aid, cid).Delete(&AppCluster{}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteApp(id uint) error {
	if err := db.Where("id = ?", id).Delete(&App{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllApp() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&App{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistAppByID(id uint) (bool, error) {
	var app App
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&app).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if app.ID > 0 {
		return true, nil
	}

	return false, nil
}
func GetAppEnv(id uint) (*AppEnv, error) {
	var appenv AppEnv
	err := db.Where("id = ? and deleted_at is NULL", id).First(&appenv).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &appenv, nil
}

func GetAppEnvs() ([]*AppEnv, error) {
	var appenvs []*AppEnv
	err := db.Where("deleted_at is NULL").Find(&appenvs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return appenvs, nil
}

func EditAppEnv(id uint, data interface{}) error {
	var appenv AppEnv
	if err := db.Model(&appenv).Where("id = ? AND deleted_at is NULL", id).Update(data).Error; err != nil {
		return err
	}

	return nil
}

func AddAppEnv(data map[string]interface{}) error {
	appenv := AppEnv{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		CreatedBy:   data["created_by"].(uint),
	}
	if err := db.Create(&appenv).Error; err != nil {
		return err
	}

	return nil
}

func DeleteAppEnv(id uint) error {
	if err := db.Where("id = ? and deleted_at is NULL", id).Delete(&AppEnv{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllAppEnv() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&AppEnv{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistAppEnvByID(id uint) (bool, error) {
	var appenv AppEnv
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&appenv).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if appenv.ID > 0 {
		return true, nil
	}

	return false, nil
}
