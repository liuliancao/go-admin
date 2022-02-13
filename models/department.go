package models

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model

	Name        string `json:"name" gorm:"unique;not null;comment:'部门表'"`
	Description string `json:"description"`
	Parent      uint   `json:"parent" gorm:"comment:'上级部门id'"`
	Status      uint   `json:"status"`
	CreatedBy   uint   `json:"created_by"`
	UpdateBy    uint   `json:"updated_by"`
}

func AddDepartment(data map[string]interface{}) error {
	department := Department{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		Parent:      data["parent"].(uint),
		Status:      data["status"].(uint),
		CreatedBy:   data["created_by"].(uint),
	}
	if err := db.Create(&department).Error; err != nil {
		return err
	}

	return nil
}
func GetDepartment(id uint) (*Department, error) {
	var department Department
	err := db.Where("id = ? and deleted_at is NULL", id).First(&department).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &department, nil
}

func GetDepartments() ([]*Department, error) {
	var departments []*Department
	err := db.Where("deleted_at is NULL").Find(&departments).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return departments, nil
}

func EditDepartment(id uint, data interface{}) error {
	var department Department
	if err := db.Model(&department).Where("id = ? AND deleted_at is NULL", id).Update(data).Error; err != nil {
		return err
	}

	return nil
}

func DeleteDepartment(id uint) error {
	if err := db.Where("id = ? and deleted_at is NULL", id).Delete(&Department{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllDepartment() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&Department{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistDepartmentByID(id uint) (bool, error) {
	var department Department
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&department).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if department.ID > 0 {
		return true, nil
	}

	return false, nil
}
