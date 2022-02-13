package models

import "github.com/jinzhu/gorm"

// HostDeploy used for host deploy info
type HostDeploy struct {
	gorm.Model

	AID uint `json:"aid" gorm:"not null;comment:'应用ID'"`
	HID uint `json:"hid" gorm:"not null;comment:'主机ID'"`

	Status    string `json:"status" gorm:"default:'init';comment:'当前服务器发布状态'"`
	Version   string `json:"version" gorm:"default:null;comment:'当前版本'"`
	Stage     string `json:"stage" gorm:"comment:'阶段'"`
	Message   string `json:"message" gorm:"default:null;comment:'callback等信息'"`
	AppMacros string `json:"app_macros" gorm:"default:null;comment:'json like macros list, 保留当时config'"`
	CreatedBy uint   `json:"created_by"`
	UpdateBy  uint   `json:"updated_by"`
}

func AddHostDeploy(data map[string]interface{}) error {
	hostdeploy := HostDeploy{
		AID:       data["aid"].(uint),
		HID:       data["hid"].(uint),
		Status:    data["status"].(string),
		Version:   data["version"].(string),
		Stage:     data["stage"].(string),
		Message:   data["message"].(string),
		AppMacros: data["app_macros"].(string),
		CreatedBy: data["created_by"].(uint),
	}
	if err := db.Create(&hostdeploy).Error; err != nil {
		return err
	}

	return nil
}
func GetHostDeploy(id uint) (*HostDeploy, error) {
	var hostdeploy HostDeploy
	err := db.Where("id = ? and deleted_at is NULL", id).First(&hostdeploy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &hostdeploy, nil
}

func GetHostDeploys(data map[string]interface{}) ([]*HostDeploy, error) {
	var hostdeploys []*HostDeploy
	newdb := db.New()
	aid := data["a_id"].(uint)
	hid := data["h_id"].(uint)
	status := data["status"].(string)
	version := data["version"].(string)
	stage := data["stage"].(string)
	createdBy := data["created_by"].(uint)

	if aid != 0 {
		newdb = newdb.Where("a_id = ?", aid)
	}
	if hid != 0 {
		newdb = newdb.Where("h_id = ?", hid)
	}
	if status != "null" && status != "" {
		newdb = newdb.Where("status = ?", status)
	}
	if version != "null" && version != "" {
		newdb = newdb.Where("version = ?", version)
	}
	if stage != "null" && stage != "" {
		newdb = newdb.Where("stage = ?", stage)
	}
	if createdBy != 0 {
		newdb = newdb.Where("created_by = ?", createdBy)
	}

	err := newdb.Debug().Where("deleted_at is NULL").Find(&hostdeploys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return hostdeploys, nil
}

func EditHostDeploy(id uint, data interface{}) error {
	var hostdeploy HostDeploy
	if err := db.Model(&hostdeploy).Where("id = ? AND deleted_at is NULL", id).Update(data).Error; err != nil {
		return err
	}

	return nil
}

func DeleteHostDeploy(id uint) error {
	if err := db.Where("id = ? and deleted_at is NULL", id).Delete(&HostDeploy{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllHostDeploy() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&HostDeploy{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistHostDeployByID(id uint) (bool, error) {
	var hostdeploy HostDeploy
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&hostdeploy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if hostdeploy.ID > 0 {
		return true, nil
	}

	return false, nil
}
