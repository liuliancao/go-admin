package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model

	Name        string `json:"name" gorm:"unique;not null;comment:'标签信息,用户向的分类'"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

type TagHost struct {
	gorm.Model

	TID       uint   `json:"t_id" gorm:"comment:'主机标签'"`
	HID       uint   `json:"h_id"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type TagApp struct {
	gorm.Model

	TID       uint   `json:"t_id" gorm:"comment:'应用标签id'"`
	AID       uint   `json:"a_id" gorm:"comment:'应用id'"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type TagCluster struct {
	gorm.Model

	TID       uint   `json:"t_id" gorm:"comment:'集群标签标签id'"`
	CID       uint   `json:"c_id" gorm:"comment:'集群标签集群id'"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type TagUser struct {
	gorm.Model

	TID       uint   `json:"t_id" gorm:"comment:'用户标签标签id'"`
	UID       uint   `json:"u_id" gorm:"comment:'用户标签用户id'"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

func GetTag(id uint) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ? and deleted_at is NULL", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &tag, nil
}

func GetTags() ([]*Tag, error) {
	var tags []*Tag
	err := db.Where("deleted_at is NULL").Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func EditTag(id uint, data interface{}) error {
	var tag Tag
	if err := db.Model(&tag).Where("id = ? AND deleted_at is NULL", id).Update(data).Error; err != nil {
		return err
	}

	return nil
}

func AddTag(data map[string]interface{}) error {
	tag := Tag{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		CreatedBy:   data["created_by"].(string),
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTag(id uint) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllTag() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistTagByID(id uint) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}
