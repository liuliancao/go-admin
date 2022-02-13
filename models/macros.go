package models

import "github.com/jinzhu/gorm"

// Macro is used for go admin config center
type Macro struct {
	gorm.Model
	Namespace   string `json:"namespace" gorm:"comment: '宏的命名空间'"`
	Key         string `json:"key" gorm:"unique;not null;comment:'宏的key，不能重复'"`
	MType       string `json:"m_type" gorm:"default:'string';comment:'宏的数据类型，默认是字符串'"`
	Name        string `json:"name" gorm:"comment:'宏的显示名称'"`
	Value       string `json:"value" grom:"comment:'宏的值'"`
	Encrypt     uint   `json:"encrypt" gorm:"comment: '是否加密，并且前后端隐藏'"`
	Description string `json:"description" gorm:"comment:'宏的描述信息'"`
	CreatedBy   uint   `json:"created_by"`
	UpdatedBy   uint   `json:"updated_by"`
}

func AddMacro(data map[string]interface{}) error {
	macro := Macro{
		Namespace:   data["namespace"].(string),
		Key:         data["key"].(string),
		Name:        data["name"].(string),
		Encrypt:     data["encrypt"].(uint),
		Value:       data["value"].(string),
		Description: data["description"].(string),
		MType:       data["m_type"].(string),
		CreatedBy:   data["created_by"].(uint),
	}
	if err := db.Create(&macro).Error; err != nil {
		return err
	}

	return nil
}
func GetMacro(id uint) (*Macro, error) {
	var macro Macro
	err := db.Where("id = ? and deleted_at is NULL", id).First(&macro).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &macro, nil
}

func GetMacros(data map[string]interface{}) ([]*Macro, error) {
	var macros []*Macro
	mType := data["m_type"].(string)
	namespace := data["namespace"].(string)
	newdb := db.New()
	if mType != "" {
		newdb = newdb.Where("m_type = ?", mType)
	}
	if namespace != "" {
		newdb = newdb.Where("namespace = ?", namespace)
	}

	err := newdb.Where("deleted_at is NULL").Find(&macros).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return macros, nil
}

func EditMacro(id uint, data interface{}) error {
	var macro Macro
	if err := db.Model(&macro).Where("id = ? AND deleted_at is NULL", id).Update(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteMacro(id uint) error {
	if err := db.Where("id = ? and deleted_at is NULL", id).Delete(&Macro{}).Error; err != nil {
		return err
	}
	return nil
}

func CleanAllMacro() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&Macro{}).Error; err != nil {
		return err
	}
	return nil
}

func ExistMacroByID(id uint) (bool, error) {
	var macro Macro
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&macro).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if macro.ID > 0 {
		return true, nil
	}
	return false, nil
}
func ExistMacroByNamespaceKey(namespace, key string) (uint, error) {
	var macro Macro
	err := db.Debug().Select("id").Where("namespace = ?", namespace).Where("`key` = ?", key).First(&macro).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	if macro.ID > 0 {
		return macro.ID, nil
	}
	return 0, nil
}
