package clusterservice

import "go-admin/models"

type Cluster struct {
	ID          uint
	Name        string
	Description string
	Status      uint
	CreatedBy   uint
	UpdatedBy   uint
}
type ClusterHost struct {
	CID       uint `json:"c_id"`
	HID       uint `json:"h_id"`
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}

type ClusterUser struct {
	CID       uint `json:"c_id"`
	UID       uint `json:"u_id"`
	CreatedBy uint `json:"created_by"`
}

func (c *Cluster) Add() (id uint, err error) {
	cluster := map[string]interface{}{
		"name":        c.Name,
		"description": c.Description,
		"status":      c.Status,
		"created_by":  c.CreatedBy,
	}

	id, err = models.AddCluster(cluster)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *Cluster) Edit() error {
	return models.EditCluster(c.ID, map[string]interface{}{
		"name":        c.Name,
		"description": c.Description,
		"status":      c.Status,
		"updated_by":  c.UpdatedBy,
	})
}

func (c *Cluster) Get() (*models.Cluster, error) {
	cluster, err := models.GetCluster(c.ID)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (c *Cluster) GetAll(aid uint) ([]*models.Cluster, error) {
	return models.GetClusters(map[string]interface{}{
		"aid": aid,
	})
}

func (c *Cluster) ExistByID() (bool, error) {
	return models.ExistClusterByID(c.ID)
}

func (c *Cluster) Delete() error {
	return models.DeleteCluster(c.ID)
}
func (c *ClusterHost) Add() (id uint, err error) {
	return models.AddClusterHost(c.CID, c.HID, c.CreatedBy)
}
func (cu *ClusterUser) Add() (id uint, err error) {
	return models.AddClusterUser(cu.CID, cu.UID, cu.CreatedBy)
}

func (cu *ClusterUser) Get(cid uint) ([]*models.User, error) {
	return models.GetClusterUsers(cid)
}
