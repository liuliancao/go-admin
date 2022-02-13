package models

import (
	"errors"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Cluster struct {
	gorm.Model

	Name        string `json:"name" gorm:"unique;not null;comment:'集群名称'"`
	Description string `json:"description"`
	Status      uint   `json:"status" gorm:"comment:'集群状态'"`
	CreatedBy   uint   `json:"created_by"`
	UpdatedBy   uint   `json:"updated_by"`
}

type ClusterHost struct {
	gorm.Model
	CID       uint `json:"c_id" gorm:"comment:'集群主机关联集群id'"`
	HID       uint `json:"h_id" gorm:"comment:'集群主机关联主机id'"`
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}
type ClusterUser struct {
	gorm.Model
	CID       uint `json:"c_id" gorm:"comment:'集群用于关联集群id'"`
	UID       uint `json:"u_id" gorm:"comment:'集群用户关联用户id'"`
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

func GetCluster(id uint) (*Cluster, error) {
	var cluster Cluster
	err := db.Where("id = ? and deleted_at is NULL", id).First(&cluster).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &cluster, nil
}

func GetClusters(data map[string]interface{}) ([]*Cluster, error) {
	var clusters []*Cluster

	aid := data["aid"].(uint)

	newdb := db.New()
	newdb = newdb.Debug().Table("cluster").Select("cluster.id as id, cluster.name as name, cluster.description as description, cluster.status as status, cluster.created_by as created_by, cluster.updated_by as updated_by").Joins("left join app_cluster on app_cluster.c_id = cluster.id")
	print(aid)
	if aid != 0 {
		newdb = newdb.Where("app_cluster.a_id = ?", aid)
	}
	err := newdb.Find(&clusters).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return clusters, nil
}

func EditCluster(id uint, data interface{}) error {
	var cluster Cluster
	if err := db.Model(&cluster).Where("id = ? AND deleted_at is NULL", id).Update(data).Error; err != nil {
		return err
	}

	return nil
}

func AddCluster(data map[string]interface{}) (id uint, err error) {
	cluster := Cluster{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		Status:      data["status"].(uint),
		CreatedBy:   data["created_by"].(uint),
	}
	if err := db.Create(&cluster).Error; err != nil {
		return 0, err
	}

	return cluster.ID, nil
}
func AddClusterHost(cid, hid uint, createdBy uint) (id uint, err error) {

	if exists, _ := ExistClusterByID(cid); exists == false {
		return 0, errors.New("cluster not found " + strconv.Itoa(int(cid)))
	}
	if exists, _ := ExistHostByID(hid); exists == false {
		return 0, errors.New("host not found" + strconv.Itoa(int(hid)))
	}

	clusterHost := ClusterHost{
		CID:       cid,
		HID:       hid,
		CreatedBy: createdBy,
	}
	if err := db.Create(&clusterHost).Error; err != nil {
		log.Fatalf("create cluster %d with host %d failed: %v", cid, hid, err)
		return 0, err
	}
	return clusterHost.ID, nil
}
func DeleteClusterHost(cid, hid uint) error {
	if err := db.Where("cid = ? AND hid = ? AND deleted_at is NULL", cid, hid).Delete(&ClusterHost{}).Error; err != nil {
		return err
	}

	return nil
}
func AddClusterUser(cid, uid uint, createdBy uint) (id uint, err error) {

	if exists, _ := ExistClusterByID(cid); exists == false {
		return 0, errors.New("cluster not found " + strconv.Itoa(int(cid)))
	}
	if exists, _ := ExistUserByID(uid); exists == false {
		return 0, errors.New("host not found" + strconv.Itoa(int(uid)))
	}

	clusterUser := ClusterUser{
		CID:       cid,
		UID:       uid,
		CreatedBy: createdBy,
	}
	if err := db.Create(&clusterUser).Error; err != nil {
		log.Fatalf("create cluster %d with host %d failed: %v", cid, uid, err)
		return 0, err
	}
	return clusterUser.ID, nil
}

func GetClusterUsers(cid uint) ([]*User, error) {
	var users []*User
	err := db.Debug().Table("user").Where("id = any(select uid from cluster_user where c_id = ?)", cid).Scan(&users).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}

func DeleteClusterUser(cid, uid uint) error {
	if err := db.Where("cid = ? AND uid = ? AND deleted_at is NULL", cid, uid).Delete(&ClusterUser{}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteCluster(id uint) error {
	if err := db.Where("id = ?", id).Delete(&Cluster{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllCluster() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&Cluster{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistClusterByID(id uint) (bool, error) {
	var cluster Cluster
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&cluster).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if cluster.ID > 0 {
		return true, nil
	}

	return false, nil
}
