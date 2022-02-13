package hostservice

import (
	"errors"
	"go-admin/models"
	"log"
	"strings"
)

type Host struct {
	ID       uint
	Hostname string
	OS       string
	Cores    uint
	IP       string
	GID      uint
	//cluster id
	CID uint
	//app id
	AID         uint
	MemorySize  uint64
	Status      uint
	Description string
	Extras      string
	Uptime      uint64
	CreatedBy   uint
	UpdatedBy   uint
	PageSize    uint
	PageNum     uint
	Key         string
}

type Guarder struct {
	ID         uint
	Datacenter string `json:"datacenter"`
	Provider   string `json:"provider"`
	Port       uint   `json:"port"`
	IP         string `json:"ip"`
	Extras     string `json:"extras"`
}

func (h *Host) Add() (id uint, err error) {
	ipList := strings.Split(h.IP, ",")

	guarder, err := models.GetGuarderByID(h.GID)
	if guarder == nil {
		return 0, errors.New("guarder not found")
	}
	if err != nil {
		return 0, err
	}

	host := map[string]interface{}{
		"hostname":    h.Hostname,
		"os":          h.OS,
		"gid":         h.GID,
		"cores":       h.Cores,
		"memory_size": h.MemorySize,
		"status":      h.Status,
		"description": h.Description,
		"extras":      h.Extras,
		"uptime":      h.Uptime,
		"created_by":  h.CreatedBy,
	}

	id, err = models.AddHost(host)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// ip allows duplicate for multiple vpcs and idc rooms.
	for _, ip := range ipList {
		if ip != "" {
			_, err := models.AddHostIP(id, ip)
			if err != nil {
				log.Println(err)
				return 0, err
			}
		}
	}
	return id, nil
}

func (h *Host) Edit(ip string) error {
	host := &models.Host{
		Hostname:    h.Hostname,
		GID:         h.GID,
		OS:          h.OS,
		Cores:       h.Cores,
		MemorySize:  h.MemorySize,
		Status:      h.Status,
		Description: h.Description,
		Extras:      h.Extras,
		Uptime:      h.Uptime,
		UpdatedBy:   h.UpdatedBy,
	}
	return models.EditHost(h.ID, ip, host)
}

func (h *Host) Get() (*models.Host, error) {
	host, err := models.GetHost(h.ID)
	if err != nil {
		return nil, err
	}

	return host, nil
}

func (h *Host) GetByHostnameAndIP() (*models.Host, error) {
	host, err := models.GetHostByHostnameAndIP(h.Hostname, h.IP)
	if err != nil {
		return nil, err
	}
	h.ID = host.ID
	return host, nil
}
func (h *Host) GetAll(cid uint) ([]*models.Host, error) {
	return models.GetHosts(map[string]interface{}{
		"cid": cid,
	})
}

func (h *Host) ExistByID() (bool, error) {
	return models.ExistHostByID(h.ID)
}

func (h *Host) GetTotal() (total uint, err error) {
	return models.GetHostTotal()
}
func (h *Host) Delete() error {
	return models.DeleteHost(h.ID)
}
func (h *Host) GetVerboseHostsWithPage() (int, []*models.VHost, error) {
	/*if h.PageNum == 0 || h.PageSize == 0 {
		return nil, errors.New("need page number and page size params")
	}*/

	return models.GetVerboseHostsWithPage(map[string]interface{}{
		"hostname":   h.Hostname,
		"os":         h.OS,
		"cores":      h.Cores,
		"ip":         h.IP,
		"cid":        h.CID,
		"aid":        h.AID,
		"gid":        h.GID,
		"status":     h.Status,
		"created_by": h.CreatedBy,
		"updated_by": h.UpdatedBy,
		"pagenum":    h.PageNum,
		"pagesize":   h.PageSize,
		"key":        h.Key,
	})
}

func (g *Guarder) Add() (id uint, err error) {
	guarder := map[string]interface{}{
		"datacenter": g.Datacenter,
		"ip":         g.IP,
		"port":       g.Port,
		"provider":   g.Provider,
		"extras":     g.Extras,
	}
	id, err = models.AddGuarder(guarder)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}
func (g *Guarder) GetAll() ([]*models.Guarder, error) {
	return models.GetGuarders()
}

func (g *Guarder) ExistByID() (bool, error) {
	return models.ExistGuarderByID(g.ID)
}
func (g *Guarder) Edit() error {
	guarder := &models.Guarder{
		Datacenter: g.Datacenter,
		IP:         g.IP,
		Port:       g.Port,
		Provider:   g.Provider,
		Extras:     g.Extras,
	}
	return models.EditGuarder(g.ID, guarder)
}
func (g *Guarder) Delete() error {
	return models.DeleteGuarder(g.ID)
}
