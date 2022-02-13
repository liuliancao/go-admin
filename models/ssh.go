package models

import "github.com/jinzhu/gorm"

type SSHKey struct {
	gorm.Model

	UID       int    `json:"uid"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}
