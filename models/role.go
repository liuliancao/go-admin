package models

import (
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model

	Name        string `json:"name" gorm:"unique;not null;comment:'角色名称'"`
	Description string `json:"description"`
	Permissions string `json:"permissions" gorm:"comment:'角色权限列表'"`
	Status      uint   `json:"status"`
	CreatedBy   uint   `json:"created_by"`
	UpdatedBy   uint   `json:"updated_by"`
}
type RoleUser struct {
	gorm.Model

	RID       uint `json:"r_id" gorm:"comment:'用户关联角色的角色id'"`
	UID       uint `json:"u_id" gorm:"comment:'用户关联角色的用户id'"`
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

func GetRole(id uint) (*Role, error) {
	var role Role
	err := db.Where("id = ? and deleted_at is NULL", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &role, nil
}

func GetRoles() ([]*Role, error) {
	var roles []*Role
	err := db.Where("deleted_at is NULL").Find(&roles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return roles, nil
}

func EditRole(id uint, data interface{}) error {
	var role Role
	if err := db.Model(&role).Where("id = ? AND deleted_at is NULL", id).Update(data).Error; err != nil {
		return err
	}

	return nil
}

func AddRole(data map[string]interface{}) error {
	role := Role{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		Status:      data["status"].(uint),
		CreatedBy:   data["created_by"].(uint),
	}
	if err := db.Create(&role).Error; err != nil {
		return err
	}

	return nil
}

//AddRoleUser for add role and user relative table
//params
//uuid: created user
func AddRoleUser(rid, uid, createdBy uint) error {
	roleUser := RoleUser{
		RID:       rid,
		UID:       uid,
		CreatedBy: createdBy,
	}
	if err := db.Create(&roleUser).Error; err != nil {
		return err
	}

	return nil
}

func DeleteRoleUser(rid, uid, createdBy uint) error {
	exists, err := ExistUserByID(createdBy)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("not exists user " + strconv.Itoa(int(createdBy)))
	}
	if err := db.Where("uid = ? AND rid = ? AND deleted_at is NULL", createdBy, rid).Delete(&RoleUser{}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteRole(id uint) error {
	if err := db.Where("id = ? and deleted_at is NULL", id).Delete(&Role{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllRole() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&Role{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistRoleByID(id uint) (bool, error) {
	var role Role
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if role.ID > 0 {
		return true, nil
	}

	return false, nil
}
