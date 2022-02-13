package models

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/jinzhu/gorm"
)

// describe  guarder and more idc or cloud providers info
type Guarder struct {
	gorm.Model
	Datacenter string `json:"datacenter" gorm:"unique;not null;comment:'aliyun-i,huaweiyun-j,office'"`
	Provider   string `json:"provider" gorm:"not null;comment:'provider like aliyun, office, so on'"`
	IP         string `json:"ip" gorm:"not null;comment:'the ip of the guarder'"`
	Port       uint   `json:"port" gorm:"default 1;comment:'the port of guarder for using'"`
	Extras     string `json:"extras" gorm:"comment:'guarder额外信息'"`
}

type Host struct {
	gorm.Model

	Hostname    string `json:"hostname" gorm:"comment:'主机名'"`
	GID         uint   `json:"gid" gorm:"comment:'主机类型guarder对应id'"`
	OS          string `json:"os" gorm:"comment:'操作系统类型'"`
	Cores       uint   `json:"cores" gorm:"comment:'核心数'"`
	MemorySize  uint64 `json:"memory_size" gorm:"comment:'内存大小'"`
	Status      uint   `json:"status" gorm:"comment:'主机状态'"`
	Description string `json:"description"`
	Extras      string `json:"extras" gorm:"comment:'主机额外信息'"`
	Uptime      uint64 `json:"uptime" gorm:"comment:'主机启动时间'"`
	CreatedBy   uint   `json:"created_by"`
	UpdatedBy   uint   `json:"updated_by"`
}

type HostIP struct {
	gorm.Model
	HID uint `json:"hid" gorm:"comment:'主机id'"`
	// ip not unique for idcs and vpcs
	IP string `json:"ip" gorm:"comment:'ip';not null;"`
}

type VHost struct {
	ID          uint   `json:"id"`
	Hostname    string `json:"hostname"`
	Provider    string `json:"provider"`
	Datacenter  string `json:"datacenter"`
	OS          string `json:"os"`
	IP          string `json:"ip"`
	Extras      string `json:"extras"`
	Status      uint   `json:"status"`
	CreatedAt   string `json:"createdat"`
	UpdatedAt   string `json:"updatedat"`
	Description string `json:"description"`
	CID         string `json:"cid"`
	Cluster     string `json:"cluster"`
	AID         string `json:"aid"`
	App         string `json:"app"`
	PID         string `json:"pid"`
	Product     string `json:"product"`
}

func GetHost(id uint) (*Host, error) {
	var host Host
	err := db.Where("id = ? and deleted_at is NULL", id).First(&host).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &host, nil
}

func GetHostByHostnameAndIP(hostname, ip string) (*Host, error) {
	var host Host
	err := db.Table("host").Joins("left join host_ip on host_ip.h_id = host.id").Where("hostname = ? and host_ip.ip = ?", hostname, ip).First(&host).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &host, nil
}

