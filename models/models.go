package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"go-admin/pkg/setting"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(
		&Guarder{},
		&Product{},
		&ProductUser{},
		&App{},
		&AppUser{},
		&AppEnv{},
		&AppCluster{},
		&Cluster{},
		&ClusterHost{},
		&ClusterUser{},
		&Department{},
		&Host{},
		&HostIP{},
		&Role{},
		&RoleUser{},
		&SSHKey{},
		&Tag{},
		&TagHost{},
		&TagApp{},
		&TagCluster{},
		&TagUser{},
		&User{},
		&UserType{},
		&ResourceType{},
		&Resource{},
		&Macro{},
		&HostDeploy{},
	)
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
