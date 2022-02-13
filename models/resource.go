package models

import "github.com/jinzhu/gorm"

type ResourceType struct {
	gorm.Model

	Name         string `json:"name" gorm:"unique;not null"`
	Nickname     string `json:"nickname" gorm"unique"`
	Tag          string `json:"tag"`
	Description  string `json:"description"`
	HTMLFormJSON string `json:"html_form_json" gorm:"size:65535"`

	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

type Resource struct {
	gorm.Model

	Name        string `json:"name" gorm:"unique;not null"`
	TID         uint   `json:"t_id"`
	Description string `json:"description"`
	DataJSON    string `json:"data_json"`

	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

func GetResourceType(id uint) (*ResourceType, error) {
	var resourceType ResourceType
	err := db.Where("id = ? and deleted_at is NULL", id).First(&resourceType).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &resourceType, nil
}

func GetResourceTypes() ([]*ResourceType, error) {
	var resourceTypes []*ResourceType
	err := db.Where("deleted_at is NULL").Find(&resourceTypes).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return resourceTypes, nil
}

func EditResourceType(id uint, data interface{}) error {
	var resourceType ResourceType
	if err := db.Model(&resourceType).Where("id = ? AND deleted_at is ?", id, "NULL").Update(data).Error; err != nil {
		return err
	}

	return nil
}

func AddResourceType(data map[string]interface{}) (id uint, err error) {
	resourceType := ResourceType{
		Name:         data["name"].(string),
		Nickname:     data["nickname"].(string),
		Tag:          data["tag"].(string),
		Description:  data["description"].(string),
		HTMLFormJSON: data["html_form_json"].(string),
		CreatedBy:    data["created_by"].(uint),
	}
	if err := db.Create(&resourceType).Error; err != nil {
		return 0, err
	}

	return resourceType.ID, nil
}
func ExistResourceTypeByID(id uint) (bool, error) {
	var resourceType ResourceType
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&resourceType).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if resourceType.ID > 0 {
		return true, nil
	}

	return false, nil
}
func DeleteResourceType(id uint) error {
	if err := db.Where("id = ?", id).Delete(&ResourceType{}).Error; err != nil {
		return err
	}

	return nil
}

func GetResource(id uint) (*Resource, error) {
	var resource Resource
	err := db.Where("id = ? and deleted_at is NULL", id).First(&resource).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &resource, nil
}

func GetResources() ([]*Resource, error) {
	var resources []*Resource
	err := db.Where("deleted_at is NULL").Find(&resources).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return resources, nil
}

func EditResource(id uint, data interface{}) error {
	var resource Resource
	if err := db.Model(&resource).Where("id = ? AND deleted_at is ?", id, "NULL").Update(data).Error; err != nil {
		return err
	}

	return nil
}

func AddResource(data map[string]interface{}) (id uint, err error) {
	resource := Resource{
		Name:        data["name"].(string),
		TID:         data["t_id"].(uint),
		Description: data["description"].(string),
		DataJSON:    data["data_json"].(string),
		CreatedBy:   data["created_by"].(uint),
	}
	if err := db.Create(&resource).Error; err != nil {
		return 0, err
	}

	return resource.ID, nil
}
func ExistResourceByID(id uint) (bool, error) {
	var resource Resource
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&resource).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if resource.ID > 0 {
		return true, nil
	}

	return false, nil
}
func DeleteResource(id uint) error {
	if err := db.Where("id = ?", id).Delete(&Resource{}).Error; err != nil {
		return err
	}

	return nil
}