func GetHosts(data map[string]interface{}) ([]*Host, error) {
	var hosts []*Host

	cid := data["cid"].(uint)
	db = db.Table("host").Joins("left join cluster_host on cluster_host.h_id = host.id")
	if cid != 0 {
		db = db.Where("host_cluster.c_id = ?", cid)
	}
	err := db.Find(&hosts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return hosts, nil
}

func GetGuarders() ([]*Guarder, error) {
	var guarders []*Guarder
	err := db.Where("deleted_at is NULL").Find(&guarders).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return guarders, nil
}

func EditHost(id uint, ip string, h *Host) error {
	var host Host
	var hostIP HostIP

	ips := strings.Replace(ip, " ", ",", -1)
	ipList := strings.Split(ips, ",")
	fmt.Println("ipList is ", ipList)
	if ip != "" {
		db.Unscoped().Where("h_id = ?", id).Delete(&hostIP)
		for _, ip := range ipList {
			if ip == "" {
				continue
			}
			if net.ParseIP(ip).To4() != nil {
				_, err := AddHostIP(id, ip)
				if err != nil {
					log.Println(err)
					return err
				}
			} else {
				log.Println("ip is not ipv4 ", ip)
			}

		}
	}

	if err := db.Model(&host).Where("id = ? AND deleted_at is NULL", id).Update(*h).Error; err != nil {
		return err
	}

	return nil
}

func AddHost(data map[string]interface{}) (id uint, err error) {
	host := Host{
		Hostname:    data["hostname"].(string),
		OS:          data["os"].(string),
		GID:         data["gid"].(uint),
		Cores:       data["cores"].(uint),
		MemorySize:  data["memory_size"].(uint64),
		Description: data["description"].(string),
		Status:      data["status"].(uint),
		Extras:      data["extras"].(string),
		Uptime:      data["uptime"].(uint64),
		CreatedBy:   data["created_by"].(uint),
	}

	log.Print(host)

	var guarder Guarder
	err = db.Select("id").Where("id = ?", host.GID).First(&guarder).Error
	if err != nil {
		log.Printf("the gid %d not exists, please check!", host.GID)
		return 0, err
	}

	if err = db.Create(&host).Error; err != nil {
		return 0, err
	}

	return host.ID, nil
}

func AddHostIP(hid uint, ip string) (id uint, err error) {
	var hostIP = HostIP{
		HID: hid,
		IP:  ip,
	}
	var hip HostIP
	err = db.Debug().Select("id").Where("h_id = ? and ip = ?", hid, ip).First(&hip).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	if hip.ID <= 0 {
		if err := db.Create(&hostIP).Error; err != nil {
			return 0, err
		}
	}
	return hostIP.ID, nil
}

func DeleteHost(id uint) error {
	if err := db.Where("id = ?", id).Delete(&Host{}).Error; err != nil {
		return err
	}

	if err := db.Where("hid = ?", id).Delete(&HostIP{}).Error; err != nil {
		return err
	}
	return nil
}

func CleanAllHost() error {
	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&Host{}).Error; err != nil {
		return err
	}

	if err := db.Unscoped().Where("deleted_at is not NULL").Delete(&HostIP{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistHostByID(id uint) (bool, error) {
	var host Host
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&host).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if host.ID > 0 {
		return true, nil
	}

	return false, nil
}
func GetGuarderByID(id uint) (*Guarder, error) {
	var guarder Guarder
	err := db.Where("id = ? and deleted_at is NULL", id).First(&guarder).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if guarder.ID > 0 {
		return &guarder, nil
	}

	return nil, nil
}

func GetHostTotal() (uint, error) {
	var count uint
	if err := db.Model(&Host{}).Where("deleted_at is null").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func QueryHosts(data map[string]interface{}) ([]*Host, error) {
	var hosts []*Host
	hostname := data["hostname"].(string)
	os := data["os"].(string)
	cores := data["cores"].(uint)
	status := data["status"].(uint)
	createdBy := data["created_by"].(uint)
	updatedBy := data["updated_by"].(uint)
	pageSize := data["pagesize"].(uint)
	pageNum := data["pagenum"].(uint)
	newdb := db.New()
	newdb = newdb.Where("deleted_at is null")
	//defer newdb.Close()
	if hostname != "null" {
		newdb = newdb.Where("hostname = ?", hostname)
	}
	if os != "null" {
		newdb = newdb.Where("os = ?", os)
	}
	if cores != 0 {
		newdb = newdb.Where("cores = ?", cores)
	}
	if status != 9999 {
		newdb = newdb.Where("status = ?", status)
	}
	if createdBy != 0 {
		newdb = newdb.Where("created_by = ?", createdBy)
	}
	if updatedBy != 0 {
		newdb = newdb.Where("updated_by = ?", updatedBy)
	}
	newdb = newdb.Limit(pageSize).Offset((pageNum - 1) * pageSize)

	if err := newdb.Find(&hosts).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return hosts, nil
}

func GetVerboseHostsWithPage(data map[string]interface{}) (int, []*VHost, error) {
	var vhosts []*VHost
	newdb := db.New()
	hostname := data["hostname"].(string)
	os := data["os"].(string)
	cores := data["cores"].(uint)
	status := data["status"].(uint)
	createdBy := data["created_by"].(uint)
	updatedBy := data["updated_by"].(uint)
	pagesize := data["pagesize"].(uint)
	pagenum := data["pagenum"].(uint)
	cid := data["cid"].(uint)
	aid := data["aid"].(uint)
	gid := data["gid"].(uint)
	key := data["key"].(string)

	if hostname != "null" && hostname != "" {
		newdb = newdb.Where("hostname = ?", hostname)
	}
	if os != "null" && os != "" {
		newdb = newdb.Where("os = ?", os)
	}
	if cores != 0 {
		newdb = newdb.Where("cores = ?", cores)
	}
	if status != 9999 {
		newdb = newdb.Where("status = ?", status)
	}
	if createdBy != 0 {
		newdb = newdb.Where("created_by = ?", createdBy)
	}
	if updatedBy != 0 {
		newdb = newdb.Where("updated_by = ?", updatedBy)
	}
	if cid != 0 {
		newdb = newdb.Where("cluster_host.c_id = ?", cid)
	}
	if aid != 0 {
		newdb = newdb.Where("app_cluster.a_id = ?", aid)
	}
	if gid != 0 {
		newdb = newdb.Where("g_id = ?", gid)
	}

	newdb = newdb.Debug().Table("host").Select("host.id as id, host.created_at as created_at, host.updated_at as updated_at, host.hostname as hostname, guarder.datacenter as datacenter,guarder.provider as provider, host.os as os, host.extras as extras, host.status as status, group_concat(distinct(host_ip.ip)) as ip, host.description as description,group_concat(distinct(cluster.id)) as c_id, group_concat(distinct(cluster.name)) as cluster, group_concat(distinct(app.id)) as a_id, group_concat(distinct(app.name)) as app, group_concat(distinct(product.id)) as p_id, group_concat(distinct(product.name)) as product").Order("host.hostname ASC").Joins("left join host_ip on host_ip.h_id = host.id").Group("hostname").Joins("left join cluster_host on host.id = cluster_host.h_id").Joins("left join cluster on cluster.id = cluster_host.c_id").Joins("left join app_cluster on cluster.id = app_cluster.c_id").Joins("left join app on app.id = app_cluster.a_id").Joins("left join product on app.p_id = product.id").Joins("left join guarder on host.g_id = guarder.id")
	if key != "null" && key != "" {
		newdb = newdb.Where(fmt.Sprintf("host.hostname LIKE '%%%s%%' or host.os LIKE '%%%s%%' or host.status LIKE '%%%s%%' or host_ip.ip LIKE '%%%s%%' or cluster.name LIKE '%%%s%%' or app.name LIKE '%%%s%%' or product.name LIKE '%%%s%%'", key, key, key, key, key, key, key))
	}

	err := newdb.Scan(&vhosts).Error
	if err != nil {
		log.Println("get verbose join caught error")
		return 0, nil, err
	}
	total := len(vhosts)

	if pagesize != 0 && pagenum != 0 {
		newdb = newdb.Limit(pagesize).Offset((pagenum - 1) * pagesize)
	}
	err = newdb.Scan(&vhosts).Error
	if err != nil {
		log.Println("get verbose join caught error")
		return 0, nil, err
	}
	return total, vhosts, nil
}

func AddGuarder(data map[string]interface{}) (id uint, err error) {
	guarder := Guarder{
		Datacenter: data["datacenter"].(string),
		Provider:   data["provider"].(string),
		IP:         data["ip"].(string),
		Port:       data["port"].(uint),
		Extras:     data["extras"].(string),
	}
	if err := db.Create(&guarder).Error; err != nil {
		return 0, err
	}
	return guarder.ID, nil
}
func ExistGuarderByID(id uint) (bool, error) {
	var guarder Guarder
	err := db.Select("id").Where("id = ? and deleted_at is NULL", id).First(&guarder).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if guarder.ID > 0 {
		return true, nil
	}

	return false, nil
}
func EditGuarder(id uint, g *Guarder) error {
	var guarder Guarder

	if err := db.Model(&guarder).Where("id = ? AND deleted_at is NULL", id).Update(*g).Error; err != nil {
		return err
	}

	return nil
}
func DeleteGuarder(id uint) error {
	if err := db.Where("id = ?", id).Delete(&Guarder{}).Error; err != nil {
		return err
	}
	return nil
}
